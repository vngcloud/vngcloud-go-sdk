package pagination

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

type Pager struct {
	client         *client.ServiceClient
	initialURL     string
	createPage     func(interface{}) IPage
	firstPage      IPage
	createResponse func() interface{}
	Err            error
	Headers        map[string]string
	Response       interface{}
	PageOpts       IPageOpts
}

func (s *Pager) EachPage(pHandler func(IPage) (bool, error)) error {
	if s.Err != nil {
		return s.Err
	}

	curUrl := s.initialURL
	for {
		var curPage IPage
		if s.firstPage != nil {
			curPage = s.firstPage
			s.firstPage = nil
		} else {
			var err error
			curPage, err = s.fetchNextPage(curUrl)
			if err != nil {
				return err
			}

			empty, err := curPage.IsEmpty()
			if err != nil {
				return err
			}

			if empty {
				return nil
			}

			ok, err := pHandler(curPage)
			if err != nil {
				return err
			}

			if !ok {
				return nil
			}

			curUrl, err = curPage.NextPageURL(s.PageOpts)
			if err != nil {
				return err
			}

			if curUrl == "" {
				return nil
			}
		}
	}
}

func (s *Pager) fetchNextPage(pUrl string) (IPage, error) {
	s.Response = s.createResponse()
	_, err := s.client.Get(pUrl, &client.RequestOpts{
		MoreHeaders:  s.Headers,
		OkCodes:      []int{200, 204, 300},
		JSONResponse: s.Response,
	})
	if err != nil {
		return nil, err
	}

	return s.createPage(s.Response), nil
}
