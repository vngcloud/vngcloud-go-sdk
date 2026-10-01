package v2

import (
	"encoding/json"
	"testing"
)

func TestCreateListenerRequestCarriesAcl(t *testing.T) {
	req := NewCreateListenerRequest("listener_01", ListenerProtocolHTTP, 80).
		WithBlockedCidrs("203.0.113.9/32", "192.0.2.0/24").
		WithDefaultAction(ListenerDefaultActionDrop)
	body, _ := json.Marshal(req.ToRequestBody())
	var got map[string]any
	_ = json.Unmarshal(body, &got)
	if got["blockedCidrs"] != "203.0.113.9/32,192.0.2.0/24" || got["defaultAction"] != "drop" {
		t.Fatalf("acl not in create body: %s", body)
	}
}

func TestCreateListenerRequestOmitsUnsetAcl(t *testing.T) {
	body, _ := json.Marshal(NewCreateListenerRequest("listener_01", ListenerProtocolHTTP, 80).ToRequestBody())
	var got map[string]any
	_ = json.Unmarshal(body, &got)
	if _, ok := got["blockedCidrs"]; ok {
		t.Fatalf("unset blockedCidrs must be omitted: %s", body)
	}
	if _, ok := got["defaultAction"]; ok {
		t.Fatalf("unset defaultAction must be omitted: %s", body)
	}
}

// An empty, non-nil pointer is how a caller clears the list, so it must reach the wire.
func TestUpdateListenerRequestSendsAnEmptyBlockedListButOmitsNil(t *testing.T) {
	empty := ""
	body, _ := json.Marshal((&UpdateListenerRequest{BlockedCidrs: &empty}).ToRequestBody())
	var got map[string]any
	_ = json.Unmarshal(body, &got)
	if v, ok := got["blockedCidrs"]; !ok || v != "" {
		t.Fatalf("empty blockedCidrs must be sent: %s", body)
	}
	if _, ok := got["defaultAction"]; ok {
		t.Fatalf("nil defaultAction must be omitted: %s", body)
	}
}

func TestListenerResponseMapsAcl(t *testing.T) {
	var resp GetListenerByIdResponse
	_ = json.Unmarshal([]byte(`{"data":{"uuid":"lis-1","blockedCidrs":"203.0.113.9/32","defaultAction":"drop"}}`), &resp)
	l := resp.Data.toEntityListener()
	if l.BlockedCidrs != "203.0.113.9/32" || l.DefaultAction != "drop" {
		t.Fatalf("acl not mapped: %+v", l)
	}
}
