package volume

import (
	"fmt"
	"testing"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

func PointerOf[T any](t T) *T {
	return &t
}

var (
	projectID = "pro-462803f3-6858-466f-bf05-df2b33faa360"
	volID     = "vol-64b97016-373e-4292-bac0-8df9ebdfaf7e"
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vServerURL   = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2"
		clientID     = "cc13636c74480a"
		clientSecret = "f595636fabe28"
	)

	provider, _ := vngcloud.NewClient(identityURL)
	vngcloud.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vngcloud.NewServiceClient(
		vServerURL,
		provider,
		"vlb")
	return vlb
}

func TestGetVolume(t *testing.T) {
	vserverClient := NewSC()

	opt := NewGetOpts(projectID, volID)

	resp, err := Get(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestDeleteVolume(t *testing.T) {
	vserverClient := NewSC()

	opt := NewDeleteOpts(projectID, volID)

	err := Delete(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("Delete successfully")
}

func TestCreateVolumeFromSnapshot(t *testing.T) {
	vserverClient := NewSC()

	opt := NewCreateOpts(projectID, &CreateOpts{
		Name:         "test-volume",
		CreatedFrom:  CreateFromSnapshot,
		Size:         20,
		VolumeTypeId: "vtype-93a22a9f-1ec0-4e61-84fb-75ac181c13dc",
		ConfigureVolumeRestore: &ConfigureVolumeRestore{
			SnapshotVolumePointId: "snap-vol-pt-ee5103b1-c8d9-4c42-98f7-fe840ca808d3",
			VolumeTypeId:          "vtype-93a22a9f-1ec0-4e61-84fb-75ac181c13dc",
		},
	})

	vol, err := Create(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("Create successfully: %v", vol)
}
