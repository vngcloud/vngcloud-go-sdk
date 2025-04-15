package entity

type ListenerInsertHeader struct {
	HeaderName  string `json:"headerName"`
	HeaderValue string `json:"headerValue"`
}

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
	CertificateAuthorities          []string
	DisplayStatus                   string
	CreatedAt                       string
	UpdatedAt                       string
	DefaultCertificateAuthority     *string
	ClientCertificateAuthentication *string
	ProgressStatus                  string
	InsertHeaders                   []ListenerInsertHeader
}

type ListListeners struct {
	Items []*Listener
}

func (s *ListListeners) Add(listeners ...*Listener) {
	s.Items = append(s.Items, listeners...)
}

func (s *ListListeners) Len() int {
	return len(s.Items)
}

func (s *ListListeners) Empty() bool {
	return s.Len() < 1
}

func (s *Listener) GetId() string {
	return s.UUID
}

func (s *Listener) GetDefaultPoolId() string {
	return s.DefaultPoolId
}

func (s *ListListeners) At(index int) *Listener {
	if index < 0 || index >= len(s.Items) {
		return nil
	}

	return s.Items[index]
}
