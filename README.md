# Newscatcher Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern) [![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/Newscatcher/newscatcher-go)

The Newscatcher Go library lets you access the Newscatcher API from Go.

## Documentation

Find the complete API reference [here](https://www.newscatcherapi.com/docs/v3/api-reference).

## Requirements

This module requires Go version >= 1.13.

## Installation

Run this command to use the Newscatcher Go library in your module:

```sh
go get github.com/Newscatcher/newscatcher-go
```

## Usage

Build and use the client like this:

```go
import (
    "context"
    newscatcher "github.com/Newscatcher/newscatcher-go"
    newscatcherclient "github.com/Newscatcher/newscatcher-go/client"
    "github.com/Newscatcher/newscatcher-go/option"
)

client := newscatcherclient.NewClient(
    option.WithApiKey("YOUR_API_KEY"),
)

response, err := client.Search.Post(
    context.TODO(),
 &newscatcher.SearchPostRequest{
        Q: "renewable energy",
        PredefinedSources: []string{"top 50 US"},
        Lang: []string{"en"},
        From: "2024-01-01",
        To: "2024-06-30",
 },
)
```

## Error Handling

When the API responds with a non-success status code, the SDK returns an error.

```go
import (
    "github.com/Newscatcher/newscatcher-go/core"
    "fmt"
)

response, err := client.Search.Post(...)
if err != nil {
    if apiErr, ok := err.(*core.APIError); ok {
        fmt.Println(apiErr.Error())
        fmt.Println(apiErr.StatusCode)
 }
    return err
}
```

You can use Go's standard error handling features with these errors:

```go
import (
    import "errors"
    newscatcher "github.com/Newscatcher/newscatcher-go"
)

response, err := client.Search.Post(...)
if err != nil {
    var badRequestErr *newscatcher.BadRequestError
    if errors.As(err, &badRequestErr) {
        // Handle bad request
        fmt.Println(badRequestErr.Error())
 }
    // Wrap error with additional context
    return fmt.Errorf("search failed: %w", err)
}
```

The SDK provides specific error types for different cases:

```go
response, err := client.Search.Post(...)
if err != nil {
    switch e := err.(type) {
    case *newscatcher.BadRequestError:
        // Handle 400 error - Invalid request parameters
    case *newscatcher.UnauthorizedError:
        // Handle 401 error - Authentication failed
    case *newscatcher.ForbiddenError:
        // Handle 403 error - Server refuses action
    case *newscatcher.RequestTimeoutError:
        // Handle 408 error - Request timeout
    case *newscatcher.UnprocessableEntityError:
        // Handle 422 error - Validation error
    case *newscatcher.TooManyRequestsError:
        // Handle 429 error - Rate limit exceeded
    case *newscatcher.InternalServerError:
        // Handle 500 error - Server error
    default:
        return err
 }
}
```

## Advanced Features

### Retries

The SDK automatically retries failed requests with exponential backoff. When a request fails with any of these status codes, the SDK retries up to 2 times:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to your liking. For example, if you want to disable retries for the client entirely, you can set this value to 1 like so:

```go
client := newscatcherclient.NewClient(
    option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
response, err := client.Search.Post(
    context.TODO(),
    ...,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Set timeouts using the standard `context` package:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.Search.Post(
    ctx,
    request,
)
```

### Request Options

Configure the client using options at initialization or per request:

```go
// Client-level configuration
client := newscatcherclient.NewClient(
    option.WithApiKey("YOUR_API_KEY"),
    option.WithHTTPClient(&http.Client{
        Timeout: 30 * time.Second,
    }),
    option.WithMaxAttempts(3),
)

// Request-level configuration
response, err := client.Search.Post(
    context.TODO(),
    request,
    option.WithMaxAttempts(1),      // Override retries for this request
)
```

Available options include:

- `WithApiKey(string)`: Set authentication key.
- `WithHTTPClient(*http.Client)`: Use custom HTTP client.
- `WithMaxAttempts(uint)`: Configure retry attempts.
- `WithBaseURL(string)`: Override API base URL.
- `WithHTTPHeader(http.Header)`: Add custom headers.

> **Caution:** We recommend using your own `*http.Client`. By default, the SDK uses `http.DefaultClient`, which waits indefinitely for responses unless you set a context-based timeout.

## Contributing

We generate this SDK programmatically, so we can't accept direct code contributions. Instead:

1. [Open an issue](https://github.com/Newscatcher/newscatcher-go/issues) to discuss your ideas with us first.
2. If you want to show an implementation, create a PR as a proof of concept.
3. We'll work with you to implement the changes in our generator.

On the other hand, contributions to the README are always very welcome!
