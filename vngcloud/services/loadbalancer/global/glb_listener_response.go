package global

// import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

// type CreateListenerResponse struct {
// 	UUID string `json:"uuid"`
// }

// type ListListenersByLoadBalancerIdResponse struct {
// 	Data []Listener `json:"data"`
// }

// type GetListenerByIdResponse struct {
// 	Data Listener `json:"data"`
// }

// type Listener struct {
// 	UUID                            string   `json:"uuid"`
// 	Name                            string   `json:"name"`
// 	Description                     string   `json:"description,omitempty"`
// 	Protocol                        string   `json:"protocol"`
// 	ProtocolPort                    int      `json:"protocolPort"`
// 	ConnectionLimit                 int      `json:"connectionLimit"`
// 	DefaultPoolId                   string   `json:"defaultPoolId"`
// 	DefaultPoolName                 string   `json:"defaultPoolName"`
// 	TimeoutClient                   int      `json:"timeoutClient"`
// 	TimeoutMember                   int      `json:"timeoutMember"`
// 	TimeoutConnection               int      `json:"timeoutConnection"`
// 	AllowedCidrs                    string   `json:"allowedCidrs"`
// 	Headers                         []string `json:"headers"`
// 	CertificateAuthorities          []string `json:"certificateAuthorities"`
// 	DisplayStatus                   string   `json:"displayStatus"`
// 	CreatedAt                       string   `json:"createdAt"`
// 	UpdatedAt                       string   `json:"updatedAt"`
// 	DefaultCertificateAuthority     *string  `json:"defaultCertificateAuthority"`
// 	ClientCertificateAuthentication *string  `json:"clientCertificateAuthentication"`
// 	ProgressStatus                  string   `json:"progressStatus"`
// }

// func (s *CreateListenerResponse) ToEntityListener() *lsentity.Listener {
// 	return &lsentity.Listener{
// 		UUID: s.UUID,
// 	}
// }

// func (s *ListListenersByLoadBalancerIdResponse) ToEntityListListeners() *lsentity.ListListeners {
// 	listeners := &lsentity.ListListeners{}
// 	for _, itemListener := range s.Data {
// 		listeners.Add(itemListener.toEntityListener())
// 	}
// 	return listeners
// }

// func (s *Listener) toEntityListener() *lsentity.Listener {
// 	return &lsentity.Listener{
// 		UUID:                            s.UUID,
// 		Name:                            s.Name,
// 		Description:                     s.Description,
// 		Protocol:                        s.Protocol,
// 		ProtocolPort:                    s.ProtocolPort,
// 		ConnectionLimit:                 s.ConnectionLimit,
// 		DefaultPoolId:                   s.DefaultPoolId,
// 		DefaultPoolName:                 s.DefaultPoolName,
// 		TimeoutClient:                   s.TimeoutClient,
// 		TimeoutMember:                   s.TimeoutMember,
// 		TimeoutConnection:               s.TimeoutConnection,
// 		AllowedCidrs:                    s.AllowedCidrs,
// 		Headers:                         s.Headers,
// 		CertificateAuthorities:          s.CertificateAuthorities,
// 		DisplayStatus:                   s.DisplayStatus,
// 		CreatedAt:                       s.CreatedAt,
// 		UpdatedAt:                       s.UpdatedAt,
// 		DefaultCertificateAuthority:     s.DefaultCertificateAuthority,
// 		ClientCertificateAuthentication: s.ClientCertificateAuthentication,
// 		ProgressStatus:                  s.ProgressStatus,
// 	}
// }

// func (s *GetListenerByIdResponse) ToEntityListener() *lsentity.Listener {
// 	return s.Data.toEntityListener()
// }
