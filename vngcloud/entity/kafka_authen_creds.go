package entity

// KafkaAuthenCreds is the result of unzipping the response of
// GET /vdb-kafka/clusters/{clusterId}/users/{userId}/authen-creds.
// The zip always contains both mtls/ and sasl/ subdirectories so the caller
// can pick the auth mode at consume time (V1: mtls preferred).
type KafkaAuthenCreds struct {
	UserId string          // folder root inside the zip: user-{vdb_user_id}
	MTLS   *KafkaMTLSCreds // nil if zip does not contain mtls/
	SASL   *KafkaSASLCreds // nil if zip does not contain sasl/
}

// KafkaMTLSCreds holds PEM-encoded bytes for kafka-go's tls.Config.
// PKCS12 files (.p12) and their passwords are intentionally dropped — those
// are Java keystore artifacts unused by Go clients.
type KafkaMTLSCreds struct {
	CACert   []byte
	UserCert []byte
	UserKey  []byte
}

// KafkaSASLCreds carries SCRAM-SHA-512 credentials plus CA for SASL_SSL.
type KafkaSASLCreds struct {
	CACert    []byte
	Mechanism string // "SCRAM-SHA-512" (only supported mechanism today)
	Username  string
	Password  string
}

func (s *KafkaAuthenCreds) HasMTLS() bool {
	return s.MTLS != nil && len(s.MTLS.CACert) > 0 && len(s.MTLS.UserCert) > 0 && len(s.MTLS.UserKey) > 0
}

func (s *KafkaAuthenCreds) HasSASL() bool {
	return s.SASL != nil && len(s.SASL.CACert) > 0 && s.SASL.Username != "" && s.SASL.Password != ""
}
