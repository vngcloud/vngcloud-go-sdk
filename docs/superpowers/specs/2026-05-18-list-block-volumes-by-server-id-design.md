# Design — `ListBlockVolumesByServerId` for `VolumeServiceV2`

**Date:** 2026-05-18
**Status:** Approved for implementation planning
**Author:** tytv2

## Background

`vserver-api-spec.json` declares the endpoint:

| Verb | Path | Operation |
|------|------|-----------|
| `GET` | `/v2/{projectId}/volumes/servers/{serverId}` | `getVolumeByInstanceIdUsingGET` |

The spec summary is **"Get volume by server id"** (singular), but the response schema `VolumeResponse` is actually an **envelope** carrying a `volumes` array:

```json
{
  "errorCode": int,
  "errorMsg": string,
  "extra": object,
  "success": bool,
  "volumes": [Volume, ...]
}
```

So the endpoint returns **all volumes attached to a server**, not one. The spec wording is misleading.

The SDK currently has no method that calls this endpoint. `VolumeServiceV2` already has `ListBlockVolumes` (project-scoped list) and `GetBlockVolumeById` (single volume by volume ID), but no way to list a server's volumes.

## Goal

Add a single new method `ListBlockVolumesByServerId(serverId)` on `VolumeServiceV2` that calls this endpoint and returns `*lsentity.ListVolumes`.

## Non-goals

- Implementing the adjacent `/v2/{projectId}/volumes/servers/{serverId}/boot` ("Get boot volume by server id") endpoint — that's a separate scope.
- Refactoring or reusing `ListBlockVolumesResponse` — its envelope shape (`page`/`pageSize`/`listData`) is different from this endpoint's `VolumeResponse` shape (`volumes`).
- Parsing the envelope's `success`/`errorCode`/`errorMsg` fields — the SDK's HTTP/error layer already handles non-200s via `lserr.SdkErrorHandler`. Match existing pattern (`ListBlockVolumes`, `GetBlockVolumeById`) which ignores those fields on 200 responses.

## Architecture

### File layout

Append-only modifications to existing files in `vngcloud/services/volume/v2/`:

| File | Change |
|------|--------|
| `vngcloud/services/volume/v2/blockvolume.go` | append `ListBlockVolumesByServerId` service method |
| `vngcloud/services/volume/v2/blockvolume_request.go` | append `ListBlockVolumesByServerIdRequest` struct + `NewListBlockVolumesByServerIdRequest` constructor |
| `vngcloud/services/volume/v2/blockvolume_response.go` | append `ListBlockVolumesByServerIdResponse` struct + `ToEntityListVolumes` mapper |
| `vngcloud/services/volume/v2/irequest.go` | append `IListBlockVolumesByServerIdRequest` interface |
| `vngcloud/services/volume/v2/url.go` | append `listBlockVolumesByServerIdUrl` URL builder |
| `vngcloud/services/volume/ivolume.go` | 1 line registering the method in `IVolumeServiceV2` |
| `test/volume_test.go` | append `TestListBlockVolumesByServerIdSuccess` integration test |

No new files created.

### Files NOT modified

- All other services (`compute/`, `loadbalancer/`, `server/`, `network/`, etc.) — untouched.
- The `boot` variant endpoint — out of scope.

## Public API

```go
vngcloud.VServerGateway().V2().VolumeService().ListBlockVolumesByServerId(
    lsvolumeV2.NewListBlockVolumesByServerIdRequest(serverId))
// returns (*lsentity.ListVolumes, lserr.IError)
```

## Component design

### Request interface

```go
type IListBlockVolumesByServerIdRequest interface {
    GetServerId() string
    ParseUserAgent() string
    AddUserAgent(pagent ...string) IListBlockVolumesByServerIdRequest
    ToMap() map[string]interface{}
}
```

`ToMap()` is included to match the `WithParameters(popts.ToMap())` pattern used by `ListBlockVolumes` and other v2 methods for error attribution.

### Request struct + constructor

```go
type ListBlockVolumesByServerIdRequest struct {
    lscommon.UserAgent
    lscommon.ServerCommon  // exposes GetServerId()
}

func NewListBlockVolumesByServerIdRequest(pserverId string) IListBlockVolumesByServerIdRequest {
    opt := new(ListBlockVolumesByServerIdRequest)
    opt.ServerId = pserverId
    return opt
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

### Response

The endpoint's response envelope differs from `ListBlockVolumesResponse` (which is page-based), so a new struct is required. The array element type is the existing `BlockVolume` — reused for consistency.

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

The envelope's `errorCode`/`errorMsg`/`extra`/`success` fields are intentionally not parsed — non-2xx responses are handled by `lserr.SdkErrorHandler` via the `JsonError` registered on the request. This matches the pattern of `ListBlockVolumesResponse` and `GetBlockVolumeByIdResponse`.

### URL builder

```go
func listBlockVolumesByServerIdUrl(psc lsclient.IServiceClient, popts IListBlockVolumesByServerIdRequest) string {
    return psc.ServiceURL(
        psc.GetProjectId(),
        "volumes", "servers", popts.GetServerId())
}
```

Produces `/v2/{projectId}/volumes/servers/{serverId}` (the `v2` prefix and host come from `psc.ServiceURL`, matching every other endpoint in this package).

### Service method

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

### Interface registration

In `vngcloud/services/volume/ivolume.go`, add this line inside the `IVolumeServiceV2` interface block (after `MigrateBlockVolumeById`, before the tag methods or `// Snapshot` section):

```go
ListBlockVolumesByServerId(popts lsvolumeSvcV2.IListBlockVolumesByServerIdRequest) (*lsentity.ListVolumes, lserr.IError)
```

## Testing

Integration test matching existing SDK style (hits live API, requires credentials):

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

The hardcoded server ID matches the convention in [test/server_test.go:399](../../../test/server_test.go#L399) (developer replaces with a valid ID at runtime).

Build-time verification (TDD-via-compile gate as established in the prior tag-resource feature):
- `go build ./...` must exit 0
- `go vet ./vngcloud/...` must exit 0
- `go build ./test/...` must exit 0

No unit tests with HTTP mocks (SDK has no such infrastructure).

## Error handling

`lserr.NormalErrorType` with no extra error helpers, matching the pattern in `ListBlockVolumes` ([blockvolume.go:50-67](../../../vngcloud/services/volume/v2/blockvolume.go#L50-L67)). Attribution via `WithKVparameters("projectId", ..., "serverId", ...)` so error messages identify the offending server.

If the API returns a 4xx/5xx with a structured error body, `lserr.SdkErrorHandler` parses it from `errResp`. No new `WithError*` helpers added unless implementation reveals a server-specific error code worth surfacing.

## Edge cases

| Case | Behavior |
|------|----------|
| Server has no attached volumes | API returns `{ "volumes": [] }`; `ToEntityListVolumes` returns `&lsentity.ListVolumes{}` with empty `Items`. No panic. |
| Server does not exist | API returns 4xx; `SdkErrorHandler` surfaces it with `serverId` in error attribution. |
| Response envelope's `success: false` on a 200 | Not parsed (out of scope). Same behavior as `ListBlockVolumes` today. |

## Out of scope

- `/v2/{projectId}/volumes/servers/{serverId}/boot` endpoint.
- Pagination — this endpoint doesn't paginate.
- Filtering / query params — endpoint has none.
- Refactoring existing volume response types.
