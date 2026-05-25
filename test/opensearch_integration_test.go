package test

import (
	lctx "context"
	ltesting "testing"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lsopensearchV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/opensearch/v1"
)

const envVdbOpenSearchEndpoint = "https://vdb.console.vngcloud.vn"

// validOpenSearchSdkConfig builds an SDK client wired to IAM + vDB OpenSearch
// endpoints. Distinct from validKafkaSdkConfig because the OS gateway also
// needs projectId (paths embed it).
func validOpenSearchSdkConfig(t *ltesting.T) lsclient.IClient {
	t.Helper()
	clientId, clientSecret := getEnv()
	if clientId == "" || clientSecret == "" {
		t.Skip("env.yaml missing VNGCLOUD_CLIENT_ID / VNGCLOUD_CLIENT_SECRET")
	}
	portalUserId := getValueOfEnv("VNGCLOUD_USER_ID")
	projectId := getValueOfEnv("VNGCLOUD_PROJECT_ID")
	if portalUserId == "" || projectId == "" {
		t.Skip("env.yaml missing VNGCLOUD_USER_ID / VNGCLOUD_PROJECT_ID")
	}

	sdkCfg := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(portalUserId).
		WithProjectId(projectId).
		WithIamEndpoint(envIamEndpoint).
		WithVdbOpenSearchEndpoint(envVdbOpenSearchEndpoint)

	return lsclient.NewClient(lctx.TODO()).
		WithRetryCount(1).
		WithSleep(2 * ltime.Second).
		Configure(sdkCfg)
}

func TestOpenSearch_ListClusters(t *ltesting.T) {
	c := validOpenSearchSdkConfig(t)
	svc := c.VDBOpenSearchGateway().V1().OpenSearchService()

	clusters, err := svc.ListClusters(lsopensearchV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil {
		t.Fatal("expected non-nil response")
	}
	t.Logf("found %d open-search cluster(s)", len(*clusters))
	for i, cl := range *clusters {
		t.Logf("  [%d] id=%s name=%q status=%s version=%s nodes=%d tls=%v public=%v private=%v region=%s",
			i, cl.Id, cl.Name, cl.Status, cl.Version, cl.NumberOfNodes,
			cl.EnableTls, cl.PublicAccess, cl.PrivateAccess, cl.Region)
		t.Logf("      package=%s ram=%dGB cpu=%d storage=%dGB type=%s",
			cl.PackageDetail.Name, cl.PackageDetail.Ram, cl.PackageDetail.Cpu,
			cl.StorageSize, cl.StorageType.Name)
	}
}

func TestOpenSearch_ListEndpoints(t *ltesting.T) {
	c := validOpenSearchSdkConfig(t)
	svc := c.VDBOpenSearchGateway().V1().OpenSearchService()

	clusters, err := svc.ListClusters(lsopensearchV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no clusters")
	}
	clusterId := (*clusters)[0].Id

	eps, eerr := svc.ListEndpoints(lsopensearchV1.NewListEndpointsRequest(clusterId))
	if eerr != nil {
		t.Fatalf("ListEndpoints(%s) failed: %+v", clusterId, eerr.GetError())
	}
	t.Logf("cluster %s has %d endpoint(s):", clusterId, len(*eps))
	for i, ep := range *eps {
		t.Logf("  [%d] type=%s url=%s private=%v cidrs=%v", i, ep.Type, ep.Url, ep.Private, ep.Cidrs)
	}

	// Sanity: there should be at least one log-type endpoint usable by workers
	var logEp string
	for _, ep := range *eps {
		if ep.IsLogIngest() {
			logEp = ep.Url
			break
		}
	}
	if logEp == "" {
		t.Errorf("no log-type endpoint found — worker bulk write target unknown")
	} else {
		t.Logf("worker would use log endpoint: %s", logEp)
	}
}

func TestOpenSearch_GetCluster(t *ltesting.T) {
	c := validOpenSearchSdkConfig(t)
	svc := c.VDBOpenSearchGateway().V1().OpenSearchService()

	clusters, err := svc.ListClusters(lsopensearchV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no open-search clusters — skipping GetCluster")
	}
	want := (*clusters)[0]

	got, gerr := svc.GetCluster(lsopensearchV1.NewGetClusterRequest(want.Id))
	if gerr != nil {
		t.Fatalf("GetCluster(%s) failed: %+v", want.Id, gerr.GetError())
	}
	if got.Id != want.Id {
		t.Errorf("id mismatch: got=%s want=%s", got.Id, want.Id)
	}
	t.Logf("cluster %s: name=%q status=%s version=%s nodes=%d storage=%dGB tls=%v public=%v",
		got.Id, got.Name, got.Status, got.Version, got.NumberOfNodes,
		got.StorageSize, got.EnableTls, got.PublicAccess)
	t.Logf("  vpc=%s/%s subnet=%s vpcDns=%v encrypt=%v",
		got.VpcId, got.VpcCidr, got.SubnetId, got.EnableVpcDns, got.EncryptVolume)
	t.Logf("  package=%s (%dGB ram, %d cpu, %s) storage=%s billing=%s",
		got.PackageDetail.Name, got.PackageDetail.Ram, got.PackageDetail.Cpu,
		got.PackageDetail.NetworkPerformance, got.StorageType.Name, got.BillingStatus.Status)
}
