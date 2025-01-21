package test

import (
	"testing"

	v2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
)

func TestListCertificates(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewListCertificatesRequest()
	certs, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListCertificates(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if certs == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", certs)
	for _, pkg := range certs.Certificates {
		t.Logf("Package: %+v", pkg)
	}
	t.Log("PASS")
}

func TestGetCertificateById(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewGetCertificateByIdRequest("secret-84cb7a5e-b949-4f1b-a2e8-d2752e6e1181")
	cert, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().GetCertificateById(opt)
	if sdkerr != nil {
		t.Fatalf("Expect nil but got %+v", sdkerr)
	}

	if cert == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", cert)
	t.Log("PASS")
}

func TestCreateCertificate(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateCertificateRequest(
		"annd2-haha",
		"-----BEGIN CERTIFICATE-----\nMIIFPTCCAyUCFDmqCQVw9BUEsdQ52Rz3RZlUnovyMA0GCSqGSIb3DQEBCwUAMFsx\nCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl\ncm5ldCBXaWRnaXRzIFB0eSBMdGQxFDASBgNVBAMMC2V4YW1wbGUuY29tMB4XDTI0\nMDIxOTA5NDYwMVoXDTI1MDIxODA5NDYwMVowWzELMAkGA1UEBhMCQVUxEzARBgNV\nBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0\nZDEUMBIGA1UEAwwLZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAw\nggIKAoICAQCWj83KzHw9SG9L9AYfz4gmGtzOKEMaPj7+PSJR3UI1cj8lWyWyNXOZ\nz7nt75mVBTrESwd925RSiONxRNYif3dbEV3b+z4p0Wtq7lvDnCiaYnFoPAR4SFZl\nBA46eTsX/eWRXMaujOVzVC343g6kEZIM73brQUnhhfzrcGtBazj2UYnfT2UXIAdz\nMo7ZOSKZbBYsMFxOwXle2ocy8JY57lOZ3Y8p+9ZwipPC6pyz6KiSJcgxehCRGN6I\nWeIBnj0hpE3flVy4pq17BjkS84sHv3WT6PzPqVckq9dpxkgw5rTfo5kzZgbBZ60e\nXOzuPmqank+siU8wRAi+mZP0h1SyC5mZpSc7+MI9k6TpmBuKy+QsdMAUwFVNI3SN\nUoOeGuLdyuVO8Zogw+3BEL+HWX8lBPGwAuaO78B6jajTpSiWQXog2v8OmvtWdVp1\n4NU71bNgKBvTHXWsTVBoF3X6ZJw+/Ra59p/IPdcH6xlYoHV1CThNXaR2LP9AWh0r\nJr2uKHJGYr+k1Fndauq/KWaBv5hYwJ8wO8DG8JzHfbqqobYn9e4H6DEXqpK9lkXI\n+ocUZOJdn8zLYV638wB8q7zAqOBf88/6e+w7fbie2Fvuzc2A1Yyt/TixaytnaSl5\nijy5ibfEJoENIZx0REuuHc1hlVB42n3R5Lg4HXNaSWCR9MoXew659QIDAQABMA0G\nCSqGSIb3DQEBCwUAA4ICAQAPbT3OtUMFcoH4glny1LShqXf4b73t5CWeQWtrveB7\nfAeiz19G2y8QFjLg4YXS4zcEznjSPzU1rT4YcsaQKkNtRepGsMYYD6dSPYowU+8B\neXPMv+aNUQlviFABmb5BHvtVhScAF3pvysmcSeVYPdmDhR0pHATJ9cDIBxWSxGKI\nWJ2kNgxq139ik+c5hFCShYxCdgQEks+azB2XWUf5JctEJRu5dViKfh8LyraQrRbj\nOk2tntsiBLLQi32vmOiDgmii1/T3gkyweqjR7T2/PtdPD4bLHgj4oTHReI0KwBTY\nnPYvQxl2xRkgGtsQfg1LUc6oAefFMM5yMbxl2riQabR0CpsZ+XXPwr6b5VO/Hrrw\nyk2PBXXXXdmc9gPkUAIb4S3MUp5oh9i4wtspES5XN+tIWlaSDWKooDGr2Eu0lJmc\nLSqonVlfBnh4Wdpv9HeTbfBv6Tw8QauTvTMhdTvjBhswHbQP2Jf37HeJYT4Kjcui\nda1HjscJfGC5vFNWDD4sLeENVUCrtJ6Aj4cisEbjf405xqo4eSq9YIfzNCWELoPO\njtlt2lzLnTxNlzQKanw014VVFgxg51ryWsbk0P8zTDMUSuuXTM282vdayATjNuok\nW/yNBd1YYMXvztpvD1nwJks00Y+IjAR/8GR7fxrztchtPOEpchQ249wpX5A5qJjr\n3Q==\n-----END CERTIFICATE-----\n",
		v2.ImportOptsTypeOptTLS,
	).WithPrivateKey("-----BEGIN PRIVATE KEY-----\n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\nKQMcZALMoH96GAFsmvxHl7EPJMlZJw==\n-----END PRIVATE KEY-----\n")

	cert, err := vngcloud.VLBGateway().V2().LoadBalancerService().CreateCertificate(opt)
	if err != nil {
		t.Fatalf("Expect nil but got %+v", err)
	}

	if cert == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Logf("Result: %+v", cert)
	t.Log("PASS")
}

func TestDeleteCertificateById(t *testing.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewDeleteCertificateByIdRequest("secret-92318075-6622-48a3-88a7-daf2a8917917")
	err := vngcloud.VLBGateway().V2().LoadBalancerService().DeleteCertificateById(opt)
	if err != nil {
		t.Fatalf("Expect nil but got %+v", err)
	}

	t.Log("PASS")
}
