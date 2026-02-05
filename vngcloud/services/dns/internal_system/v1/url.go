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

func listRecordsUrl(psc lsclient.IServiceClient, popts IListRecordsRequest) string {
	url := psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record")
	if popts.GetName() != "" {
		url += "?name=" + popts.GetName()
	}
	return url
}

func createHostedZoneUrl(psc lsclient.IServiceClient) string {
	return psc.ServiceURL("dns", "hosted-zone")
}

func deleteHostedZoneUrl(psc lsclient.IServiceClient, popts IDeleteHostedZoneRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId())
}

func updateHostedZoneUrl(psc lsclient.IServiceClient, popts IUpdateHostedZoneRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId())
}

func getRecordUrl(psc lsclient.IServiceClient, popts IGetRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func updateRecordUrl(psc lsclient.IServiceClient, popts IUpdateRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func deleteRecordUrl(psc lsclient.IServiceClient, popts IDeleteRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record", popts.GetRecordId())
}

func createDnsRecordUrl(psc lsclient.IServiceClient, popts ICreateDnsRecordRequest) string {
	return psc.ServiceURL("dns", "hosted-zone", popts.GetHostedZoneId(), "record")
}
