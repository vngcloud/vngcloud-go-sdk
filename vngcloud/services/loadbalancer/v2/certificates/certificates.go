package certificates

import (
	"fmt"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Import(pSc *client.ServiceClient, pOpts IImportOptsBuilder) (*objects.Certificate, error) {
	response := &ImportResponse{}
	body := pOpts.ToRequestBody()
	res, err := pSc.Post(importURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{201},
	})

	fmt.Println(res)

	if err != nil {
		return nil, err
	}

	return response.ToCertificateObject(), nil
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Certificate, error) {
	response := &GetResponse{}
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToCertificateObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{204},
	})

	if err != nil {
		return err
	}

	return nil
}

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) ([]*objects.Certificate, error) {
	response := &ListResponse{}
	_, err := pSc.Get(listURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToListCertificateObjects(), nil
}
