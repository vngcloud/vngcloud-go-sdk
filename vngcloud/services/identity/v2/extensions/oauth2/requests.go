package oauth2

import (
	"encoding/base64"
	"fmt"
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

type AuthOptions struct {
	ClientID     string
	ClientSecret string

	tokens.AuthOptionsBuilder
}

type (
	GetAccessTokenRequestBody struct {
		GrantType string `json:"grant_type"`
	}

	GetAccessTokenResponse struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		TokenType        string `json:"token_type"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
	}
)

func (s *AuthOptions) ToTokenV2BodyMap(interface{}) (interface{}, error) {
	return &GetAccessTokenRequestBody{
		GrantType: "client_credentials",
	}, nil
}

func (s *AuthOptions) ToTokenV2HeadersMap(map[string]interface{}) (map[string]string, error) {
	return map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(s.ClientID+":"+s.ClientSecret)),
	}, nil
}

func (s *AuthOptions) CanReauth() bool {
	return true
}

func Create(pSc *client.ServiceClient, pOpts tokens.AuthOptionsBuilder) (string, error) {
	var result GetAccessTokenResponse
	body, _ := pOpts.ToTokenV2BodyMap(nil)
	headers, _ := pOpts.ToTokenV2HeadersMap(nil)

	resp, err := pSc.Post(authURL(pSc), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: &result,
		MoreHeaders:  headers,
		OkCodes:      []int{200},
	})

	if err != nil {
		fmt.Println("The failed response is: ", resp)
		return "", err
	}

	return result.AccessToken, nil
}
