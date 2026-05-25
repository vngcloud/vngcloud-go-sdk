package v1

import (
	larchivezip "archive/zip"
	lbytes "bytes"
	lerrors "errors"
	lfmt "fmt"
	lio "io"
	lpath "path"
	lregexp "regexp"
	lstrings "strings"

	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)

type AuthMode string

const (
	AuthMTLS AuthMode = "mtls"
	AuthSASL AuthMode = "sasl"

	saslMechanismSCRAMSHA512 = "SCRAM-SHA-512"
)

// publicPortMatrix maps each auth mode to the broker listener port on public IPs.
// Confirmed by vDB May 2026:
//
//	mTLS public  → :9194
//	SASL public  → :9196
//	(private endpoints :9094 / :9096 are deferred to V2)
var publicPortMatrix = map[AuthMode]int{
	AuthMTLS: 9194,
	AuthSASL: 9196,
}

// jaasUsernamePasswordRe extracts username + password from a SASL JAAS config
// line like:
//
//	org.apache.kafka.common.security.scram.ScramLoginModule required username="..." password="...";
var jaasUsernamePasswordRe = lregexp.MustCompile(`username="([^"]+)"\s+password="([^"]+)"`)

// ParseAuthenCredsZip decodes the ZIP returned by GetUserAuthenCreds.
// The zip contains both mtls/ and sasl/ subfolders so the caller can pick the
// auth mode at consume time. Java keystore artifacts (*.p12 + *.password +
// config.properties) are skipped — Go clients work directly with PEM bytes.
func ParseAuthenCredsZip(pzipBytes []byte) (*lsentity.KafkaAuthenCreds, error) {
	if len(pzipBytes) == 0 {
		return nil, lerrors.New("empty zip bytes")
	}

	zr, err := larchivezip.NewReader(lbytes.NewReader(pzipBytes), int64(len(pzipBytes)))
	if err != nil {
		return nil, lfmt.Errorf("open zip: %w", err)
	}

	creds := &lsentity.KafkaAuthenCreds{}
	mtls := &lsentity.KafkaMTLSCreds{}
	sasl := &lsentity.KafkaSASLCreds{Mechanism: saslMechanismSCRAMSHA512}

	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}

		// Folder root is "user-{uuid}/" — strip the first segment.
		parts := lstrings.SplitN(f.Name, "/", 2)
		if len(parts) < 2 {
			continue
		}
		if creds.UserId == "" && lstrings.HasPrefix(parts[0], "user-") {
			creds.UserId = lstrings.TrimPrefix(parts[0], "user-")
		}

		rel := parts[1]
		base := lpath.Base(rel)

		// Skip macOS metadata + Java artifacts (worker doesn't need PKCS12 + keystore passwords).
		if base == ".DS_Store" || lstrings.HasSuffix(base, ".p12") ||
			base == "user.password" || base == "ca.password" || base == "config.properties" {
			continue
		}

		content, rerr := readZipEntry(f)
		if rerr != nil {
			return nil, lfmt.Errorf("read %s: %w", f.Name, rerr)
		}

		switch {
		case lstrings.HasPrefix(rel, "mtls/"):
			switch base {
			case "ca.crt":
				mtls.CACert = content
			case "user.crt":
				mtls.UserCert = content
			case "user.key":
				mtls.UserKey = content
			}

		case lstrings.HasPrefix(rel, "sasl/"):
			switch base {
			case "ca.crt":
				sasl.CACert = content
			case "password":
				sasl.Password = lstrings.TrimSpace(string(content))
			case "sasl.jaas.config":
				// Username lives only inside the JAAS line — must parse.
				if m := jaasUsernamePasswordRe.FindSubmatch(content); len(m) == 3 {
					sasl.Username = string(m[1])
					if sasl.Password == "" {
						sasl.Password = string(m[2])
					}
				}
			}
		}
	}

	if mtls.CACert != nil || mtls.UserCert != nil || mtls.UserKey != nil {
		creds.MTLS = mtls
	}
	if sasl.CACert != nil || sasl.Username != "" || sasl.Password != "" {
		creds.SASL = sasl
	}

	return creds, nil
}

func readZipEntry(f *larchivezip.File) ([]byte, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return lio.ReadAll(rc)
}

// BuildBootstrapServers assembles a comma-separated "ip:port,ip:port,ip:port"
// string from a KafkaCluster + chosen auth mode.
//
// V1 only supports the **public** endpoint set (floatingIps). The cluster must
// have `publicAccess=true` and ≥1 floating IP; otherwise the function returns
// an error so the caller can surface a clear "enable public access first"
// message back to the tenant.
func BuildBootstrapServers(pcluster *lsentity.KafkaCluster, pauth AuthMode) (string, error) {
	if pcluster == nil {
		return "", lerrors.New("cluster is nil")
	}
	if !pcluster.PublicAccess {
		return "", lfmt.Errorf("cluster %s: publicAccess is not enabled — tenant must turn it on via vDB portal", pcluster.Id)
	}
	if len(pcluster.FloatingIps) == 0 {
		return "", lfmt.Errorf("cluster %s: no floating IPs available", pcluster.Id)
	}
	port, ok := publicPortMatrix[pauth]
	if !ok {
		return "", lfmt.Errorf("unsupported auth mode: %q (expected %q or %q)", pauth, AuthMTLS, AuthSASL)
	}

	parts := make([]string, 0, len(pcluster.FloatingIps))
	for _, ip := range pcluster.FloatingIps {
		parts = append(parts, lfmt.Sprintf("%s:%d", ip, port))
	}
	return lstrings.Join(parts, ","), nil
}
