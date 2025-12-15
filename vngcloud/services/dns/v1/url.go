package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getHostedZoneByIdUrl(psc lsclient.IServiceClient, popts IGetHostedZoneByIdRequest) string {
	return psc.ServiceURL(
		"dns",
		"hosted-zone",
		popts.GetHostedZoneId())
}
