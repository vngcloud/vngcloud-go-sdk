# Design — Add `tag/resource/{resourceId}` API support to Compute (Server) and Volume services

**Date:** 2026-05-17
**Status:** Approved for implementation planning
**Author:** tytv2

## Background

`vserver-api-spec.json` declares two endpoints for managing tags on a resource:

| Verb | Path | Operation |
|------|------|-----------|
| `GET` | `/v2/{projectId}/tag/resource/{resourceId}` | List tags of a resource |
| `PUT` | `/v2/{projectId}/tag/resource/{resourceId}` | Replace tag list on a resource (`resourceType` in body distinguishes Server / Volume / LoadBalancer) |

The SDK currently implements these endpoints **only for LoadBalancer** ([`vngcloud/services/loadbalancer/v2/tag.go`](../../../vngcloud/services/loadbalancer/v2/tag.go), [`url.go`](../../../vngcloud/services/loadbalancer/v2/url.go)). Server and Volume have no equivalents:

- Server v1 has `ServerServiceInternalV1.CreateSystemTags` ([`systemtag.go`](../../../vngcloud/services/server/v1/systemtag.go)) — calls a different endpoint, `POST /<projectId>/tags`, not the spec endpoint above.
- Volume only supports inline tags on `CreateBlockVolume` via `WithTags` ([`blockvolume_request.go:250`](../../../vngcloud/services/volume/v2/blockvolume_request.go#L250)) — no per-resource tag list management.

## Goal

Add three methods — `ListTags`, `CreateTags`, `UpdateTags` — to both `ComputeServiceV2` (which manages Server resources) and `VolumeServiceV2`, mirroring the LoadBalancer pattern exactly.

## Non-goals

- Refactoring the existing LoadBalancer tag implementation.
- Touching `ServerServiceInternalV1.CreateSystemTags` (legacy endpoint, kept as-is).
- Adding a `DeleteTags` method (no DELETE endpoint exists for `/tag/resource/`).
- Introducing a generic shared tag service or refactoring tag code across services.

## Architecture

### File layout

```
vngcloud/services/compute/v2/
├── tag.go                # 3 service methods on ComputeServiceV2
├── tag_request.go        # request structs + builders
└── tag_response.go       # response struct + entity mapping

vngcloud/services/volume/v2/
├── tag.go                # 3 service methods on VolumeServiceV2
├── tag_request.go
└── tag_response.go
```

### Files to modify

- `vngcloud/services/compute/icompute.go` — add 3 methods to `IComputeServiceV2`.
- `vngcloud/services/compute/v2/irequest.go` — add `IListTagsRequest`, `ICreateTagsRequest`, `IUpdateTagsRequest`.
- `vngcloud/services/compute/v2/url.go` — add `listTagsUrl`, `createTagsUrl`, `updateTagsUrl`.
- `vngcloud/services/volume/ivolume.go` — add 3 methods to `IVolumeServiceV2`.
- `vngcloud/services/volume/v2/irequest.go` — add 3 tag request interfaces.
- `vngcloud/services/volume/v2/url.go` — add 3 URL builders.

### Files NOT modified

- `vngcloud/services/loadbalancer/v2/*` — LB tag code stays as-is.
- `vngcloud/services/server/v1/systemtag*` — legacy `CreateSystemTags` stays.

## Public API

```go
// Server tags
vngcloud.VServerGateway().V2().ComputeService().ListTags(
    lscomputeV2.NewListTagsRequest(serverId))

vngcloud.VServerGateway().V2().ComputeService().CreateTags(
    lscomputeV2.NewCreateTagsRequest(serverId).WithTags("env", "prod"))

vngcloud.VServerGateway().V2().ComputeService().UpdateTags(
    lscomputeV2.NewUpdateTagsRequest(serverId).WithTags("env", "prod"))

// Volume tags
vngcloud.VServerGateway().V2().VolumeService().ListTags(
    lsvolumeV2.NewListTagsRequest(volumeId))

vngcloud.VServerGateway().V2().VolumeService().CreateTags(
    lsvolumeV2.NewCreateTagsRequest(volumeId).WithTags("env", "prod"))

vngcloud.VServerGateway().V2().VolumeService().UpdateTags(
    lsvolumeV2.NewUpdateTagsRequest(volumeId).WithTags("env", "prod"))
```

## Component design

### URL builders

All three methods hit the same path; only the HTTP verb differs.

```go
// compute/v2/url.go
func listTagsUrl(psc lsclient.IServiceClient, popts IListTagsRequest) string {
    return psc.ServiceURL(psc.GetProjectId(), "tag", "resource", popts.GetServerId())
}
func createTagsUrl(psc lsclient.IServiceClient, popts ICreateTagsRequest) string {
    return psc.ServiceURL(psc.GetProjectId(), "tag", "resource", popts.GetServerId())
}
func updateTagsUrl(psc lsclient.IServiceClient, popts IUpdateTagsRequest) string {
    return psc.ServiceURL(psc.GetProjectId(), "tag", "resource", popts.GetServerId())
}
```

Volume version is identical, with `popts.GetBlockVolumeId()` instead.

### Request structs

```go
// compute/v2/tag_request.go (Volume mirrors with BlockVolumeCommon and "VOLUME")

const serverResourceType = "SERVER"

func NewListTagsRequest(pserverId string) IListTagsRequest {
    o := new(ListTagsRequest)
    o.ServerId = pserverId
    return o
}

func NewCreateTagsRequest(pserverId string) ICreateTagsRequest {
    o := new(CreateTagsRequest)
    o.ServerId = pserverId
    o.ResourceID = pserverId
    o.ResourceType = serverResourceType
    o.TagRequestList = make([]lscommon.Tag, 0)
    return o
}

func NewUpdateTagsRequest(pserverId string) IUpdateTagsRequest {
    o := new(UpdateTagsRequest)
    o.ServerId = pserverId
    o.ResourceID = pserverId
    o.ResourceType = serverResourceType
    o.TagRequestList = make([]lscommon.Tag, 0)
    return o
}

type ListTagsRequest struct {
    lscommon.UserAgent
    lscommon.ServerCommon  // exposes GetServerId()
}

type CreateTagsRequest struct {
    ResourceID     string         `json:"resourceId"`
    ResourceType   string         `json:"resourceType"`
    TagRequestList []lscommon.Tag `json:"tagRequestList"`
    lscommon.UserAgent
    lscommon.ServerCommon
}

type UpdateTagsRequest struct {
    ResourceID     string         `json:"resourceId"`
    ResourceType   string         `json:"resourceType"`
    TagRequestList []lscommon.Tag `json:"tagRequestList"`
    lscommon.UserAgent
    lscommon.ServerCommon
}
```

Plus the following builder methods on each struct, copied verbatim in shape from LB ([`loadbalancer/v2/tag_request.go`](../../../vngcloud/services/loadbalancer/v2/tag_request.go)):

- `AddUserAgent(pagent ...string)`
- `(*CreateTagsRequest).WithTags(ptags ...string)`
- `(*CreateTagsRequest).ToRequestBody() interface{}`
- `(*UpdateTagsRequest).WithTags(ptags ...string)`
- `(*UpdateTagsRequest).ToRequestBody(plstTags *lsentity.ListTags) interface{}` — implements system-tag preservation via dedup map
- `(*UpdateTagsRequest).ToMap() map[string]interface{}`

### Interfaces

```go
// compute/v2/irequest.go
type IListTagsRequest interface {
    GetServerId() string
    ParseUserAgent() string
    AddUserAgent(pagent ...string) IListTagsRequest
}

type ICreateTagsRequest interface {
    GetServerId() string
    ToRequestBody() interface{}
    ParseUserAgent() string
    WithTags(ptags ...string) ICreateTagsRequest
    AddUserAgent(pagent ...string) ICreateTagsRequest
}

type IUpdateTagsRequest interface {
    GetServerId() string
    ToRequestBody(plstTags *lsentity.ListTags) interface{}
    ParseUserAgent() string
    WithTags(ptags ...string) IUpdateTagsRequest
    ToMap() map[string]interface{}
    AddUserAgent(pagent ...string) IUpdateTagsRequest
}
```

Volume interfaces are identical except `GetServerId() string` is replaced by `GetBlockVolumeId() string`.

### Response

```go
// compute/v2/tag_response.go (Volume identical, just different package)

type ListTagResponse struct {
    Key       string `json:"key"`
    Value     string `json:"value"`
    CreatedAt string `json:"createdAt"`
    SystemTag bool   `json:"systemTag,omitempty"`
}
type ListTagsResponse []ListTagResponse

func (s ListTagResponse) ToEntityTag() *lsentity.Tag {
    return &lsentity.Tag{Key: s.Key, Value: s.Value, SystemTag: s.SystemTag}
}
func (s *ListTagsResponse) ToEntityListTags() *lsentity.ListTags {
    result := new(lsentity.ListTags)
    if s == nil {
        return result
    }
    for _, item := range *s {
        result.Add(item.ToEntityTag())
    }
    return result
}
```

### Service methods

```go
// compute/v2/tag.go
func (s *ComputeServiceV2) ListTags(popts IListTagsRequest) (*lsentity.ListTags, lserr.IError) {
    url := listTagsUrl(s.VServerClient, popts)
    resp := new(ListTagsResponse)
    errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
    req := lsclient.NewRequest().
        WithHeader("User-Agent", popts.ParseUserAgent()).
        WithOkCodes(200).
        WithJsonResponse(resp).
        WithJsonError(errResp)
    if _, sdkErr := s.VServerClient.Get(url, req); sdkErr != nil {
        return nil, lserr.SdkErrorHandler(sdkErr, errResp)
    }
    return resp.ToEntityListTags(), nil
}

func (s *ComputeServiceV2) CreateTags(popts ICreateTagsRequest) lserr.IError {
    url := createTagsUrl(s.VServerClient, popts)
    errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
    req := lsclient.NewRequest().
        WithHeader("User-Agent", popts.ParseUserAgent()).
        WithOkCodes(200).
        WithJsonBody(popts.ToRequestBody()).
        WithJsonError(errResp)
    if _, sdkErr := s.VServerClient.Put(url, req); sdkErr != nil {
        return lserr.SdkErrorHandler(sdkErr, errResp)
    }
    return nil
}

func (s *ComputeServiceV2) UpdateTags(popts IUpdateTagsRequest) lserr.IError {
    tmpTags, sdkErr := s.ListTags(NewListTagsRequest(popts.GetServerId()))
    if sdkErr != nil {
        return sdkErr
    }
    tags := new(lsentity.ListTags)
    for _, tag := range tmpTags.Items {
        if !tag.SystemTag {
            tags.Items = append(tags.Items, tag)
        }
    }
    url := updateTagsUrl(s.VServerClient, popts)
    errResp := lserr.NewErrorResponse(lserr.NormalErrorType)
    req := lsclient.NewRequest().
        WithHeader("User-Agent", popts.ParseUserAgent()).
        WithOkCodes(200).
        WithJsonBody(popts.ToRequestBody(tags)).
        WithJsonError(errResp)
    if _, sdkErr = s.VServerClient.Put(url, req); sdkErr != nil {
        return lserr.SdkErrorHandler(sdkErr, errResp,
            lserr.WithErrorTagKeyInvalid(errResp)).WithParameters(popts.ToMap())
    }
    return nil
}
```

Volume version is character-for-character identical except: package `v2` of volume, `popts.GetBlockVolumeId()` instead of `popts.GetServerId()`, and the inner `NewListTagsRequest` reference resolves inside the volume package.

## Method semantics

| Method | HTTP | System-tag behavior | When to use |
|---|---|---|---|
| `ListTags` | `GET` | Returns both system and user tags (caller can filter via `tag.SystemTag`). | Read tag list of a resource. |
| `CreateTags` | `PUT` | **Replaces the entire tag list** — system tags are wiped unless caller includes them. | Only when caller fully owns the tag list and accepts data loss risk. |
| `UpdateTags` | `PUT` | Internally `GET`s current list, filters out system tags, merges caller's tags via dedupe-by-key, then `PUT`s. Safe — preserves system tags. | **Default recommended method** for adding/updating user tags. |

`UpdateTags` semantics match LB's `UpdateTags` exactly: empty `WithTags` clears all user tags but keeps system tags.

## Constants reuse

Resource-type constants `"SERVER"`, `"VOLUME"`, `"LOAD-BALANCER"` are already defined in [`server/v1/systemtag_request.go:6-9`](../../../vngcloud/services/server/v1/systemtag_request.go#L6-L9). They are **not reused** here — each new package defines a local `const serverResourceType = "SERVER"` (or `volumeResourceType = "VOLUME"`) to avoid creating a reverse dependency from `compute/v2` or `volume/v2` to `server/v1`. The string values match.

## Testing

Integration tests against the real API (matches existing SDK test pattern — e.g. `TestCreateSystemTags` at [`test/server_test.go:399`](../../../test/server_test.go#L399), and LB's `TestListTags`/`TestCreateTags`/`TestUpdateTags` in [`test/lb_test.go`](../../../test/lb_test.go)).

Six new test functions to add:

```go
// test/server_test.go
func TestListServerTags(t *testing.T)     { ... }
func TestCreateServerTags(t *testing.T)   { ... }
func TestUpdateServerTags(t *testing.T)   { ... }

// test/volume_test.go
func TestListVolumeTags(t *testing.T)     { ... }
func TestCreateVolumeTags(t *testing.T)   { ... }
func TestUpdateVolumeTags(t *testing.T)   { ... }
```

Each test: build request via `NewXxxTagsRequest(resourceId).WithTags(...)`, call the service method, log result, expect `sdkerr == nil`. Resource IDs are hardcoded fixtures matching the existing test style.

No unit tests with HTTP mocks (SDK has no such pattern).

## Error handling

- `ListTags`, `CreateTags`: `lserr.NormalErrorType`, no extra error helpers.
- `UpdateTags`: `lserr.NormalErrorType` plus `lserr.WithErrorTagKeyInvalid(errResp)` — mirrors LB's `UpdateTags`. Existing helper, no new helper added.
- If implementation surfaces vserver-tag-API-specific error codes not present in LB, new `lserr.WithError*` helpers may be added; default is to reuse existing ones.

## Edge cases

| Case | Behavior |
|------|----------|
| `UpdateTags` with empty `WithTags` | `ListTags` → filter system → empty user list → `PUT` replaces user tags with empty → resource has only system tags. Correct "clear all user tags" semantics. |
| `WithTags` with odd number of args | Appends `"none"` as the missing value (matches LB at [`loadbalancer/v2/tag_request.go:78`](../../../vngcloud/services/loadbalancer/v2/tag_request.go#L78)). |
| Duplicate keys in `WithTags` | In `UpdateTags.ToRequestBody`, map dedupe keeps the last value (matches LB). `CreateTags.ToRequestBody` keeps the duplicates in the order given (no dedup) — matches LB. |
| `ListTags` on resource with no tags | API returns `[]`; `ToEntityListTags` returns an empty `*lsentity.ListTags`. No panic. |
| Concurrent `UpdateTags` calls on same resource | Lost-update is possible (PUT after GET race) — same limitation as LB's `UpdateTags`. Not addressed; documented as a known limitation. |
| `CreateSystemTags` legacy call path | Untouched. Continues to call `POST /<projectId>/tags` via `ServerServiceInternalV1`. |

## Out of scope

- Refactoring LB tag code into a shared helper.
- Generic `TagService` consolidating all resource types.
- Unit tests with HTTP mocking.
- `DeleteTags` method (no DELETE endpoint exists at `/tag/resource/`).
- Migrating existing `CreateSystemTags` callers to the new method.
