package snapshot

import lParser "github.com/cuongpiger/joat/parser"

type ListOpts struct {
	Page int `q:"page"`
	Size int `q:"size"`

	VolumeID string
	Status   string
	Name     string
}

func (s *ListOpts) ToListQuery() (string, error) {
	parser, _ := lParser.GetParser()
	url, err := parser.UrlMe(s)

	if err != nil {
		return "", err
	}

	return url.String(), err
}

func (s *ListOpts) ToListQueryWithParams(pParams *map[string]interface{}) (string, error) {
	if pParams != nil {
		if value, ok := (*pParams)["page"]; ok {
			s.Page = value.(int)
		}

		if value, ok := (*pParams)["size"]; ok {
			s.Size = value.(int)
		}

		parser, _ := lParser.GetParser()
		url, err := parser.UrlMe(s)
		if err != nil {
			return "", err
		}

		return url.String(), err
	}

	return "", nil
}

func (s *ListOpts) GetVolumeID() string {
	return s.VolumeID
}

func (s *ListOpts) GetStatus() string {
	return s.Status
}

func (s *ListOpts) GetName() string {
	return s.Name
}
