# Tag/Resource API SDK Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `ListTags`, `CreateTags`, `UpdateTags` methods to `ComputeServiceV2` (Server) and `VolumeServiceV2`, calling the `/v2/{projectId}/tag/resource/{resourceId}` endpoints from `vserver-api-spec.json`.

**Architecture:** Mirror the existing LoadBalancer tag implementation (`vngcloud/services/loadbalancer/v2/{tag.go, tag_request.go, tag_response.go, url.go}`) into two new locations: `vngcloud/services/compute/v2/` (for Server) and `vngcloud/services/volume/v2/` (for Volume). `UpdateTags` does GET-then-PUT to preserve system tags.

**Tech Stack:** Go 1.x, vngcloud-go-sdk client primitives (`lsclient.IServiceClient`, `lsentity.ListTags`, `lscommon.Tag`).

**Spec:** [docs/superpowers/specs/2026-05-17-tag-resource-api-sdk-design.md](../specs/2026-05-17-tag-resource-api-sdk-design.md)

**Reference implementation:** [vngcloud/services/loadbalancer/v2/tag.go](../../../vngcloud/services/loadbalancer/v2/tag.go), [tag_request.go](../../../vngcloud/services/loadbalancer/v2/tag_request.go), [tag_response.go](../../../vngcloud/services/loadbalancer/v2/tag_response.go), [url.go:146-168](../../../vngcloud/services/loadbalancer/v2/url.go#L146-L168)

---

## Conventions used in this plan

- **TDD via compile-check:** The SDK has no unit-test infrastructure; integration tests in `test/` hit real APIs and require credentials. The "failing test" gate at each task is `go build ./test/...` returning non-zero because the referenced symbols don't exist yet. Once the implementation lands, `go build ./...` and `go vet ./...` must both pass.
- **Running the integration tests** (with `go test -run ...`) is the developer's manual step after the SDK builds; it requires real `vngcloud` credentials wired through `validSuperSdkConfig()` / `validSdkConfig()`. Don't run them as part of the plan execution.
- **Test fixture IDs:** hardcoded UUIDs that mirror the existing pattern (e.g. `TestCreateSystemTags` at [test/server_test.go:399](../../../test/server_test.go#L399), `TestListTagsSuccess` at [test/lb_test.go:556](../../../test/lb_test.go#L556)). Developer replaces with valid IDs at runtime.
- **Commit message format:** `feat: <short summary>` to match existing project history.

---

## Task 1: Compute (Server) `ListTags`

**Files:**
- Create: `vngcloud/services/compute/v2/tag.go`
- Create: `vngcloud/services/compute/v2/tag_request.go`
- Create: `vngcloud/services/compute/v2/tag_response.go`
- Modify: `vngcloud/services/compute/v2/irequest.go` (append at end)
- Modify: `vngcloud/services/compute/v2/url.go` (append at end)
- Modify: `vngcloud/services/compute/icompute.go` (add 1 method to `IComputeServiceV2`)
- Test: `test/server_test.go` (append at end)

- [ ] **Step 1: Write the failing integration test**

Append to `test/server_test.go`:

```go
func TestListServerTagsSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := lscomputeSvcV2.NewListTagsRequest("ins-da59addd-6263-4544-b405-420a65ccfb1f")
	tags, sdkErr := vngcloud.VServerGateway().V2().ComputeService().ListTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	if tags == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", tags)
	t.Log("PASS")
}
```

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewListTagsRequest` or `ListTags` not declared in `lscomputeSvcV2` / on `IComputeServiceV2`.

- [ ] **Step 3: Create the response type**

Write `vngcloud/services/compute/v2/tag_response.go`:

```go
package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type ListTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag,omitempty"`
}

type ListTagsResponse []ListTagResponse

func (s ListTagResponse) ToEntityTag() *lsentity.Tag {
	return &lsentity.Tag{
		Key:       s.Key,
		Value:     s.Value,
		SystemTag: s.SystemTag,
	}
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

- [ ] **Step 4: Create the request struct + constructor**

Write `vngcloud/services/compute/v2/tag_request.go`:

```go
package v2

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	serverResourceType = "SERVER"
)

func NewListTagsRequest(pserverId string) IListTagsRequest {
	o := new(ListTagsRequest)
	o.ServerId = pserverId
	return o
}

func (s *ListTagsRequest) AddUserAgent(pagent ...string) IListTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ListTagsRequest struct {
	lscommon.UserAgent
	lscommon.ServerCommon
}
```

- [ ] **Step 5: Add the request interface**

Append to `vngcloud/services/compute/v2/irequest.go` (at end of file, before the closing of the file):

```go

type IListTagsRequest interface {
	GetServerId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListTagsRequest
}
```

- [ ] **Step 6: Add the URL builder**

Append to `vngcloud/services/compute/v2/url.go`:

```go

func listTagsUrl(psc lsclient.IServiceClient, popts IListTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag", "resource", popts.GetServerId())
}
```

- [ ] **Step 7: Add the service method**

Write `vngcloud/services/compute/v2/tag.go`:

```go
package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

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
```

- [ ] **Step 8: Register the method on the public interface**

Modify `vngcloud/services/compute/icompute.go` — add this line to the `IComputeServiceV2` interface block (keep alphabetical-ish placement; near the existing `CreateServerGroup` lines is fine):

```go
	ListTags(popts lscomputeSvcV2.IListTagsRequest) (*lsentity.ListTags, lserr.IError)
```

- [ ] **Step 9: Verify the build passes**

Run: `go build ./...`
Expected: exit 0, no errors.

Run: `go vet ./...`
Expected: exit 0, no warnings.

Run: `go build ./test/...`
Expected: exit 0 (test file compiles, even if not run).

- [ ] **Step 10: Commit**

```bash
git add vngcloud/services/compute/v2/tag.go \
        vngcloud/services/compute/v2/tag_request.go \
        vngcloud/services/compute/v2/tag_response.go \
        vngcloud/services/compute/v2/irequest.go \
        vngcloud/services/compute/v2/url.go \
        vngcloud/services/compute/icompute.go \
        test/server_test.go
git commit -m "feat: add ListTags to ComputeServiceV2 (server tag/resource API)"
```

---

## Task 2: Compute (Server) `CreateTags`

**Files:**
- Modify: `vngcloud/services/compute/v2/tag.go` (append)
- Modify: `vngcloud/services/compute/v2/tag_request.go` (append)
- Modify: `vngcloud/services/compute/v2/irequest.go` (append)
- Modify: `vngcloud/services/compute/v2/url.go` (append)
- Modify: `vngcloud/services/compute/icompute.go` (1 line)
- Test: `test/server_test.go` (append)

- [ ] **Step 1: Write the failing integration test**

Append to `test/server_test.go`:

```go
func TestCreateServerTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewCreateTagsRequest("ins-da59addd-6263-4544-b405-420a65ccfb1f").
		WithTags("env", "dev")
	sdkErr := vngcloud.VServerGateway().V2().ComputeService().CreateTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr.GetMessage())
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}
```

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewCreateTagsRequest` or `CreateTags` not declared.

- [ ] **Step 3: Extend the request file**

Append to `vngcloud/services/compute/v2/tag_request.go`:

```go

func NewCreateTagsRequest(pserverId string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.ServerId = pserverId
	opts.ResourceID = pserverId
	opts.ResourceType = serverResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type CreateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.ServerCommon
}

func (s *CreateTagsRequest) AddUserAgent(pagent ...string) ICreateTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateTagsRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateTagsRequest) WithTags(ptags ...string) ICreateTagsRequest {
	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}
	return s
}
```

- [ ] **Step 4: Add the request interface**

Append to `vngcloud/services/compute/v2/irequest.go`:

```go

type ICreateTagsRequest interface {
	GetServerId() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) ICreateTagsRequest
	AddUserAgent(pagent ...string) ICreateTagsRequest
}
```

- [ ] **Step 5: Add the URL builder**

Append to `vngcloud/services/compute/v2/url.go`:

```go

func createTagsUrl(psc lsclient.IServiceClient, popts ICreateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag", "resource", popts.GetServerId())
}
```

- [ ] **Step 6: Add the service method**

Append to `vngcloud/services/compute/v2/tag.go`:

```go

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
```

- [ ] **Step 7: Register the method on the public interface**

In `vngcloud/services/compute/icompute.go`, add to the `IComputeServiceV2` block, immediately after the `ListTags` line from Task 1:

```go
	CreateTags(popts lscomputeSvcV2.ICreateTagsRequest) lserr.IError
```

- [ ] **Step 8: Verify the build passes**

Run: `go build ./...`
Expected: exit 0.

Run: `go vet ./...`
Expected: exit 0.

Run: `go build ./test/...`
Expected: exit 0.

- [ ] **Step 9: Commit**

```bash
git add vngcloud/services/compute/v2/tag.go \
        vngcloud/services/compute/v2/tag_request.go \
        vngcloud/services/compute/v2/irequest.go \
        vngcloud/services/compute/v2/url.go \
        vngcloud/services/compute/icompute.go \
        test/server_test.go
git commit -m "feat: add CreateTags to ComputeServiceV2 (server tag/resource API)"
```

---

## Task 3: Compute (Server) `UpdateTags`

**Files:**
- Modify: `vngcloud/services/compute/v2/tag.go` (append)
- Modify: `vngcloud/services/compute/v2/tag_request.go` (append)
- Modify: `vngcloud/services/compute/v2/irequest.go` (append)
- Modify: `vngcloud/services/compute/v2/url.go` (append)
- Modify: `vngcloud/services/compute/icompute.go` (1 line)
- Test: `test/server_test.go` (append)

- [ ] **Step 1: Write the failing integration test**

Append to `test/server_test.go`:

```go
func TestUpdateServerTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := lscomputeSvcV2.NewUpdateTagsRequest("ins-da59addd-6263-4544-b405-420a65ccfb1f").
		WithTags("env", "prod")
	sdkErr := vngcloud.VServerGateway().V2().ComputeService().UpdateTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr.GetMessage())
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}
```

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewUpdateTagsRequest` or `UpdateTags` not declared.

- [ ] **Step 3: Add the request struct + constructor**

Append to `vngcloud/services/compute/v2/tag_request.go`:

```go

func NewUpdateTagsRequest(pserverId string) IUpdateTagsRequest {
	opts := new(UpdateTagsRequest)
	opts.ServerId = pserverId
	opts.ResourceID = pserverId
	opts.ResourceType = serverResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type UpdateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.ServerCommon
}

func (s *UpdateTagsRequest) AddUserAgent(pagent ...string) IUpdateTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *UpdateTagsRequest) ToRequestBody(plstTags *lsentity.ListTags) interface{} {
	st := map[string]lscommon.Tag{}
	for _, tag := range plstTags.Items {
		st[tag.Key] = lscommon.Tag{
			IsEdited: false,
			Key:      tag.Key,
			Value:    tag.Value,
		}
	}

	for _, tag := range s.TagRequestList {
		st[tag.Key] = tag
	}

	s.TagRequestList = make([]lscommon.Tag, 0)
	for _, tag := range st {
		s.TagRequestList = append(s.TagRequestList, tag)
	}

	return s
}

func (s *UpdateTagsRequest) WithTags(ptags ...string) IUpdateTagsRequest {
	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *UpdateTagsRequest) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for _, tag := range s.TagRequestList {
		res[tag.Key] = tag.Value
	}
	return res
}
```

Also update the file's import block at the top — add `lsentity` next to `lscommon`:

```go
import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)
```

- [ ] **Step 4: Add the request interface**

Append to `vngcloud/services/compute/v2/irequest.go`:

```go

type IUpdateTagsRequest interface {
	GetServerId() string
	ToRequestBody(plstTags *lsentity.ListTags) interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) IUpdateTagsRequest
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) IUpdateTagsRequest
}
```

Also update the import block at the top of `irequest.go` to add the `lsentity` import (if not already present):

```go
import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)
```

If `irequest.go` previously had no imports, change the file to start with the import above immediately under the `package v2` line.

- [ ] **Step 5: Add the URL builder**

Append to `vngcloud/services/compute/v2/url.go`:

```go

func updateTagsUrl(psc lsclient.IServiceClient, popts IUpdateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag", "resource", popts.GetServerId())
}
```

- [ ] **Step 6: Add the service method**

Append to `vngcloud/services/compute/v2/tag.go`:

```go

func (s *ComputeServiceV2) UpdateTags(popts IUpdateTagsRequest) lserr.IError {
	tmpTags, sdkErr := s.ListTags(NewListTagsRequest(popts.GetServerId()))
	if sdkErr != nil {
		return sdkErr
	}

	// Do not update system tags
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

- [ ] **Step 7: Register the method on the public interface**

In `vngcloud/services/compute/icompute.go`, add to the `IComputeServiceV2` block, immediately after the `CreateTags` line from Task 2:

```go
	UpdateTags(popts lscomputeSvcV2.IUpdateTagsRequest) lserr.IError
```

- [ ] **Step 8: Verify the build passes**

Run: `go build ./...`
Expected: exit 0.

Run: `go vet ./...`
Expected: exit 0.

Run: `go build ./test/...`
Expected: exit 0.

- [ ] **Step 9: Commit**

```bash
git add vngcloud/services/compute/v2/tag.go \
        vngcloud/services/compute/v2/tag_request.go \
        vngcloud/services/compute/v2/irequest.go \
        vngcloud/services/compute/v2/url.go \
        vngcloud/services/compute/icompute.go \
        test/server_test.go
git commit -m "feat: add UpdateTags to ComputeServiceV2 (server tag/resource API)"
```

---

## Task 4: Volume `ListTags`

**Files:**
- Create: `vngcloud/services/volume/v2/tag.go`
- Create: `vngcloud/services/volume/v2/tag_request.go`
- Create: `vngcloud/services/volume/v2/tag_response.go`
- Modify: `vngcloud/services/volume/v2/irequest.go` (append)
- Modify: `vngcloud/services/volume/v2/url.go` (append)
- Modify: `vngcloud/services/volume/ivolume.go` (add 1 method to `IVolumeServiceV2`)
- Test: `test/volume_test.go` (append)

- [ ] **Step 1: Write the failing integration test**

Append to `test/volume_test.go` (note the existing import alias is `v2`, not `lsvolumeSvcV2`):

```go
func TestListVolumeTagsSuccess(t *ltesting.T) {
	vngcloud := validSuperSdkConfig()
	opt := v2.NewListTagsRequest("vol-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
	tags, sdkErr := vngcloud.VServerGateway().V2().VolumeService().ListTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr)
	}

	if tags == nil {
		t.Fatalf("Expect not nil but got nil")
	}

	t.Log("Result: ", tags)
	t.Log("PASS")
}
```

If `validSuperSdkConfig` is not imported in `test/volume_test.go` already, no action — `test/` is a single package; helpers are in scope across files.

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewListTagsRequest` or `ListTags` not declared in `v2` package / on `IVolumeServiceV2`.

- [ ] **Step 3: Create the response type**

Write `vngcloud/services/volume/v2/tag_response.go`:

```go
package v2

import lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"

type ListTagResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	SystemTag bool   `json:"systemTag,omitempty"`
}

type ListTagsResponse []ListTagResponse

func (s ListTagResponse) ToEntityTag() *lsentity.Tag {
	return &lsentity.Tag{
		Key:       s.Key,
		Value:     s.Value,
		SystemTag: s.SystemTag,
	}
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

- [ ] **Step 4: Create the request struct + constructor**

Write `vngcloud/services/volume/v2/tag_request.go`:

```go
package v2

import (
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)

const (
	volumeResourceType = "VOLUME"
)

func NewListTagsRequest(pvolumeId string) IListTagsRequest {
	o := new(ListTagsRequest)
	o.BlockVolumeId = pvolumeId
	return o
}

func (s *ListTagsRequest) AddUserAgent(pagent ...string) IListTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

type ListTagsRequest struct {
	lscommon.UserAgent
	lscommon.BlockVolumeCommon
}
```

- [ ] **Step 5: Add the request interface**

Append to `vngcloud/services/volume/v2/irequest.go`:

```go

type IListTagsRequest interface {
	GetBlockVolumeId() string
	ParseUserAgent() string
	AddUserAgent(pagent ...string) IListTagsRequest
}
```

- [ ] **Step 6: Add the URL builder**

Append to `vngcloud/services/volume/v2/url.go`:

```go

func listTagsUrl(psc lsclient.IServiceClient, popts IListTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag", "resource", popts.GetBlockVolumeId())
}
```

- [ ] **Step 7: Add the service method**

Write `vngcloud/services/volume/v2/tag.go`:

```go
package v2

import (
	lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/client"
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lserr "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/sdk_error"
)

func (s *VolumeServiceV2) ListTags(popts IListTagsRequest) (*lsentity.ListTags, lserr.IError) {
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
```

- [ ] **Step 8: Register the method on the public interface**

In `vngcloud/services/volume/ivolume.go`, add this line inside the `IVolumeServiceV2` block (alongside `MigrateBlockVolumeById`):

```go
	ListTags(popts lsvolumeSvcV2.IListTagsRequest) (*lsentity.ListTags, lserr.IError)
```

- [ ] **Step 9: Verify the build passes**

Run: `go build ./...`
Expected: exit 0.

Run: `go vet ./...`
Expected: exit 0.

Run: `go build ./test/...`
Expected: exit 0.

- [ ] **Step 10: Commit**

```bash
git add vngcloud/services/volume/v2/tag.go \
        vngcloud/services/volume/v2/tag_request.go \
        vngcloud/services/volume/v2/tag_response.go \
        vngcloud/services/volume/v2/irequest.go \
        vngcloud/services/volume/v2/url.go \
        vngcloud/services/volume/ivolume.go \
        test/volume_test.go
git commit -m "feat: add ListTags to VolumeServiceV2 (volume tag/resource API)"
```

---

## Task 5: Volume `CreateTags`

**Files:**
- Modify: `vngcloud/services/volume/v2/tag.go` (append)
- Modify: `vngcloud/services/volume/v2/tag_request.go` (append)
- Modify: `vngcloud/services/volume/v2/irequest.go` (append)
- Modify: `vngcloud/services/volume/v2/url.go` (append)
- Modify: `vngcloud/services/volume/ivolume.go` (1 line)
- Test: `test/volume_test.go` (append)

- [ ] **Step 1: Write the failing integration test**

Append to `test/volume_test.go`:

```go
func TestCreateVolumeTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewCreateTagsRequest("vol-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee").
		WithTags("env", "dev")
	sdkErr := vngcloud.VServerGateway().V2().VolumeService().CreateTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr.GetMessage())
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}
```

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewCreateTagsRequest` or `CreateTags` not declared.

- [ ] **Step 3: Extend the request file**

Append to `vngcloud/services/volume/v2/tag_request.go`:

```go

func NewCreateTagsRequest(pvolumeId string) ICreateTagsRequest {
	opts := new(CreateTagsRequest)
	opts.BlockVolumeId = pvolumeId
	opts.ResourceID = pvolumeId
	opts.ResourceType = volumeResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type CreateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.BlockVolumeCommon
}

func (s *CreateTagsRequest) AddUserAgent(pagent ...string) ICreateTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *CreateTagsRequest) ToRequestBody() interface{} {
	return s
}

func (s *CreateTagsRequest) WithTags(ptags ...string) ICreateTagsRequest {
	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}
	return s
}
```

- [ ] **Step 4: Add the request interface**

Append to `vngcloud/services/volume/v2/irequest.go`:

```go

type ICreateTagsRequest interface {
	GetBlockVolumeId() string
	ToRequestBody() interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) ICreateTagsRequest
	AddUserAgent(pagent ...string) ICreateTagsRequest
}
```

- [ ] **Step 5: Add the URL builder**

Append to `vngcloud/services/volume/v2/url.go`:

```go

func createTagsUrl(psc lsclient.IServiceClient, popts ICreateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag", "resource", popts.GetBlockVolumeId())
}
```

- [ ] **Step 6: Add the service method**

Append to `vngcloud/services/volume/v2/tag.go`:

```go

func (s *VolumeServiceV2) CreateTags(popts ICreateTagsRequest) lserr.IError {
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
```

- [ ] **Step 7: Register the method on the public interface**

In `vngcloud/services/volume/ivolume.go`, add to the `IVolumeServiceV2` block, immediately after the `ListTags` line from Task 4:

```go
	CreateTags(popts lsvolumeSvcV2.ICreateTagsRequest) lserr.IError
```

- [ ] **Step 8: Verify the build passes**

Run: `go build ./...`
Expected: exit 0.

Run: `go vet ./...`
Expected: exit 0.

Run: `go build ./test/...`
Expected: exit 0.

- [ ] **Step 9: Commit**

```bash
git add vngcloud/services/volume/v2/tag.go \
        vngcloud/services/volume/v2/tag_request.go \
        vngcloud/services/volume/v2/irequest.go \
        vngcloud/services/volume/v2/url.go \
        vngcloud/services/volume/ivolume.go \
        test/volume_test.go
git commit -m "feat: add CreateTags to VolumeServiceV2 (volume tag/resource API)"
```

---

## Task 6: Volume `UpdateTags`

**Files:**
- Modify: `vngcloud/services/volume/v2/tag.go` (append)
- Modify: `vngcloud/services/volume/v2/tag_request.go` (append)
- Modify: `vngcloud/services/volume/v2/irequest.go` (append)
- Modify: `vngcloud/services/volume/v2/url.go` (append)
- Modify: `vngcloud/services/volume/ivolume.go` (1 line)
- Test: `test/volume_test.go` (append)

- [ ] **Step 1: Write the failing integration test**

Append to `test/volume_test.go`:

```go
func TestUpdateVolumeTagsSuccess(t *ltesting.T) {
	vngcloud := validSdkConfig()
	opt := v2.NewUpdateTagsRequest("vol-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee").
		WithTags("env", "prod")
	sdkErr := vngcloud.VServerGateway().V2().VolumeService().UpdateTags(opt)
	if sdkErr != nil {
		t.Fatalf("Expect nil but got %+v", sdkErr.GetMessage())
	}

	t.Log("Result: ", sdkErr)
	t.Log("PASS")
}
```

- [ ] **Step 2: Verify the build fails**

Run: `go build ./test/...`
Expected: build error referencing `NewUpdateTagsRequest` or `UpdateTags` not declared.

- [ ] **Step 3: Add the request struct + constructor + builders**

Update the import block at the top of `vngcloud/services/volume/v2/tag_request.go` to add `lsentity`:

```go
import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
	lscommon "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/common"
)
```

Append to `vngcloud/services/volume/v2/tag_request.go`:

```go

func NewUpdateTagsRequest(pvolumeId string) IUpdateTagsRequest {
	opts := new(UpdateTagsRequest)
	opts.BlockVolumeId = pvolumeId
	opts.ResourceID = pvolumeId
	opts.ResourceType = volumeResourceType
	opts.TagRequestList = make([]lscommon.Tag, 0)
	return opts
}

type UpdateTagsRequest struct {
	ResourceID     string         `json:"resourceId"`
	ResourceType   string         `json:"resourceType"`
	TagRequestList []lscommon.Tag `json:"tagRequestList"`

	lscommon.UserAgent
	lscommon.BlockVolumeCommon
}

func (s *UpdateTagsRequest) AddUserAgent(pagent ...string) IUpdateTagsRequest {
	s.UserAgent.AddUserAgent(pagent...)
	return s
}

func (s *UpdateTagsRequest) ToRequestBody(plstTags *lsentity.ListTags) interface{} {
	st := map[string]lscommon.Tag{}
	for _, tag := range plstTags.Items {
		st[tag.Key] = lscommon.Tag{
			IsEdited: false,
			Key:      tag.Key,
			Value:    tag.Value,
		}
	}

	for _, tag := range s.TagRequestList {
		st[tag.Key] = tag
	}

	s.TagRequestList = make([]lscommon.Tag, 0)
	for _, tag := range st {
		s.TagRequestList = append(s.TagRequestList, tag)
	}

	return s
}

func (s *UpdateTagsRequest) WithTags(ptags ...string) IUpdateTagsRequest {
	if len(ptags)%2 != 0 {
		ptags = append(ptags, "none")
	}

	for i := 0; i < len(ptags); i += 2 {
		s.TagRequestList = append(s.TagRequestList, lscommon.Tag{Key: ptags[i], Value: ptags[i+1]})
	}

	return s
}

func (s *UpdateTagsRequest) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for _, tag := range s.TagRequestList {
		res[tag.Key] = tag.Value
	}
	return res
}
```

- [ ] **Step 4: Add the request interface**

Update the import block at the top of `vngcloud/services/volume/v2/irequest.go` to add `lsentity` (if not already present from Task 4 — it was not):

```go
import (
	lsentity "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/entity"
)
```

Append to `vngcloud/services/volume/v2/irequest.go`:

```go

type IUpdateTagsRequest interface {
	GetBlockVolumeId() string
	ToRequestBody(plstTags *lsentity.ListTags) interface{}
	ParseUserAgent() string
	WithTags(ptags ...string) IUpdateTagsRequest
	ToMap() map[string]interface{}
	AddUserAgent(pagent ...string) IUpdateTagsRequest
}
```

- [ ] **Step 5: Add the URL builder**

Append to `vngcloud/services/volume/v2/url.go`:

```go

func updateTagsUrl(psc lsclient.IServiceClient, popts IUpdateTagsRequest) string {
	return psc.ServiceURL(
		psc.GetProjectId(),
		"tag", "resource", popts.GetBlockVolumeId())
}
```

- [ ] **Step 6: Add the service method**

Append to `vngcloud/services/volume/v2/tag.go`:

```go

func (s *VolumeServiceV2) UpdateTags(popts IUpdateTagsRequest) lserr.IError {
	tmpTags, sdkErr := s.ListTags(NewListTagsRequest(popts.GetBlockVolumeId()))
	if sdkErr != nil {
		return sdkErr
	}

	// Do not update system tags
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

- [ ] **Step 7: Register the method on the public interface**

In `vngcloud/services/volume/ivolume.go`, add to the `IVolumeServiceV2` block, immediately after the `CreateTags` line from Task 5:

```go
	UpdateTags(popts lsvolumeSvcV2.IUpdateTagsRequest) lserr.IError
```

- [ ] **Step 8: Verify the build passes**

Run: `go build ./...`
Expected: exit 0.

Run: `go vet ./...`
Expected: exit 0.

Run: `go build ./test/...`
Expected: exit 0.

- [ ] **Step 9: Commit**

```bash
git add vngcloud/services/volume/v2/tag.go \
        vngcloud/services/volume/v2/tag_request.go \
        vngcloud/services/volume/v2/irequest.go \
        vngcloud/services/volume/v2/url.go \
        vngcloud/services/volume/ivolume.go \
        test/volume_test.go
git commit -m "feat: add UpdateTags to VolumeServiceV2 (volume tag/resource API)"
```

---

## Final verification (post Task 6)

- [ ] **Run the full build + vet one last time**

Run: `go build ./...`
Expected: exit 0.

Run: `go vet ./...`
Expected: exit 0.

- [ ] **Optional: run the integration tests (requires real credentials)**

These tests hit the live vserver gateway. Replace the hardcoded resource IDs in each new test with IDs from a live project before running.

Run any of:

```
go test -v -run TestListServerTagsSuccess ./test/...
go test -v -run TestCreateServerTagsSuccess ./test/...
go test -v -run TestUpdateServerTagsSuccess ./test/...
go test -v -run TestListVolumeTagsSuccess ./test/...
go test -v -run TestCreateVolumeTagsSuccess ./test/...
go test -v -run TestUpdateVolumeTagsSuccess ./test/...
```

Expected: PASS, logged tag list / `nil` sdkErr. Failures here indicate live-environment issues, not code defects — investigate API response separately.
