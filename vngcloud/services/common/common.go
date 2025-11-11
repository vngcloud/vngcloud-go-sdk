package common

import (
	"slices"
	lstr "strings"
)

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

type Zone string

const (
	HCM_03_1A_ZONE     Zone = "HCM03-1A"
	HCM_03_1B_ZONE     Zone = "HCM03-1B"
	HCM_03_1C_ZONE     Zone = "HCM03-1C"
	HCM_03_BKK_01_ZONE Zone = "HCM03-BKK-01"
	HAN_01_1A_ZONE     Zone = "HAN01-1A"
)

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
	for _, agent := range pagent {
		if !slices.Contains(s.Agent, agent) {
			s.Agent = append(s.Agent, agent)
		}
	}
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
