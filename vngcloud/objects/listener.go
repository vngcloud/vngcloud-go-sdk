package objects

type Listener struct {
	UUID                            string   `json:"uuid"`
	Name                            string   `json:"name"`
	Description                     string   `json:"description,omitempty"`
	Protocol                        string   `json:"protocol"`
	ProtocolPort                    int      `json:"protocolPort"`
	ConnectionLimit                 int      `json:"connectionLimit"`
	DefaultPoolId                   string   `json:"defaultPoolId"`
	DefaultPoolName                 string   `json:"defaultPoolName"`
	TimeoutClient                   int      `json:"timeoutClient"`
	TimeoutMember                   int      `json:"timeoutMember"`
	TimeoutConnection               int      `json:"timeoutConnection"`
	AllowedCidrs                    string   `json:"allowedCidrs"`
	Headers                         []string `json:"headers"`
	CertificateAuthorities          []string `json:"certificateAuthorities"`
	DisplayStatus                   string   `json:"displayStatus"`
	CreatedAt                       string   `json:"createdAt"`
	UpdatedAt                       string   `json:"updatedAt"`
	DefaultCertificateAuthority     *string  `json:"defaultCertificateAuthority"`
	ClientCertificateAuthentication *string  `json:"clientCertificateAuthentication"`
	ProgressStatus                  string   `json:"progressStatus"`
}
