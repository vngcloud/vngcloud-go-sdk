# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

`vngcloud-go-sdk` is a Go SDK for VNG Cloud services. Importers use module path `github.com/vngcloud/vngcloud-go-sdk/v2` (the `/v2` suffix is mandatory — Go modules major-version routing). Go 1.22+.

Operational context (incidents, invariants, farms, the Redmine task workflow) lives in the **vks-harness** repo (`knowledge/`, `AGENTS.md`). Management farms are read-only for agents.

Consumers construct an `IClient` from `client/`, configure endpoints + IAM creds via `ISdkConfigure`, then call typed service methods through gateway accessors (e.g. `vngcloud.VLBGateway().V2().LoadBalancerService().CreateLoadBalancer(opt)`). The README shows the canonical usage shape.

## Build / lint / test

```bash
make verify-fast     # vet + lint (new code only) + unit tests, no credentials; run before finishing a change
make test            # unit tests, excludes ./test/...
make lint            # golangci-lint pinned to v2.6.x in ./bin, only code new vs LINT_BASE (default origin/main)
make test-integration  # opt-in, live APIs, needs test/env.yaml; never part of verify-fast
go build ./...
gitleaks detect --config .gitleaks.toml    # CI secret scan
```

`make lint` installs golangci-lint with `go install`, which needs a recent Go toolchain (1.25+); the library itself targets Go 1.22. CI runs the same linter version with `.golangci.yml`.

The `test/` package contains **integration tests that hit real VNG Cloud APIs** — they are not unit tests. They read credentials and resource IDs from `test/env.yaml` via `joho/godotenv` (see `getValueOfEnv` / `validSdkConfig` in `test/identity_test.go`). `test/env.yaml` is gitignored. Plain `go test ./...` (which includes `./test`) will fail without it or will issue real API calls; do not run the whole `test/` suite blind.

Run a single integration test (must have `test/env.yaml` populated):

```bash
go test ./test -run TestCreateLoadBalancerSuccess -v
```

`test/data_test.go` holds fake-but-realistic-looking certs/keys/tokens — they are mock fixtures, not real secrets. Don't "redact" them.

## Release

Tag `vX.Y.Z` and push tags. `.github/workflows/release_build.yml` (GoReleaser) builds the release; a Telegram notification fires on every commit. The module path enforces the `v2.x` major line — bumping past v3 requires updating `go.mod` and every internal import.

```bash
git tag -am "[release] release new version" v2.X.Y
git push --tags
```

## Architecture

Three nested layers — **client → gateway → service** — with versioned subpackages on the service layer. Read top-down when unsure where to put something.

### Top-level `client/` (the public entry point)

- `client.IClient` (in `client/iclient.go`) is what consumers hold. Builder methods (`WithRetryCount`, `WithProjectId`, …) plus per-product gateway accessors: `IamGateway()`, `VServerGateway()`, `VLBGateway()`, `VNetworkGateway()`, `GLBGateway()`, `VDnsGateway()`, `VDBKafkaGateway()`, `VDBOpenSearchGateway()`. **Adding a new product means adding both an accessor here and a gateway under `vngcloud/gateway/`.**
- `client.ISdkConfigure` is a separate fluent builder for endpoints + credentials. Each new product endpoint needs matching `With…Endpoint` / `Get…Endpoint` methods.
- `Configure()` lazily constructs each gateway only when its endpoint is non-empty — callers can opt in to subsets of products.

### `vngcloud/gateway/` (product routing)

- `igateway.go` defines the interface hierarchy: `IXxxGateway → V1()/V2()/Internal()` → `IXxxGatewayV2 → XxxService()`. This is the only place version selection happens — services themselves never know their version.
- `gateway.go` (and `iam_gateway.go`, `vlb_gateway.go`, etc.) construct one `IServiceClient` per `endpoint+version` pair (`endpoint + "v2"`, `endpoint + "internal"`, …) and inject it into the service struct. **The `+ "v2"` style URL suffixing happens here, not in the service layer.**

### `vngcloud/services/<product>/<version>/` (the API calls)

Each version folder follows a fixed file layout:

- `base.go` — service struct holding `IServiceClient`s (a service can hold multiple clients, e.g. VLBv2 holds both `VLBClient` and `VServerClient` because cross-gateway calls are needed).
- `url.go` — pure functions building URLs via `psc.ServiceURL(parts...)`. Query strings appended manually.
- `irequest.go` — fluent request interfaces (`ICreateXxxRequest`).
- `xxx_request.go` / `xxx_response.go` — request body builders and response → entity converters.
- `xxx.go` — the service methods themselves. Each one is uniform: build URL → build `errResp` of the right type → `lsclient.NewRequest().WithJsonBody/Response/Error().WithOkCodes(...)` → call `s.XxxClient.Get/Post/Put/Delete/Patch(url, req)` → on error wrap via `lserr.SdkErrorHandler(sdkErr, errResp, lserr.WithErrorXxx(errResp), …).AppendCategories(lserr.ErrCatProductXxx)`.

### `vngcloud/entity/` (returned types)

Pure data structs — `LoadBalancer`, `KafkaCluster`, `Volume`, … Service responses define `ToEntityXxx()` to convert wire types into these. Consumers receive entities, never raw responses.

### `vngcloud/sdk_error/` (error model)

Everything fallible returns `(*Entity, lserr.IError)`, **never** plain `error`. `SdkError` carries an error code, a category set, parameters, and the wrapped underlying error. Pattern: define product-specific `WithErrorXxxNotFound` matchers in a `loadbalancer.go`-style file, then chain them into `SdkErrorHandler(...)`. `ErrorCategory` (e.g. `ErrCatProductVlb`) lets callers branch by product without parsing codes. New error responses with non-standard JSON shapes plug in via `NewErrorResponse(ptype)` and a new `…ErrorType` const in `errors.go`.

### `vngcloud/client/` (HTTP plumbing — distinct from top-level `client/`)

- `IHttpClient` — the auth-aware HTTP layer (built on `imroc/req/v3`). Owns retry/sleep, default headers, and the reauth callback (`WithReauthFunc(IamOauth2, fn)`).
- `IServiceClient` — wraps `IHttpClient` plus product context (`projectId`, `zoneId`, `userId`, endpoint base). `ServiceURL(parts...)` is the URL builder every service uses.
- `IRequest` (`request.go`) — the per-call options: `WithJsonBody`, `WithJsonResponse`, `WithJsonError`, `WithOkCodes`, headers. `request.go` and `http.go` are where to look for header/retry/auth behavior changes.

## Code conventions in this repo

These are **deliberate, repo-wide patterns** — match them when editing or adding code:

- **Aliased imports with `ls`/`l` prefix**: `lsclient`, `lsentity`, `lserr`, `lctx "context"`, `lfmt "fmt"`, `ltime "time"`, `lreq "github.com/imroc/req/v3"`. Stick to existing aliases for the package being imported; see neighboring files for the convention.
- **Interface-first APIs**: every public type is an `I…` interface; the struct is unexported (or lowercase) and a `New…` constructor returns the interface. Fluent builders return the interface itself (`With…(…) IClient`).
- **Parameter prefix `p`**: function parameters are named `pctx`, `popts`, `pclient`, `pendpoint`, etc. Local vars and struct fields are not prefixed.
- **`sdkerr`/`lserr.IError` everywhere**: do not return `error`. New service methods must wrap via `lserr.SdkErrorHandler` and attach a category from `sdk_error/categories.go`.
- **URL builder per endpoint** in `url.go`, never inline string concatenation in the service method. Query strings go after `ServiceURL(...)`.
- **Versioned services are physical directories** (`v1/`, `v2/`, `inter/`, `internal/`). Don't add version branching inside a single file — add a new directory and wire it in via the gateway.
- **Tests for new services live in `test/`** as `<product>_test.go` and reuse the `validSdkConfig()` family from `test/identity_test.go`. They are real-API integration tests, not unit tests.

## Sensitive paths

- `vngcloud/services/*/*/` delete/teardown methods (`Delete*`, `Remove*`) call real cloud APIs that destroy resources; URL, method and OK-codes must match the API exactly. Never exercise them from `test/` against a shared project.
- `vngcloud/client/http.go`, `request.go`: auth, reauth callback, retry; a change affects every service.
- `vngcloud/sdk_error/`: error codes and categories are public API consumed by controllers; do not rename or renumber.
- `client/` and `vngcloud/gateway/`: public interface surface and endpoint wiring; breaking changes need a new major version.
- `go.mod` module path (`/v2`) and `.github/workflows/release_build.yml` / `.goreleaser.yaml`: release flow.

## Things easy to get wrong

- Module path is `…/v2`, not `…`. New files importing other internal packages must use the `/v2/` prefix.
- `WithProjectId` on a configured client re-creates gateways in place (see `client.go`). Don't cache gateway references across project switches.
- `WithVNetworkEndpoint` is sometimes called twice in tests to override — the second call wins. Not a bug.
- Lint config (`.golangci.yml`) relaxes `dupl`/`gocyclo`/`goconst` only for `*_test.go`, `test/*`, `*request.go`, `*response.go`. Don't push duplicated logic outside those.
- English only in code, comments, docs and commit messages. Run `make verify-fast` before finishing.
