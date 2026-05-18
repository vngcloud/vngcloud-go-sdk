# ListBlockVolumesByServerId Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `ListBlockVolumesByServerId(serverId)` method to `VolumeServiceV2`, calling `GET /v2/{projectId}/volumes/servers/{serverId}` and returning `*lsentity.ListVolumes`.

**Architecture:** Append-only modifications to existing files in `vngcloud/services/volume/v2/`. Reuse the existing `BlockVolume` struct as the response array element. Mirror the patterns of `ListBlockVolumes` / `GetBlockVolumeById` already in the package. No new files.

**Tech Stack:** Go 1.x, vngcloud-go-sdk client primitives (`lsclient.IServiceClient`, `lsentity.ListVolumes`, `lscommon.ServerCommon`).

**Spec:** [docs/superpowers/specs/2026-05-18-list-block-volumes-by-server-id-design.md](../specs/2026-05-18-list-block-volumes-by-server-id-design.md)

**Reference implementation:** [vngcloud/services/volume/v2/blockvolume.go:50-67 (`ListBlockVolumes`)](../../../vngcloud/services/volume/v2/blockvolume.go#L50-L67), [blockvolume_response.go:9-15, 67-74 (`ListBlockVolumesResponse`)](../../../vngcloud/services/volume/v2/blockvolume_response.go)

---

## Conventions

- **TDD via compile-check:** SDK has no unit-test infrastructure; integration tests in `test/` need real credentials. "Failing test" gate = `go build ./test/...` returning non-zero because the integration test references symbols that don't exist yet. "Test passes" gate = `go build ./...`, `go vet ./vngcloud/...`, and `go build ./test/...` all exit 0.
- **Running the integration test** is the developer's manual step after the SDK builds — requires real `vngcloud` credentials and a valid server ID. Not part of plan execution.
- **Test fixture ID:** `ins-da59addd-6263-4544-b405-420a65ccfb1f` (same hardcoded ID used in `TestCreateSystemTags` at [test/server_test.go:399](../../../test/server_test.go#L399)).
- **Commit message format:** `feat: <short summary>` to match project history.

---

## Task 1: `ListBlockVolumesByServerId`

**Files:**
- Modify: `vngcloud/services/volume/v2/blockvolume.go` (append service method)
- Modify: `vngcloud/services/volume/v2/blockvolume_request.go` (append request struct + constructor)
- Modify: `vngcloud/services/volume/v2/blockvolume_response.go` (append response struct + mapper)
- Modify: `vngcloud/services/volume/v2/irequest.go` (append interface)
- Modify: `vngcloud/services/volume/v2/url.go` (append URL builder)
- Modify: `vngcloud/services/volume/ivolume.go` (1 line in `IVolumeServiceV2` block)
- Test: `test/volume_test.go` (append integration test)

- [ ] **Step 1: Write the failing integration test**

Append to `test/volume_test.go`:

```go
func TestListBlockVolumesByServerIdSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := v2.NewListBlockVolumesByServerIdRequest("ins-da59addd-6263-4544-b405-420a65ccfb1f")
	volumes, sdkErr := vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumesByServerId(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	if volumes == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", volumes)
	t.Log("PASS")
}
```

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewListBlockVolumesByServerIdRequest` or `ListBlockVolumesByServerId` not declared in `v2` package / on `IVolumeServiceV2`.

- [ ] **Step 3: Add the request interface**

Append to `vngcloud/services/volume/v2/irequest.go`:

```go

type IListBlockVolumesByServerIdRequest interface {
	GetServerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListBlockVolumesByServerIdRequest
	ToMap() map[string]interface{}
}
```

- [ ] **Step 4: Add the request struct + constructor**

Append to `vngcloud/services/volume/v2/blockvolume_request.go`:

```go

func NewListBlockVolumesByServerIdRequest(pserverId string) IListBlockVolumesByServerIdRequest {
	opt := new(ListBlockVolumesByServerIdRequest)
	opt.ServerId = pserverId
	return opt
}

type ListBlockVolumesByServerIdRequest struct {
	lscommon.UserAgent
	lscommon.ServerCommon
}

func (s *ListBlockVolumesByServerIdRequest) AddUserAgent(pagent ...string) IListBlockVolumesByServerIdRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *ListBlockVolumesByServerIdRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"serverId": s.ServerId,
	}
}
```

Note: `lscommon.UserAgent` and `lscommon.ServerCommon` are already used elsewhere in this file via embedded types; the existing import `lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"` covers it. If for some reason `lscommon` is not yet imported, add it.

- [ ] **Step 5: Add the response struct + entity mapper**

Append to `vngcloud/services/volume/v2/blockvolume_response.go`:

```go

type ListBlockVolumesByServerIdResponse struct {
	Volumes []BlockVolume `json:"volumes"`
}

func (s *ListBlockVolumesByServerIdResponse) ToEntityListVolumes() *lsentity.ListVolumes {
	lst := new(lsentity.ListVolumes)
	for _, v := range s.Volumes {
		lst.Items = append(lst.Items, v.toEntityVolume())
	}
	return lst
}
```

- [ ] **Step 6: Add the URL builder**

Append to `vngcloud/services/volume/v2/url.go`:

```go

func listBlockVolumesByServerIdUrl(psc lsclient.IServiceClient, popts IListBlockVolumesByServerIdRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"volumes", "servers", popts.GetServerId())
}
```

- [ ] **Step 7: Add the service method**

Append to `vngcloud/services/volume/v2/blockvolume.go`:

```go

func (s *VolumeServiceV2) ListBlockVolumesByServerId(popts IListBlockVolumesByServerIdRequest) (*lsentity.ListVolumes, lserr.IError) {
	url := listBlockVolumesByServerIdUrl(s.VServerClient, popts)
	resp := new(ListBlockVolumesByServerIdResponse)
	errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
	req := lsclient.NewRequest().
		WithHeader("User-Agent", popts.ParseUserAgent()).
		WithOkCodes(200).
		WithJsonResponse(resp).
		WithJsonError(errResp)

	if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
		return nil, lserr.SdkErrorHandler(sdkErr, errResp).
			WithKVparameters(
				"projectId", s.getProjectId(),
				"serverId", popts.GetServerId())
	}

	return resp.ToEntityListVolumes(), nil
}
```

- [ ] **Step 8: Register the method on the public interface**

In `vngcloud/services/volume/ivolume.go`, add this line inside the `IVolumeServiceV2` interface block, immediately after `MigrateBlockVolumeById` (before the existing tag methods or the `// Snapshot` section comment):

```go
	ListBlockVolumesByServerId(popts lsvolumeSvcV2.IListBlockVolumesByServerIdRequest) (*lsentity.ListVolumes, lserr.IError)
```

- [ ] **Step 9: Verify the build passes**

Run: `go build ./...`
Expected: exit 0, no errors.

Run: `go vet ./vngcloud/...`
Expected: exit 0, no warnings.

Run: `go build ./test/...`
Expected: exit 0 (test file compiles).

- [ ] **Step 10: Commit**

```bash
git add vngcloud/services/volume/v2/blockvolume.go \
        vngcloud/services/volume/v2/blockvolume_request.go \
        vngcloud/services/volume/v2/blockvolume_response.go \
        vngcloud/services/volume/v2/irequest.go \
        vngcloud/services/volume/v2/url.go \
        vngcloud/services/volume/ivolume.go \
        test/volume_test.go
git commit -m "feat: add ListBlockVolumesByServerId to VolumeServiceV2"
```

(Use `git -c commit.gpgsign=false commit ...` if signing is configured.)

---

## Final verification

- [ ] **Run full build + vet one last time**

Run: `go build ./...` → exit 0
Run: `go vet ./vngcloud/...` → exit 0
Run: `go build ./test/...` → exit 0

- [ ] **Optional: run the integration test (requires real credentials and a valid server ID)**

Replace the hardcoded server ID in the test with a real one and run:

```
go test -v -run TestListBlockVolumesByServerIdSuccess ./test/...
```

Expected: PASS, logged volume list. Failures here indicate live-environment issues, not code defects.
