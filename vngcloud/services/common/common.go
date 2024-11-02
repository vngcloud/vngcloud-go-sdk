package common

import lstr "strings"

type Project struct {
	Id string
}

func (s *Project) GetPorjectId() string {
	return s.Id
}

func (s *Project) SetProjectId(pid string) {
	s.Id = pid
}

type Paging struct {
	Page int
	Size int
}

func (s *Paging) GetPage() int {
	return s.Page
}

func (s *Paging) GetSize() int {
	return s.Size
}

func (s *Paging) SetPage(ppage int) *Paging {
	s.Page = ppage
	return s
}

func (s *Paging) SetSize(psize int) *Paging {
	s.Size = psize
	return s
}

type Tag struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	IsEdited bool   `json:"isEdited,omitempty"`
}

type UserAgent struct {
	Agent []string
}

func (s *UserAgent) ParseUserAgent() string {
	// Parse the array into string
	return lstr.Join(s.Agent, "; ")
}

func (s *UserAgent) AddUserAgent(pagent ...string) *UserAgent {
	s.Agent = append(s.Agent, pagent...)
	return s
}

type PortalUser struct {
	Id string
}

func (s *PortalUser) GetPortalUserId() string {
	return s.Id
}

func (s *PortalUser) SetPortalUserId(pid string) {
	s.Id = pid
}

func (s *PortalUser) GetMapHeaders() map[string]string {
	return map[string]string{
		"portal-user-id": s.Id,
	}
}
