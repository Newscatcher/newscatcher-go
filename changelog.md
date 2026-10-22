# Changelog

## [v2.0.0] - 2026-10-22
### Breaking Changes
- **`aggregation`, `searchsimilar`, `searchlink`, and `latestheadlines` sub-packages** — entirely removed; replace imports with the new `aggregationcount`, `searchbylink`, and updated `latestheadlines` packages respectively.
- **`AggregationGetRequest/PostRequest`, `SearchSimilarGetRequest/PostRequest`, `SearchUrlGetRequest/PostRequest`, `LatestHeadlinesGetRequest/PostRequest`** — all removed and replaced by renamed equivalents (`GetAggregationCountRequest`, `GetSearchByLinkRequest`, `GetLatestHeadlinesRequest`, etc.); update all type references.
- **`NewClient`** across `authors`, `search`, `sources`, `subscription`, and `latestheadlines` packages now accepts a single `*core.RequestOptions` pointer instead of variadic `option.RequestOption`; update all call sites to pass a constructed `*core.RequestOptions` directly.
- **`Client.SearchLink`**, **`Client.Searchsimilar`**, **`Client.Latestheadlines`**, and **`Client.Aggregation`** root-client fields renamed or replaced: use `Client.SearchByLink`, `Client.LatestHeadlines`, and `Client.AggregationCount` instead.
- **Request/response types across `authors`, `search`, and `sources` packages** renamed from `VerbNounRequest/Response` to `VerbNounRequest/Response` pattern (e.g., `AuthorsGetRequest` → `GetAuthorsRequest`, `SearchGetRequest` → `GetSearchRequest`, `SourcesGetRequest` → `GetSourcesRequest`); update all type references.

### Added
- **`aggregationcount`** and **`searchbylink`** packages — new sub-clients (`Client` + `RawClient`) for the `/api/aggregation_count` and `/api/search_by_link` endpoints, replacing the removed packages.
- **`Client.WithRawResponse`** (`*RawClient`) field added to all sub-package clients, exposing raw HTTP status codes and headers via `core.Response[T]` alongside decoded response bodies.
- **`Set*` explicit-field setter methods** added to all major request and response types to prevent null/zero values from being omitted during JSON serialization.
- **`WithoutRetries()`, `WithMaxStreamBufSize()`, `WithMaxStreamReconnectAttempts()`, and `WithoutStreamReconnection()`** — new per-request option functions for fine-grained retry and streaming control.
- **`internal.ExplicitFieldsFromJSON`**, **`NewDateList`/`NewDateMap`/`NewDateTimeList`/`NewDateTimeMap`** helpers, and **`ErrorCodes`** exported variable added for improved JSON fidelity, date collection handling, and structured error decoding.

### Changed
- **Retry delay defaults** widened from 500ms–5s to 1s–60s; the retrier now honors `Retry-After` and `X-RateLimit-Reset` headers before falling back to exponential backoff.
- **`DateTime` serialization** — `UnmarshalJSON` now accepts unix epoch integers, space-separated strings, and fractional-second variants; `MarshalJSON` now uses `time.RFC3339Nano` to preserve sub-second precision.

### Fixed
- **`internal.EncodeURL`** — now correctly dereferences pointer arguments before path-escaping, preventing malformed `%!v(MISSING)` values in constructed URLs.

