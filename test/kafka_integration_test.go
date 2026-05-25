package test

import (
	larchivezip "archive/zip"
	lbytes "bytes"
	lctx "context"
	lfmt "fmt"
	los "os"
	lstrings "strings"
	ltesting "testing"
	ltime "time"

	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
	lskafkaV1 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/kafka/v1"
)

const (
	envIamEndpoint      = "https://iamapis.vngcloud.vn/accounts-api"
	envVdbKafkaEndpoint = "https://vdb-gateway.vngcloud.vn"

	// Opt-in toggle to run the destructive CreateTopic test against real vDB.
	// Set VKS_TEST_CREATE_TOPIC=1 to enable. Skipped by default.
	envCreateTopicToggle = "VKS_TEST_CREATE_TOPIC"
)

// validKafkaSdkConfig builds an SDK client wired to IAM + vDB Kafka endpoints
// using credentials in test/env.yaml. Other endpoints are unset — these tests
// only touch the kafka gateway.
func validKafkaSdkConfig(t *ltesting.T) lsclient.IClient {
	t.Helper()
	clientId, clientSecret := getEnv()
	if clientId == "" || clientSecret == "" {
		t.Skip("env.yaml missing VNGCLOUD_CLIENT_ID / VNGCLOUD_CLIENT_SECRET — skipping integration test")
	}

	portalUserId := getValueOfEnv("VNGCLOUD_USER_ID")
	if portalUserId == "" {
		t.Skip("env.yaml missing VNGCLOUD_USER_ID — skipping integration test")
	}

	sdkCfg := lsclient.NewSdkConfigure().
		WithClientId(clientId).
		WithClientSecret(clientSecret).
		WithUserId(portalUserId).
		WithIamEndpoint(envIamEndpoint).
		WithVdbKafkaEndpoint(envVdbKafkaEndpoint)

	return lsclient.NewClient(lctx.TODO()).
		WithRetryCount(1).
		WithSleep(2 * ltime.Second).
		Configure(sdkCfg)
}

func TestKafka_ListClusters(t *ltesting.T) {
	c := validKafkaSdkConfig(t)
	svc := c.VDBKafkaGateway().V1().KafkaService()

	clusters, err := svc.ListClusters(lskafkaV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil {
		t.Fatal("expected non-nil response (may be empty slice)")
	}

	t.Logf("found %d kafka cluster(s) for portal user", len(*clusters))
	for i, cl := range *clusters {
		t.Logf("  [%d] id=%s name=%q status=%s mtls=%v sasl=%v public=%v floatingIps=%v",
			i, cl.Id, cl.Name, cl.Status, cl.MtlsAuthen, cl.SaslAuthen, cl.PublicAccess, cl.FloatingIps)
	}
}

func TestKafka_GetCluster(t *ltesting.T) {
	c := validKafkaSdkConfig(t)
	svc := c.VDBKafkaGateway().V1().KafkaService()

	clusters, err := svc.ListClusters(lskafkaV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no kafka clusters in account — skipping GetCluster")
	}

	clusterId := (*clusters)[0].Id
	got, gerr := svc.GetCluster(lskafkaV1.NewGetClusterRequest(clusterId))
	if gerr != nil {
		t.Fatalf("GetCluster(%s) failed: %+v", clusterId, gerr.GetError())
	}
	if got.Id != clusterId {
		t.Errorf("cluster id mismatch: got %s, want %s", got.Id, clusterId)
	}
	t.Logf("cluster %s: name=%q status=%s brokers=%d mtls=%v sasl=%v public=%v",
		got.Id, got.Name, got.Status, got.KafkaBrokerCount, got.MtlsAuthen, got.SaslAuthen, got.PublicAccess)
	t.Logf("  fixedIps=%v floatingIps=%v", got.FixedIps, got.FloatingIps)
}

func TestKafka_ListUsers(t *ltesting.T) {
	c := validKafkaSdkConfig(t)
	svc := c.VDBKafkaGateway().V1().KafkaService()

	clusters, err := svc.ListClusters(lskafkaV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no kafka clusters in account")
	}
	clusterId := (*clusters)[0].Id

	users, uerr := svc.ListUsers(lskafkaV1.NewListUsersRequest(clusterId))
	if uerr != nil {
		t.Fatalf("ListUsers(%s) failed: %+v", clusterId, uerr.GetError())
	}
	t.Logf("found %d user(s) on cluster %s", len(*users), clusterId)
	for i, u := range *users {
		t.Logf("  [%d] id=%s name=%q mtls=%v sasl=%v produceAll=%v produceTopics=%v",
			i, u.Id, u.Name, u.MtlsAuthen, u.SaslAuthen, u.ProduceAll, u.ProduceTopicNames)
	}
}

func TestKafka_ListTopics(t *ltesting.T) {
	c := validKafkaSdkConfig(t)
	svc := c.VDBKafkaGateway().V1().KafkaService()

	clusters, err := svc.ListClusters(lskafkaV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no kafka clusters in account")
	}
	clusterId := (*clusters)[0].Id

	topics, terr := svc.ListTopics(lskafkaV1.NewListTopicsRequest(clusterId))
	if terr != nil {
		t.Fatalf("ListTopics(%s) failed: %+v", clusterId, terr.GetError())
	}
	t.Logf("found %d topic(s) on cluster %s", len(*topics), clusterId)
	for i, tp := range *topics {
		t.Logf("  [%d] name=%q partitions=%d replicas=%d retentionSec=%d status=%s",
			i, tp.Name, tp.Partitions, tp.Replicas, tp.RetentionSeconds, tp.Status)
	}
}

func TestKafka_GetUserAuthenCreds(t *ltesting.T) {
	c := validKafkaSdkConfig(t)
	svc := c.VDBKafkaGateway().V1().KafkaService()

	clusters, err := svc.ListClusters(lskafkaV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no kafka clusters in account")
	}

	// Find a cluster that has at least one user (downloading creds requires a user id)
	var clusterId, userId string
	for _, cl := range *clusters {
		if cl.Status != "ACTIVE" {
			continue
		}
		users, uerr := svc.ListUsers(lskafkaV1.NewListUsersRequest(cl.Id))
		if uerr != nil || users == nil || len(*users) == 0 {
			continue
		}
		clusterId = cl.Id
		userId = (*users)[0].Id
		t.Logf("using cluster=%s user=%s for authen-creds test", clusterId, userId)
		break
	}
	if clusterId == "" {
		t.Skip("no active cluster with a user — skipping authen-creds test")
	}

	zipBytes, zerr := svc.GetUserAuthenCreds(lskafkaV1.NewGetUserAuthenCredsRequest(clusterId, userId))
	if zerr != nil {
		t.Fatalf("GetUserAuthenCreds(%s, %s) failed: %+v", clusterId, userId, zerr.GetError())
	}
	if len(zipBytes) == 0 {
		t.Fatal("expected non-empty zip bytes")
	}
	t.Logf("downloaded authen-creds zip: %d bytes", len(zipBytes))

	// Verify it's actually a zip
	zr, zererr := larchivezip.NewReader(lbytes.NewReader(zipBytes), int64(len(zipBytes)))
	if zererr != nil {
		t.Fatalf("response is not a valid zip: %v", zererr)
	}
	t.Logf("zip contains %d entries:", len(zr.File))
	for _, f := range zr.File {
		t.Logf("  %s (%d bytes)", f.Name, f.UncompressedSize64)
	}

	// Parse via SDK helper
	creds, perr := lskafkaV1.ParseAuthenCredsZip(zipBytes)
	if perr != nil {
		t.Fatalf("ParseAuthenCredsZip failed: %v", perr)
	}
	t.Logf("parsed userId=%s hasMTLS=%v hasSASL=%v", creds.UserId, creds.HasMTLS(), creds.HasSASL())
	if creds.HasMTLS() {
		t.Logf("  mtls ca.crt=%d bytes, user.crt=%d bytes, user.key=%d bytes",
			len(creds.MTLS.CACert), len(creds.MTLS.UserCert), len(creds.MTLS.UserKey))
	}
	if creds.HasSASL() {
		t.Logf("  sasl username=%s mechanism=%s password=%d chars",
			creds.SASL.Username, creds.SASL.Mechanism, len(creds.SASL.Password))
	}

	// Build bootstrap servers from cluster info
	cluster, gerr := svc.GetCluster(lskafkaV1.NewGetClusterRequest(clusterId))
	if gerr != nil {
		t.Fatalf("GetCluster failed: %+v", gerr.GetError())
	}
	if cluster.PublicAccess && len(cluster.FloatingIps) > 0 {
		if cluster.MtlsAuthen {
			if brokers, berr := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthMTLS); berr != nil {
				t.Errorf("BuildBootstrapServers(mTLS) failed: %v", berr)
			} else {
				t.Logf("  bootstrap mTLS: %s", brokers)
			}
		}
		if cluster.SaslAuthen {
			if brokers, berr := lskafkaV1.BuildBootstrapServers(cluster, lskafkaV1.AuthSASL); berr != nil {
				t.Errorf("BuildBootstrapServers(SASL) failed: %v", berr)
			} else {
				t.Logf("  bootstrap SASL: %s", brokers)
			}
		}
	} else {
		t.Logf("  cluster public access disabled — skipping BuildBootstrapServers (V1 only supports public)")
	}
}

// TestKafka_CreateTopic actually mutates the cluster. Opt-in only.
// Run with:  VKS_TEST_CREATE_TOPIC=1 go test ...
// Optional VKS_TEST_CLUSTER_NAME=<name> picks a specific cluster (else uses index 0).
func TestKafka_CreateTopic(t *ltesting.T) {
	if los.Getenv(envCreateTopicToggle) != "1" {
		t.Skipf("destructive — set %s=1 to enable", envCreateTopicToggle)
	}
	c := validKafkaSdkConfig(t)
	svc := c.VDBKafkaGateway().V1().KafkaService()

	clusters, err := svc.ListClusters(lskafkaV1.NewListClustersRequest())
	if err != nil {
		t.Fatalf("ListClusters failed: %+v", err.GetError())
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Skip("no kafka clusters in account")
	}

	wantName := los.Getenv("VKS_TEST_CLUSTER_NAME")
	var clusterId string
	if wantName != "" {
		for _, cl := range *clusters {
			if cl.Name == wantName {
				clusterId = cl.Id
				t.Logf("selected cluster by name: %s (id=%s)", cl.Name, cl.Id)
				break
			}
		}
		if clusterId == "" {
			t.Fatalf("no cluster with name=%q found", wantName)
		}
	} else {
		clusterId = (*clusters)[0].Id
		t.Logf("using cluster index 0: %s", clusterId)
	}

	topicName := lfmt.Sprintf("vks-sdk-test-%d", ltime.Now().Unix())
	topic, terr := svc.CreateTopic(lskafkaV1.NewCreateTopicRequest(clusterId).
		WithName(topicName).
		WithPartitions(1).
		WithReplicas(1))
	if terr != nil {
		// Many vDB clusters have replicas constraint; surface error verbatim
		t.Fatalf("CreateTopic(%s) failed: %+v", topicName, terr.GetError())
	}
	t.Logf("created topic: name=%s id=%s status=%s partitions=%d replicas=%d retentionSec=%d",
		topic.Name, topic.Id, topic.Status, topic.Partitions, topic.Replicas, topic.RetentionSeconds)

	if !lstrings.Contains(topic.Name, "vks-sdk-test-") {
		t.Errorf("topic name not preserved: got %s", topic.Name)
	}

	// Verify via ListTopics
	topics, lerr := svc.ListTopics(lskafkaV1.NewListTopicsRequest(clusterId))
	if lerr != nil {
		t.Errorf("ListTopics after create failed: %+v", lerr.GetError())
		return
	}
	found := false
	for _, tp := range *topics {
		if tp.Name == topicName {
			found = true
			t.Logf("verified via ListTopics: found %s id=%s status=%s", tp.Name, tp.Id, tp.Status)
			break
		}
	}
	if !found {
		t.Errorf("created topic %s not found in ListTopics", topicName)
	}
}
