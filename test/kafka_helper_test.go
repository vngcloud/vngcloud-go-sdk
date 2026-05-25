package test

import (
	larchivezip "archive/zip"
	lbytes "bytes"
	lstrings "strings"
	ltesting "testing"

	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lskafkaV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/kafka/v1"
)

const (
	// All test fixtures are FAKE — modeled after the real sample at
	// /home/tytv2/tcp-logs-test/user-f7c2e02a-3e82-4990-a4d5-8069fc29cbcd/ but
	// regenerated with throwaway material so nothing real leaks into the repo.
	testAuthenCredsUserId = "f7c2e02a-3e82-4990-a4d5-8069fc29cbcd"

	testFakeCACertMTLS   = "-----BEGIN CERTIFICATE-----\nMIIB-FAKE-MTLS-CA\n-----END CERTIFICATE-----\n"
	testFakeUserCertMTLS = "-----BEGIN CERTIFICATE-----\nMIIB-FAKE-MTLS-USER-CRT\n-----END CERTIFICATE-----\n"
	testFakeUserKeyMTLS  = "-----BEGIN PRIVATE KEY-----\nMIIB-FAKE-MTLS-USER-KEY\n-----END PRIVATE KEY-----\n"

	testFakeCACertSASL = "-----BEGIN CERTIFICATE-----\nMIIB-FAKE-SASL-CA\n-----END CERTIFICATE-----\n"
	testFakeSASLPwd    = "fakeSaslPassword12345"
	testFakeSASLUser   = "kafkauser-50077359-71df-4785-84b6-55364bedf444"
	testFakeJaasConfig = `org.apache.kafka.common.security.scram.ScramLoginModule required username="kafkauser-50077359-71df-4785-84b6-55364bedf444" password="fakeSaslPassword12345";`

	// Files that should be IGNORED by the parser (Java keystore artifacts + macOS junk).
	testFakeP12         = "PK\x03\x04 fake p12 content"
	testFakeKeystorePwd = "garbage12345"
	testFakeConfigProps = "security.protocol=SSL\nssl.keystore.location=user.p12\n"
)

// buildTestAuthenCredsZip creates a zip mirroring the real vDB authen-creds layout:
//
//	user-{userId}/
//	├── mtls/
//	│   ├── ca.crt, ca.p12, ca.password
//	│   ├── user.crt, user.key, user.p12, user.password
//	│   └── config.properties
//	└── sasl/
//	    ├── ca.crt, ca.p12, ca.password
//	    ├── password, sasl.jaas.config, config.properties
func buildTestAuthenCredsZip(t *ltesting.T) []byte {
	t.Helper()
	var buf lbytes.Buffer
	zw := larchivezip.NewWriter(&buf)
	root := "user-" + testAuthenCredsUserId + "/"

	add := func(path, content string) {
		w, err := zw.Create(path)
		if err != nil {
			t.Fatalf("zip create %s: %v", path, err)
		}
		if _, err := w.Write([]byte(content)); err != nil {
			t.Fatalf("zip write %s: %v", path, err)
		}
	}

	// mtls/
	add(root+"mtls/ca.crt", testFakeCACertMTLS)
	add(root+"mtls/user.crt", testFakeUserCertMTLS)
	add(root+"mtls/user.key", testFakeUserKeyMTLS)
	add(root+"mtls/ca.p12", testFakeP12)
	add(root+"mtls/user.p12", testFakeP12)
	add(root+"mtls/ca.password", testFakeKeystorePwd)
	add(root+"mtls/user.password", testFakeKeystorePwd)
	add(root+"mtls/config.properties", testFakeConfigProps)

	// sasl/
	add(root+"sasl/ca.crt", testFakeCACertSASL)
	add(root+"sasl/password", testFakeSASLPwd)
	add(root+"sasl/sasl.jaas.config", testFakeJaasConfig)
	add(root+"sasl/ca.p12", testFakeP12)
	add(root+"sasl/ca.password", testFakeKeystorePwd)
	add(root+"sasl/config.properties", testFakeConfigProps)

	// macOS junk
	add(root+".DS_Store", "should be skipped")

	if err := zw.Close(); err != nil {
		t.Fatalf("zip close: %v", err)
	}
	return buf.Bytes()
}

func TestParseAuthenCredsZip_BothMTLSAndSASL(t *ltesting.T) {
	zipBytes := buildTestAuthenCredsZip(t)

	creds, err := lskafkaV1.ParseAuthenCredsZip(zipBytes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if creds.UserId != testAuthenCredsUserId {
		t.Errorf("UserId = %q, want %q", creds.UserId, testAuthenCredsUserId)
	}

	if !creds.HasMTLS() {
		t.Fatal("expected HasMTLS() = true")
	}
	if got := string(creds.MTLS.CACert); got != testFakeCACertMTLS {
		t.Errorf("MTLS.CACert mismatch:\n  got:  %q\n  want: %q", got, testFakeCACertMTLS)
	}
	if got := string(creds.MTLS.UserCert); got != testFakeUserCertMTLS {
		t.Errorf("MTLS.UserCert mismatch")
	}
	if got := string(creds.MTLS.UserKey); got != testFakeUserKeyMTLS {
		t.Errorf("MTLS.UserKey mismatch")
	}

	if !creds.HasSASL() {
		t.Fatal("expected HasSASL() = true")
	}
	if creds.SASL.Mechanism != "SCRAM-SHA-512" {
		t.Errorf("SASL.Mechanism = %q, want SCRAM-SHA-512", creds.SASL.Mechanism)
	}
	if creds.SASL.Username != testFakeSASLUser {
		t.Errorf("SASL.Username = %q, want %q", creds.SASL.Username, testFakeSASLUser)
	}
	if creds.SASL.Password != testFakeSASLPwd {
		t.Errorf("SASL.Password = %q, want %q", creds.SASL.Password, testFakeSASLPwd)
	}
	if string(creds.SASL.CACert) != testFakeCACertSASL {
		t.Errorf("SASL.CACert mismatch")
	}
}

func TestParseAuthenCredsZip_EmptyBytes(t *ltesting.T) {
	if _, err := lskafkaV1.ParseAuthenCredsZip(nil); err == nil {
		t.Fatal("expected error for nil input")
	}
	if _, err := lskafkaV1.ParseAuthenCredsZip([]byte{}); err == nil {
		t.Fatal("expected error for empty bytes")
	}
}

func TestParseAuthenCredsZip_OnlyMTLS(t *ltesting.T) {
	var buf lbytes.Buffer
	zw := larchivezip.NewWriter(&buf)
	root := "user-onlymtls/"
	_, _ = zw.Create(root + "mtls/ca.crt")
	w, _ := zw.Create(root + "mtls/user.crt")
	_, _ = w.Write([]byte(testFakeUserCertMTLS))
	w, _ = zw.Create(root + "mtls/user.key")
	_, _ = w.Write([]byte(testFakeUserKeyMTLS))
	w, _ = zw.Create(root + "mtls/ca.crt")
	_, _ = w.Write([]byte(testFakeCACertMTLS))
	_ = zw.Close()

	creds, err := lskafkaV1.ParseAuthenCredsZip(buf.Bytes())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !creds.HasMTLS() {
		t.Error("expected HasMTLS() = true")
	}
	if creds.HasSASL() {
		t.Error("expected HasSASL() = false")
	}
}

func TestBuildBootstrapServers_MTLSPublic(t *ltesting.T) {
	cluster := &lsentity.KafkaCluster{
		Id:           "test-cluster-mtls",
		PublicAccess: true,
		FloatingIps:  []string{"61.28.236.199", "61.28.235.33", "61.28.235.4"},
	}

	got, err := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthMTLS)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "61.28.236.199:9194,61.28.235.33:9194,61.28.235.4:9194"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBuildBootstrapServers_SASLPublic(t *ltesting.T) {
	cluster := &lsentity.KafkaCluster{
		Id:           "test-cluster-sasl",
		PublicAccess: true,
		FloatingIps:  []string{"61.28.236.199", "61.28.235.33", "61.28.235.4"},
	}

	got, err := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthSASL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "61.28.236.199:9196,61.28.235.33:9196,61.28.235.4:9196"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBuildBootstrapServers_PublicAccessDisabled(t *ltesting.T) {
	cluster := &lsentity.KafkaCluster{
		Id:           "test-cluster",
		PublicAccess: false,
		FloatingIps:  []string{"61.28.236.199"},
	}

	_, err := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthMTLS)
	if err == nil {
		t.Fatal("expected error when publicAccess=false")
	}
	if !lstrings.Contains(err.Error(), "publicAccess") {
		t.Errorf("error %q should mention publicAccess", err.Error())
	}
}

func TestBuildBootstrapServers_NoFloatingIPs(t *ltesting.T) {
	cluster := &lsentity.KafkaCluster{
		Id:           "test-cluster",
		PublicAccess: true,
		FloatingIps:  []string{},
	}

	_, err := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthMTLS)
	if err == nil {
		t.Fatal("expected error when floatingIps is empty")
	}
}

func TestBuildBootstrapServers_UnsupportedAuthMode(t *ltesting.T) {
	cluster := &lsentity.KafkaCluster{
		Id:           "test-cluster",
		PublicAccess: true,
		FloatingIps:  []string{"61.28.236.199"},
	}

	_, err := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthMode("plaintext"))
	if err == nil {
		t.Fatal("expected error for unsupported auth mode")
	}
}

func TestBuildBootstrapServers_NilCluster(t *ltesting.T) {
	if _, err := lskafkaV1.BuildBootstrapServers(nil, lskafkaV1.AuthMTLS); err == nil {
		t.Fatal("expected error for nil cluster")
	}
}
