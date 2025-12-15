package v1

import lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"

func getHostedZoneByIdUrl(psc lsclient.IServiceClient, popts IGetHostedZoneByIdRequest) string {
	return psc.ServiceURL(
		"dns",
		"hosted-zone",
		popts.GetHostedZoneId())
}

func listHostedZonesUrl(psc lsclient.IServiceClient, popts IListHostedZonesRequest) string {
	url := psc.ServiceURL("dns", "hosted-zone")
	if popts.GetName() != "" {
		url += "?name=" + popts.GetName()
	}
	return url
}
