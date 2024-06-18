package entity

type Listener struct {
	UUID                            string
	Name                            string
	Description                     string
	Protocol                        string
	ProtocolPort                    int
	ConnectionLimit                 int
	DefaultPoolId                   string
	DefaultPoolName                 string
	TimeoutClient                   int
	TimeoutMember                   int
	TimeoutConnection               int
	AllowedCidrs                    string
	Headers                         []string
	CertificateAuthorities          []string
	DisplayStatus                   string
	CreatedAt                       string
	UpdatedAt                       string
	DefaultCertificateAuthority     *string
	ClientCertificateAuthentication *string
	ProgressStatus                  string
}
