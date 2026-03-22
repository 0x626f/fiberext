<div align="center">
    <pre style="background: none;">
██████████ █████ ███████████  ██████████ ███████████      ██████████ █████ █████ ███████████
░░███░░░░░█░░███ ░░███░░░░░███░░███░░░░░█░░███░░░░░███    ░░███░░░░░█░░███ ░░███ ░█░░░███░░░█
 ░███   █ ░  ░███  ░███    ░███ ░███  █ ░  ░███    ░███     ░███  █ ░  ░░███ ███  ░   ░███  ░
 ░███████    ░███  ░██████████  ░██████    ░██████████      ░██████     ░░█████       ░███
 ░███░░░█    ░███  ░███░░░░░███ ░███░░█    ░███░░░░░███     ░███░░█      ███░███      ░███
 ░███  ░     ░███  ░███    ░███ ░███ ░   █ ░███    ░███     ░███ ░   █  ███ ░░███     ░███
 █████       █████ ███████████  ██████████ █████   █████    ██████████ █████ █████    █████
░░░░░       ░░░░░ ░░░░░░░░░░░  ░░░░░░░░░░ ░░░░░   ░░░░░    ░░░░░░░░░░ ░░░░░ ░░░░░    ░░░░░
    </pre>
</div>

<div align="center">
    <h3>FiberExt - Extension of fiber library to streamline and simplify the development process</h3>
    <h6>Currently under active development and breaking changes are possible</h6>
</div>

---

## Quick Start

```go
package main

import (
    "context"
    "net/http"

    "github.com/0x626f/fiberext"
)

func main() {
    auth := func(c fiberext.Context) error {
        if c.Get("Authorization") == "" {
            return fiberext.Unauthorized(c)
        }
        return c.Next()
    }

    api := fiberext.NewController("/api/v1").
        AddNewResource(http.MethodGet,  "/users", listUsers).
        AddNewResource(http.MethodPost, "/users", createUser)

    cfg := fiberext.NewConfig().
        WithHost("0.0.0.0").
        WithPort(8080).
        WithDisableStartupMessage(true).
        WithMiddleware(auth).
        WithResource(fiberext.NewResource(http.MethodGet, "/health", healthHandler)).
        WithController(api)

    fiberext.Run(context.Background(), cfg)
}
```

---

## Reference

### Server

```go
server := fiberext.Run(ctx context.Context, cfg *Config) Server
```

Returns a `*fiber.App`. Registers middlewares, resources, and controllers in order, then starts the listener in the background. Cancelling `ctx` triggers graceful shutdown.

---

### Config

```go
cfg := fiberext.NewConfig()
```

All methods return `*Config` for chaining.

| Method | Description |
|---|---|
| `WithHost(string)` | Bind address (default `""`) |
| `WithPort(int)` | Bind port |
| `WithTLS(bool)` | Enable TLS |
| `WithMutualTLS(bool)` | Enable mutual TLS |
| `WithCertFile(string)` | Path to TLS certificate file |
| `WithKeyFile(string)` | Path to TLS key file |
| `WithClientCertFile(string)` | Path to client CA file (mTLS) |
| `WithCertificate(tls.Certificate)` | In-memory TLS certificate |
| `WithClientCerts(*x509.CertPool)` | In-memory client CA pool (mTLS) |
| `WithMiddleware(Handler)` | Append a global middleware |
| `WithResource(*Resource)` | Append a top-level route |
| `WithController(*Controller)` | Append a route group |
| `WithErrorHandler(ErrorHandler)` | Custom fiber error handler |
| `WithPrefork(bool)` | Multi-process listen |
| `WithBodyLimit(int)` | Max request body size in bytes |
| `WithReadTimeout(time.Duration)` | Read deadline |
| `WithWriteTimeout(time.Duration)` | Write deadline |
| `WithIdleTimeout(time.Duration)` | Keep-alive idle timeout |
| `WithStrictRouting(bool)` | Treat `/foo` and `/foo/` as different |
| `WithCaseSensitive(bool)` | Case-sensitive routing |
| `WithDisableStartupMessage(bool)` | Suppress Fiber banner |
| `WithAppName(string)` | Application name in headers |
| `WithServerHeader(string)` | Value of the `Server` response header |
| `WithJSONEncoder(utils.JSONMarshal)` | Custom JSON encoder |
| `WithJSONDecoder(utils.JSONUnmarshal)` | Custom JSON decoder |
| _…and all other `fiber.Config` fields_ | Prefixed with `With` |

`cfg.URL()` returns `"host:port"`.

---

### Controller

Groups routes under a common path prefix.

```go
ctrl := fiberext.NewController("/api/v1").
    AddNewResource(http.MethodGet,    "/users",     listUsers).
    AddNewResource(http.MethodPost,   "/users",     createUser).
    AddNewResource(http.MethodGet,    "/users/:id", getUser).
    AddNewResource(http.MethodDelete, "/users/:id", deleteUser)

cfg.WithController(ctrl)
// → GET    /api/v1/users
// → POST   /api/v1/users
// → GET    /api/v1/users/:id
// → DELETE /api/v1/users/:id
```

---

### Resource

A single route binding.

```go
r := fiberext.NewResource(method string, path string, handler Handler) *Resource
```

Top-level resources are registered directly on the server:

```go
cfg.WithResource(fiberext.NewResource(http.MethodGet, "/health", healthHandler))
```

---

### Request Helpers

```go
// Parse URL params into a struct (uses `params` field tags)
type P struct{ ID int `params:"id"` }
p, err := fiberext.FromParams[P](ctx)

// Parse JSON body into a struct
user, err := fiberext.FromBody[User](ctx)

// Parse query string into a struct (uses `query` field tags)
type Filter struct {
    Name  string `query:"name"`
    Limit int    `query:"limit"`
}
f, err := fiberext.FromQuery[Filter](ctx)

// Single URL param with optional default
id  := fiberext.GetParam(ctx, "id")
tab := fiberext.GetParam(ctx, "tab", "overview")

// Single query arg with optional default
sort := fiberext.GetQueryArg(ctx, "sort", "asc")
page := fiberext.GetQueryArg(ctx, "page", "1")
```

---

### Response Helpers

All helpers accept `ctx Context` as the first argument. Error helpers accept optional additional arguments; success helpers with a body accept `obj ...any`.

#### 2xx Success

```go
fiberext.OK(ctx, data)                       // 200
fiberext.Created(ctx, data)                  // 201
fiberext.Accepted(ctx, data)                 // 202
fiberext.NonAuthoritativeInformation(ctx, data) // 203
fiberext.NoContent(ctx)                      // 204
fiberext.ResetContent(ctx)                   // 205
fiberext.PartialContent(ctx, data)           // 206
fiberext.MultiStatus(ctx, data)              // 207
fiberext.AlreadyReported(ctx, data)          // 208
fiberext.IMUsed(ctx, data)                   // 226
```

#### 3xx Redirection

```go
fiberext.MovedPermanently(ctx)   // 301
fiberext.Found(ctx)              // 302
fiberext.SeeOther(ctx)           // 303
fiberext.NotModified(ctx)        // 304
fiberext.TemporaryRedirect(ctx)  // 307
fiberext.PermanentRedirect(ctx)  // 308
```

#### 4xx Client Errors

```go
fiberext.BadRequest(ctx)                    // 400
fiberext.Unauthorized(ctx)                  // 401
fiberext.Forbidden(ctx)                     // 403
fiberext.NotFound(ctx)                      // 404
fiberext.MethodNotAllowed(ctx)              // 405
fiberext.Conflict(ctx)                      // 409
fiberext.Gone(ctx)                          // 410
fiberext.UnprocessableEntity(ctx)           // 422
fiberext.TooManyRequests(ctx)               // 429
fiberext.UnavailableForLegalReasons(ctx)    // 451
// … all standard 4xx codes available
```

#### 5xx Server Errors

```go
fiberext.RespondInternalError(ctx)           // 500
fiberext.NotImplemented(ctx)                 // 501
fiberext.BadGateway(ctx)                     // 502
fiberext.ServiceUnavailable(ctx)             // 503
fiberext.GatewayTimeout(ctx)                 // 504
// … all standard 5xx codes available
```

---

### Types

| Type | Alias for |
|---|---|
| `fiberext.Server` | `*fiber.App` |
| `fiberext.Context` | `*fiber.Ctx` |
| `fiberext.Handler` | `fiber.Handler` |
| `fiberext.ErrorHandler` | `fiber.ErrorHandler` |

---

### Custom Error Handler

```go
cfg.WithErrorHandler(func(c fiberext.Context, err error) error {
    code := fiber.StatusInternalServerError
    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }
    return c.Status(code).JSON(fiber.Map{
        "code":    code,
        "message": err.Error(),
    })
})
```
