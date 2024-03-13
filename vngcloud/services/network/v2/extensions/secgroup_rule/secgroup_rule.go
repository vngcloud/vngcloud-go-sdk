package secgroup_rule

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*objects.SecgroupRule, error) {
	errorResolver := new(ErrorResolver)
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		JSONError:    errorResolver,
		OkCodes:      []int{201},
	})

	if err != nil {
		return nil, errorResolver.ToError()
	}

	return response.ToSecgroupRuleObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{204}},
	)

	return err
}

func ListRulesBySecgroupID(pSc *client.ServiceClient, pOpts IListRulesBySecgroupIDOptsBuilder) ([]*objects.SecgroupRule, error) {
	response := NewListRulesBySecgroupIDResponse()
	_, err := pSc.Get(listRulesURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListSecgroupRules(), nil
}
