package vcontainer

import (
	"strings"

	"github.com/cuongpiger/joat/utils"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	vconUtils "github.com/vngcloud/vngcloud-go-sdk/error/utils"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

func NewClient(pIdentityEndpoint string) (*client.ProviderClient, error) {
	provider := new(client.ProviderClient)
	provider.IdentityEndpoint = utils.NormalizeURL(pIdentityEndpoint)
	provider.UseTokenLock()
	provider.UseHTTPClient()

	return provider, nil
}

func Authenticate(pPc *client.ProviderClient, pAuthOpts tokens.AuthOptionsBuilder) error {
	authClient, err := newIdentity(pPc)
	if err != nil {
		return err
	}

	var accessToken string
	switch pAuthOpts.(type) {
	case *oauth2.AuthOptions:
		accessToken, err = oauth2.Create(authClient, pAuthOpts)
		if pAuthOpts.CanReauth() {
			tac := *pPc
			tac.SetThrowaway(true)
			tac.ReauthFunc = nil
			tac.SetAccessToken("")
			tao := pAuthOpts
			pPc.ReauthFunc = func() error {
				reauthAT, reauthErr := reauthForOAuth2(&tac, tao)
				if reauthErr != nil {
					return reauthErr
				}

				pPc.SetAccessToken(reauthAT)
				return nil
			}
		}
	}

	if err != nil {
		return err
	}

	pPc.SetAccessToken(accessToken)

	return nil
}

func NewServiceClient(pEndpoint string, pProvider *client.ProviderClient, pType string) (*client.ServiceClient, error) {
	return &client.ServiceClient{
		ProviderClient: pProvider,
		Endpoint:       utils.NormalizeURL(pEndpoint),
		Type:           strings.ToLower(pType),
	}, nil
}

func newIdentity(pProvider *client.ProviderClient) (*client.ServiceClient, error) {
	endpoint := utils.NormalizeURL(pProvider.IdentityEndpoint)
	isV2Endpoint := strings.HasSuffix(endpoint, "v2/")
	isV1Endpoint := strings.HasSuffix(endpoint, "v1/")

	if isV2Endpoint || isV1Endpoint {
		return &client.ServiceClient{
			ProviderClient: pProvider,
			Endpoint:       endpoint,
			Type:           "identity",
		}, nil
	}

	return nil, vconUtils.NewErrInvalidEndpoint("invalid identity endpoint")
}

func reauthForOAuth2(pProvider *client.ProviderClient, pAuthOpts tokens.AuthOptionsBuilder) (string, error) {
	authClient, err := newIdentity(pProvider)
	if err != nil {
		return "", err
	}

	accessToken, err := oauth2.Create(authClient, pAuthOpts)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
