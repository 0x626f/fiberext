package fiberext_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/0x626f/fiberext"
	"github.com/gofiber/fiber/v3"
)

// ---- domain types ----

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ---- test helpers ----

// startServer picks a free port, starts the server and registers cleanup.
// Tests interact with the returned Server via srv.Test() which bypasses the
// network and exercises the full handler chain in-process.
func startServer(t *testing.T, cfg *fiberext.Config) fiberext.Server {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("find free port: %v", err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cfg.WithHost("127.0.0.1").WithPort(port)
	srv := fiberext.Run(ctx, cfg)

	t.Cleanup(func() {
		_ = srv.Shutdown()
		cancel()
	})
	return srv
}

func jsonBody(t *testing.T, v any) *bytes.Buffer {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return bytes.NewBuffer(b)
}

func newRequest(t *testing.T, method, path string, body any) *http.Request {
	t.Helper()
	var buf *bytes.Buffer
	if body != nil {
		buf = jsonBody(t, body)
	} else {
		buf = &bytes.Buffer{}
	}
	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}

func decode[T any](t *testing.T, r io.Reader) T {
	t.Helper()
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return v
}

func assertStatus(t *testing.T, resp *http.Response, want int) {
	t.Helper()
	if resp.StatusCode != want {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("status: want %d got %d — body: %s", want, resp.StatusCode, body)
	}
}

func assertHeader(t *testing.T, resp *http.Response, key, want string) {
	t.Helper()
	if got := resp.Header.Get(key); got != want {
		t.Fatalf("header %q: want %q got %q", key, want, got)
	}
}

func do(t *testing.T, srv fiberext.Server, req *http.Request) *http.Response {
	t.Helper()
	resp, err := srv.Test(req, fiber.TestConfig{Timeout: 0})
	if err != nil {
		t.Fatalf("srv.Test: %v", err)
	}
	return resp
}

// ---- tests ----

// TestServer_Resources_CRUD registers GET/POST/GET:id/PUT:id/DELETE:id routes
// via Resources and verifies the full request-response cycle for each method.
func TestServer_Resources_CRUD(t *testing.T) {
	store := map[int]User{
		1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
		2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	seq := 3

	cfg := fiberext.NewConfig().
		WithResource(fiberext.NewResource(http.MethodGet, "/users", func(c fiberext.Context) error {
			list := make([]User, 0, len(store))
			for _, u := range store {
				list = append(list, u)
			}
			return c.Status(http.StatusOK).JSON(list)
		})).
		WithResource(fiberext.NewResource(http.MethodPost, "/users", func(c fiberext.Context) error {
			u, err := fiberext.FromBody[User](c)
			if err != nil {
				return fiberext.BadRequest(c)
			}
			u.ID = seq
			seq++
			store[u.ID] = u
			return c.Status(http.StatusCreated).JSON(u)
		})).
		WithResource(fiberext.NewResource(http.MethodGet, "/users/:id", func(c fiberext.Context) error {
			type p struct {
				ID int `uri:"id"`
			}
			params, _ := fiberext.FromParams[p](c)
			u, ok := store[params.ID]
			if !ok {
				return fiberext.NotFound(c)
			}
			return c.JSON(u)
		})).
		WithResource(fiberext.NewResource(http.MethodPut, "/users/:id", func(c fiberext.Context) error {
			type p struct {
				ID int `uri:"id"`
			}
			params, _ := fiberext.FromParams[p](c)
			if _, ok := store[params.ID]; !ok {
				return fiberext.NotFound(c)
			}
			u, err := fiberext.FromBody[User](c)
			if err != nil {
				return fiberext.BadRequest(c)
			}
			u.ID = params.ID
			store[u.ID] = u
			return c.JSON(u)
		})).
		WithResource(fiberext.NewResource(http.MethodDelete, "/users/:id", func(c fiberext.Context) error {
			type p struct {
				ID int `uri:"id"`
			}
			params, _ := fiberext.FromParams[p](c)
			delete(store, params.ID)
			return fiberext.NoContent(c)
		}))

	srv := startServer(t, cfg)

	t.Run("list all users", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users", nil))
		assertStatus(t, resp, http.StatusOK)
		users := decode[[]User](t, resp.Body)
		if len(users) != 2 {
			t.Fatalf("want 2 users, got %d", len(users))
		}
	})

	t.Run("create user returns 201 with assigned ID", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodPost, "/users", User{Name: "Charlie", Email: "charlie@example.com"}))
		assertStatus(t, resp, http.StatusCreated)
		created := decode[User](t, resp.Body)
		if created.ID == 0 {
			t.Fatal("expected assigned ID, got 0")
		}
		if created.Name != "Charlie" {
			t.Fatalf("name: want Charlie, got %s", created.Name)
		}
	})

	t.Run("get existing user by ID", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users/1", nil))
		assertStatus(t, resp, http.StatusOK)
		u := decode[User](t, resp.Body)
		if u.ID != 1 || u.Name != "Alice" {
			t.Fatalf("want Alice/1, got %+v", u)
		}
	})

	t.Run("get non-existent user returns 404", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users/9999", nil))
		assertStatus(t, resp, http.StatusNotFound)
	})

	t.Run("update existing user", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodPut, "/users/1", User{Name: "Alice Updated", Email: "new@example.com"}))
		assertStatus(t, resp, http.StatusOK)
		u := decode[User](t, resp.Body)
		if u.Name != "Alice Updated" {
			t.Fatalf("name: want Alice Updated, got %s", u.Name)
		}
		if u.ID != 1 {
			t.Fatalf("ID must be preserved: want 1, got %d", u.ID)
		}
	})

	t.Run("update non-existent user returns 404", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodPut, "/users/9999", User{Name: "Ghost"}))
		assertStatus(t, resp, http.StatusNotFound)
	})

	t.Run("delete user then confirm 404", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodDelete, "/users/2", nil))
		assertStatus(t, resp, http.StatusNoContent)

		resp2 := do(t, srv, newRequest(t, http.MethodGet, "/users/2", nil))
		assertStatus(t, resp2, http.StatusNotFound)
	})

	t.Run("POST with malformed JSON returns 400", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{bad json`))
		req.Header.Set("Content-Type", "application/json")
		resp := do(t, srv, req)
		assertStatus(t, resp, http.StatusBadRequest)
	})
}

// TestServer_Controller_GroupedRoutes verifies that a Controller groups its
// Resources under the controller path prefix.
func TestServer_Controller_GroupedRoutes(t *testing.T) {
	products := map[int]Product{
		1: {ID: 1, Name: "Widget", Price: 9.99},
	}
	seq := 2

	ctrl := fiberext.NewController("/api/v1")
	ctrl.AddResource(fiberext.NewResource(http.MethodGet, "/products", func(c fiberext.Context) error {
		list := make([]Product, 0, len(products))
		for _, p := range products {
			list = append(list, p)
		}
		return c.JSON(list)
	}))
	ctrl.AddResource(fiberext.NewResource(http.MethodPost, "/products", func(c fiberext.Context) error {
		p, err := fiberext.FromBody[Product](c)
		if err != nil {
			return fiberext.BadRequest(c)
		}
		p.ID = seq
		seq++
		products[p.ID] = p
		return c.Status(http.StatusCreated).JSON(p)
	}))
	ctrl.AddResource(fiberext.NewResource(http.MethodGet, "/products/:id", func(c fiberext.Context) error {
		type params struct {
			ID int `uri:"id"`
		}
		ps, _ := fiberext.FromParams[params](c)
		p, ok := products[ps.ID]
		if !ok {
			return fiberext.NotFound(c)
		}
		return c.JSON(p)
	}))
	ctrl.AddResource(fiberext.NewResource(http.MethodDelete, "/products/:id", func(c fiberext.Context) error {
		type params struct {
			ID int `uri:"id"`
		}
		ps, _ := fiberext.FromParams[params](c)
		delete(products, ps.ID)
		return fiberext.NoContent(c)
	}))

	srv := startServer(t, fiberext.NewConfig().WithController(ctrl))

	t.Run("GET /api/v1/products", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/api/v1/products", nil))
		assertStatus(t, resp, http.StatusOK)
		list := decode[[]Product](t, resp.Body)
		if len(list) != 1 {
			t.Fatalf("want 1 product, got %d", len(list))
		}
	})

	t.Run("POST /api/v1/products creates under prefix", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodPost, "/api/v1/products", Product{Name: "Gadget", Price: 29.99}))
		assertStatus(t, resp, http.StatusCreated)
		p := decode[Product](t, resp.Body)
		if p.ID == 0 || p.Name != "Gadget" {
			t.Fatalf("unexpected product: %+v", p)
		}
	})

	t.Run("GET /api/v1/products/:id", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/api/v1/products/1", nil))
		assertStatus(t, resp, http.StatusOK)
		p := decode[Product](t, resp.Body)
		if p.ID != 1 {
			t.Fatalf("want ID 1, got %d", p.ID)
		}
	})

	t.Run("routes without prefix return 404", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/products", nil))
		assertStatus(t, resp, http.StatusNotFound)
	})

	t.Run("DELETE /api/v1/products/:id then 404", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodDelete, "/api/v1/products/1", nil))
		assertStatus(t, resp, http.StatusNoContent)

		resp2 := do(t, srv, newRequest(t, http.MethodGet, "/api/v1/products/1", nil))
		assertStatus(t, resp2, http.StatusNotFound)
	})
}

// TestServer_Middleware_AuthGate verifies that a bearer-token middleware blocks
// unauthenticated requests and allows authenticated ones through.
func TestServer_Middleware_AuthGate(t *testing.T) {
	const validToken = "secret-token"

	authMiddleware := func(c fiberext.Context) error {
		auth := c.Get("Authorization")
		if auth != "Bearer "+validToken {
			return fiberext.Unauthorized(c)
		}
		return c.Next()
	}

	cfg := fiberext.NewConfig().
		WithMiddleware(authMiddleware).
		WithResource(fiberext.NewResource(http.MethodGet, "/profile", func(c fiberext.Context) error {
			return c.JSON(fiber.Map{"user": "alice"})
		}))

	srv := startServer(t, cfg)

	t.Run("request without token is rejected with 401", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/profile", nil))
		assertStatus(t, resp, http.StatusUnauthorized)
	})

	t.Run("request with wrong token is rejected with 401", func(t *testing.T) {
		req := newRequest(t, http.MethodGet, "/profile", nil)
		req.Header.Set("Authorization", "Bearer wrong-token")
		resp := do(t, srv, req)
		assertStatus(t, resp, http.StatusUnauthorized)
	})

	t.Run("request with valid token reaches handler", func(t *testing.T) {
		req := newRequest(t, http.MethodGet, "/profile", nil)
		req.Header.Set("Authorization", "Bearer "+validToken)
		resp := do(t, srv, req)
		assertStatus(t, resp, http.StatusOK)
	})
}

// TestServer_Middleware_Order verifies that multiple middlewares execute in
// registration order and that each one can access values set by the previous.
func TestServer_Middleware_Order(t *testing.T) {
	var callOrder []string

	record := func(name string) fiberext.Handler {
		return func(c fiberext.Context) error {
			callOrder = append(callOrder, name)
			return c.Next()
		}
	}

	cfg := fiberext.NewConfig().
		WithMiddleware(record("first")).
		WithMiddleware(record("second")).
		WithMiddleware(record("third")).
		WithResource(fiberext.NewResource(http.MethodGet, "/ping", func(c fiberext.Context) error {
			callOrder = append(callOrder, "handler")
			return c.SendString("pong")
		}))

	srv := startServer(t, cfg)

	do(t, srv, newRequest(t, http.MethodGet, "/ping", nil))

	if len(callOrder) != 4 {
		t.Fatalf("expected 4 calls, got %d: %v", len(callOrder), callOrder)
	}
	expected := []string{"first", "second", "third", "handler"}
	for i, want := range expected {
		if callOrder[i] != want {
			t.Fatalf("step %d: want %q, got %q", i, want, callOrder[i])
		}
	}
}

// TestServer_Middleware_LocalsPropagation verifies that locals set by a
// middleware are readable in the handler, enabling patterns like request-scoped
// user injection after authentication.
func TestServer_Middleware_LocalsPropagation(t *testing.T) {
	const localsKey = "requestUser"

	injectUser := func(c fiberext.Context) error {
		// simulate extracting user from JWT / session
		c.Locals(localsKey, &User{ID: 42, Name: "Alice", Email: "alice@example.com"})
		return c.Next()
	}

	cfg := fiberext.NewConfig().
		WithMiddleware(injectUser).
		WithResource(fiberext.NewResource(http.MethodGet, "/me", func(c fiberext.Context) error {
			u, ok := c.Locals(localsKey).(*User)
			if !ok || u == nil {
				return fiberext.RespondInternalError(c)
			}
			return c.JSON(u)
		}))

	srv := startServer(t, cfg)

	resp := do(t, srv, newRequest(t, http.MethodGet, "/me", nil))
	assertStatus(t, resp, http.StatusOK)

	u := decode[User](t, resp.Body)
	if u.ID != 42 || u.Name != "Alice" {
		t.Fatalf("unexpected user from locals: %+v", u)
	}
}

// TestServer_Middleware_ShortCircuit verifies that when a middleware returns a
// response (e.g. 429 rate-limit), downstream middlewares and the handler are
// never called.
func TestServer_Middleware_ShortCircuit(t *testing.T) {
	const limit = 3
	var count atomic.Int32
	var handlerCalls atomic.Int32

	rateLimiter := func(c fiberext.Context) error {
		n := count.Add(1)
		if int(n) > limit {
			return fiberext.TooManyRequests(c)
		}
		return c.Next()
	}

	cfg := fiberext.NewConfig().
		WithMiddleware(rateLimiter).
		WithResource(fiberext.NewResource(http.MethodGet, "/data", func(c fiberext.Context) error {
			handlerCalls.Add(1)
			return c.SendString("ok")
		}))

	srv := startServer(t, cfg)

	for i := 0; i < limit; i++ {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/data", nil))
		assertStatus(t, resp, http.StatusOK)
	}

	// 4th request should be blocked
	resp := do(t, srv, newRequest(t, http.MethodGet, "/data", nil))
	assertStatus(t, resp, http.StatusTooManyRequests)

	if got := handlerCalls.Load(); int(got) != limit {
		t.Fatalf("handler calls: want %d, got %d", limit, got)
	}
}

// TestServer_CustomErrorHandler verifies that a custom fiber ErrorHandler
// receives errors produced by handlers and can format them uniformly.
func TestServer_CustomErrorHandler(t *testing.T) {
	errHandler := func(c fiberext.Context, err error) error {
		code := http.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		return c.Status(code).JSON(APIError{Code: code, Message: err.Error()})
	}

	cfg := fiberext.NewConfig().
		WithErrorHandler(errHandler).
		WithResource(fiberext.NewResource(http.MethodGet, "/boom", func(c fiberext.Context) error {
			return fiber.NewError(http.StatusUnprocessableEntity, "validation failed")
		})).
		WithResource(fiberext.NewResource(http.MethodGet, "/forbidden", func(c fiberext.Context) error {
			return fiber.NewError(http.StatusForbidden, "access denied")
		}))

	srv := startServer(t, cfg)

	t.Run("handler error is formatted by error handler", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/boom", nil))
		assertStatus(t, resp, http.StatusUnprocessableEntity)
		e := decode[APIError](t, resp.Body)
		if e.Code != http.StatusUnprocessableEntity {
			t.Fatalf("code: want 422, got %d", e.Code)
		}
		if e.Message != "validation failed" {
			t.Fatalf("message: want 'validation failed', got %q", e.Message)
		}
	})

	t.Run("403 error is formatted by error handler", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/forbidden", nil))
		assertStatus(t, resp, http.StatusForbidden)
		e := decode[APIError](t, resp.Body)
		if e.Code != http.StatusForbidden {
			t.Fatalf("code: want 403, got %d", e.Code)
		}
	})

	t.Run("404 for unknown route is handled", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/unknown", nil))
		assertStatus(t, resp, http.StatusNotFound)
		e := decode[APIError](t, resp.Body)
		if e.Code != http.StatusNotFound {
			t.Fatalf("code: want 404, got %d", e.Code)
		}
	})
}

// TestServer_Middleware_ResponseHeaders verifies that middleware can inject
// response headers (e.g. CORS, request tracing) for all routes.
func TestServer_Middleware_ResponseHeaders(t *testing.T) {
	corsMiddleware := func(c fiberext.Context) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Set("X-Service", "fiberext")
		return c.Next()
	}

	cfg := fiberext.NewConfig().
		WithMiddleware(corsMiddleware).
		WithResource(fiberext.NewResource(http.MethodGet, "/health", func(c fiberext.Context) error {
			return c.JSON(fiber.Map{"status": "ok"})
		})).
		WithResource(fiberext.NewResource(http.MethodPost, "/events", func(c fiberext.Context) error {
			return fiberext.Accepted(c)
		}))

	srv := startServer(t, cfg)

	for _, path := range []string{"/health", "/events"} {
		method := http.MethodGet
		if path == "/events" {
			method = http.MethodPost
		}
		resp := do(t, srv, newRequest(t, method, path, nil))

		assertHeader(t, resp, "Access-Control-Allow-Origin", "*")
		assertHeader(t, resp, "X-Service", "fiberext")
	}
}

// TestServer_Mixed_ResourcesAndControllers verifies that top-level Resources
// and grouped Controllers coexist without interfering with each other.
func TestServer_Mixed_ResourcesAndControllers(t *testing.T) {
	usersCtrl := fiberext.NewController("/api/v1")
	usersCtrl.AddResource(fiberext.NewResource(http.MethodGet, "/users", func(c fiberext.Context) error {
		return c.JSON([]User{{ID: 1, Name: "Alice"}})
	}))

	productsCtrl := fiberext.NewController("/api/v2")
	productsCtrl.AddResource(fiberext.NewResource(http.MethodGet, "/products", func(c fiberext.Context) error {
		return c.JSON([]Product{{ID: 1, Name: "Widget", Price: 9.99}})
	}))

	cfg := fiberext.NewConfig().
		WithResource(fiberext.NewResource(http.MethodGet, "/health", func(c fiberext.Context) error {
			return c.JSON(fiber.Map{"status": "ok"})
		})).
		WithController(usersCtrl).
		WithController(productsCtrl)

	srv := startServer(t, cfg)

	t.Run("top-level health resource", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/health", nil))
		assertStatus(t, resp, http.StatusOK)
	})

	t.Run("v1 users controller", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/api/v1/users", nil))
		assertStatus(t, resp, http.StatusOK)
		users := decode[[]User](t, resp.Body)
		if len(users) != 1 || users[0].Name != "Alice" {
			t.Fatalf("unexpected users: %+v", users)
		}
	})

	t.Run("v2 products controller", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/api/v2/products", nil))
		assertStatus(t, resp, http.StatusOK)
		products := decode[[]Product](t, resp.Body)
		if len(products) != 1 || products[0].Name != "Widget" {
			t.Fatalf("unexpected products: %+v", products)
		}
	})

	t.Run("cross-version routes do not leak", func(t *testing.T) {
		// v1 users prefix must not serve v2 routes
		resp := do(t, srv, newRequest(t, http.MethodGet, "/api/v1/products", nil))
		assertStatus(t, resp, http.StatusNotFound)

		resp2 := do(t, srv, newRequest(t, http.MethodGet, "/api/v2/users", nil))
		assertStatus(t, resp2, http.StatusNotFound)
	})
}

// TestServer_Middleware_PerControllerAuth demonstrates scoping authentication
// to a single controller while leaving public routes unprotected.
func TestServer_Middleware_PerControllerAuth(t *testing.T) {
	const token = "admin-token"

	// admin controller is behind a group-level middleware applied via the
	// global middleware — routes that need auth are placed in a separate
	// controller prefix.
	authMiddleware := func(c fiberext.Context) error {
		// only guard /admin/* paths
		if len(c.Path()) >= 6 && c.Path()[:6] == "/admin" {
			if c.Get("Authorization") != "Bearer "+token {
				return fiberext.Unauthorized(c)
			}
		}
		return c.Next()
	}

	adminCtrl := fiberext.NewController("/admin")
	adminCtrl.AddResource(fiberext.NewResource(http.MethodGet, "/dashboard", func(c fiberext.Context) error {
		return c.JSON(fiber.Map{"data": "sensitive"})
	}))
	adminCtrl.AddResource(fiberext.NewResource(http.MethodDelete, "/users/:id", func(c fiberext.Context) error {
		return fiberext.NoContent(c)
	}))

	cfg := fiberext.NewConfig().
		WithMiddleware(authMiddleware).
		WithResource(fiberext.NewResource(http.MethodGet, "/status", func(c fiberext.Context) error {
			return c.JSON(fiber.Map{"status": "ok"})
		})).
		WithController(adminCtrl)

	srv := startServer(t, cfg)

	t.Run("public /status is accessible without token", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/status", nil))
		assertStatus(t, resp, http.StatusOK)
	})

	t.Run("admin dashboard requires auth", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/admin/dashboard", nil))
		assertStatus(t, resp, http.StatusUnauthorized)
	})

	t.Run("admin dashboard accessible with valid token", func(t *testing.T) {
		req := newRequest(t, http.MethodGet, "/admin/dashboard", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp := do(t, srv, req)
		assertStatus(t, resp, http.StatusOK)
	})

	t.Run("admin delete requires auth", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodDelete, "/admin/users/1", nil))
		assertStatus(t, resp, http.StatusUnauthorized)
	})
}

// TestServer_Controller_NestedResources verifies typical nested resource
// patterns like /orgs/:orgID/members used in real REST APIs.
func TestServer_Controller_NestedResources(t *testing.T) {
	type Member struct {
		OrgID  int    `json:"org_id"`
		UserID int    `json:"user_id"`
		Role   string `json:"role"`
	}

	members := map[string][]Member{
		"1": {{OrgID: 1, UserID: 10, Role: "admin"}, {OrgID: 1, UserID: 11, Role: "member"}},
		"2": {{OrgID: 2, UserID: 20, Role: "owner"}},
	}

	ctrl := fiberext.NewController("/orgs")
	ctrl.AddResource(fiberext.NewResource(http.MethodGet, "/:orgID/members", func(c fiberext.Context) error {
		orgID := fiberext.GetParam(c, "orgID")
		list, ok := members[orgID]
		if !ok {
			return fiberext.NotFound(c)
		}
		return c.JSON(list)
	}))
	ctrl.AddResource(fiberext.NewResource(http.MethodPost, "/:orgID/members", func(c fiberext.Context) error {
		orgID := fiberext.GetParam(c, "orgID")
		m, err := fiberext.FromBody[Member](c)
		if err != nil {
			return fiberext.BadRequest(c)
		}
		m.OrgID, _ = strconv.Atoi(orgID)
		members[orgID] = append(members[orgID], m)
		return c.Status(http.StatusCreated).JSON(m)
	}))

	srv := startServer(t, fiberext.NewConfig().WithController(ctrl))

	t.Run("list members of org 1", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/orgs/1/members", nil))
		assertStatus(t, resp, http.StatusOK)
		list := decode[[]Member](t, resp.Body)
		if len(list) != 2 {
			t.Fatalf("want 2 members, got %d", len(list))
		}
	})

	t.Run("list members of org 2", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/orgs/2/members", nil))
		assertStatus(t, resp, http.StatusOK)
		list := decode[[]Member](t, resp.Body)
		if len(list) != 1 || list[0].Role != "owner" {
			t.Fatalf("unexpected members: %+v", list)
		}
	})

	t.Run("unknown org returns 404", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/orgs/999/members", nil))
		assertStatus(t, resp, http.StatusNotFound)
	})

	t.Run("add member to org 1", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodPost, "/orgs/1/members", Member{UserID: 99, Role: "viewer"}))
		assertStatus(t, resp, http.StatusCreated)
		m := decode[Member](t, resp.Body)
		if m.OrgID != 1 || m.UserID != 99 {
			t.Fatalf("unexpected member: %+v", m)
		}
	})
}

// TestServer_QueryParams verifies FromQuery (structured query parsing) and
// GetQueryArg (single-key access) in realistic filtering and sorting scenarios.
func TestServer_QueryParams(t *testing.T) {
	type UserFilter struct {
		Name  string `query:"name"`
		Email string `query:"email"`
	}

	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Alice", Email: "alice2@example.com"},
	}

	cfg := fiberext.NewConfig().
		// FromQuery: filter users by name and/or email struct fields
		WithResource(fiberext.NewResource(http.MethodGet, "/users", func(c fiberext.Context) error {
			f, err := fiberext.FromQuery[UserFilter](c)
			if err != nil {
				return fiberext.BadRequest(c)
			}
			result := make([]User, 0)
			for _, u := range users {
				if f.Name != "" && !strings.EqualFold(u.Name, f.Name) {
					continue
				}
				if f.Email != "" && !strings.EqualFold(u.Email, f.Email) {
					continue
				}
				result = append(result, u)
			}
			return c.JSON(result)
		})).
		// GetQueryArg: sort direction as a plain single query arg
		WithResource(fiberext.NewResource(http.MethodGet, "/products", func(c fiberext.Context) error {
			sort := fiberext.GetQueryArg(c, "sort", "asc")
			products := []Product{
				{ID: 1, Name: "Alpha", Price: 5.0},
				{ID: 2, Name: "Beta", Price: 1.0},
				{ID: 3, Name: "Gamma", Price: 9.0},
			}
			if sort == "desc" {
				for i, j := 0, len(products)-1; i < j; i, j = i+1, j-1 {
					products[i], products[j] = products[j], products[i]
				}
			}
			return c.JSON(products)
		})).
		// GetQueryArg with default: pagination using page param, defaults to "1"
		WithResource(fiberext.NewResource(http.MethodGet, "/items", func(c fiberext.Context) error {
			page := fiberext.GetQueryArg(c, "page", "1")
			return c.JSON(fiber.Map{"page": page})
		}))

	srv := startServer(t, cfg)

	t.Run("FromQuery: no filter returns all users", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users", nil))
		assertStatus(t, resp, http.StatusOK)
		result := decode[[]User](t, resp.Body)
		if len(result) != 3 {
			t.Fatalf("want 3 users, got %d", len(result))
		}
	})

	t.Run("FromQuery: filter by name returns matching users", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users?name=Alice", nil))
		assertStatus(t, resp, http.StatusOK)
		result := decode[[]User](t, resp.Body)
		if len(result) != 2 {
			t.Fatalf("want 2 Alices, got %d", len(result))
		}
		for _, u := range result {
			if !strings.EqualFold(u.Name, "Alice") {
				t.Fatalf("unexpected user in result: %+v", u)
			}
		}
	})

	t.Run("FromQuery: filter by name and email returns exact match", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users?name=Alice&email=alice@example.com", nil))
		assertStatus(t, resp, http.StatusOK)
		result := decode[[]User](t, resp.Body)
		if len(result) != 1 || result[0].ID != 1 {
			t.Fatalf("want only Alice ID=1, got %+v", result)
		}
	})

	t.Run("FromQuery: filter with no match returns empty list", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/users?name=Nobody", nil))
		assertStatus(t, resp, http.StatusOK)
		result := decode[[]User](t, resp.Body)
		if len(result) != 0 {
			t.Fatalf("want empty list, got %+v", result)
		}
	})

	t.Run("GetQueryArg: sort=asc returns products in original order", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/products?sort=asc", nil))
		assertStatus(t, resp, http.StatusOK)
		result := decode[[]Product](t, resp.Body)
		if result[0].Name != "Alpha" {
			t.Fatalf("want Alpha first, got %s", result[0].Name)
		}
	})

	t.Run("GetQueryArg: sort=desc returns products reversed", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/products?sort=desc", nil))
		assertStatus(t, resp, http.StatusOK)
		result := decode[[]Product](t, resp.Body)
		if result[0].Name != "Gamma" {
			t.Fatalf("want Gamma first, got %s", result[0].Name)
		}
	})

	t.Run("GetQueryArg: default used when param absent", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/items", nil))
		assertStatus(t, resp, http.StatusOK)
		m := decode[map[string]string](t, resp.Body)
		if m["page"] != "1" {
			t.Fatalf("want default page=1, got %q", m["page"])
		}
	})

	t.Run("GetQueryArg: explicit param overrides default", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/items?page=3", nil))
		assertStatus(t, resp, http.StatusOK)
		m := decode[map[string]string](t, resp.Body)
		if m["page"] != "3" {
			t.Fatalf("want page=3, got %q", m["page"])
		}
	})
}

// TestServer_Middleware_PanicRecovery verifies that the server continues to
// serve subsequent requests after a handler panics (via fiber's built-in
// recover middleware).
func TestServer_Middleware_PanicRecovery(t *testing.T) {
	var recovered atomic.Bool

	recoverMiddleware := func(c fiberext.Context) error {
		defer func() {
			if r := recover(); r != nil {
				recovered.Store(true)
				_ = c.Status(http.StatusInternalServerError).SendString("recovered")
			}
		}()
		return c.Next()
	}

	cfg := fiberext.NewConfig().
		WithMiddleware(recoverMiddleware).
		WithResource(fiberext.NewResource(http.MethodGet, "/panic", func(c fiberext.Context) error {
			panic("something went wrong")
		})).
		WithResource(fiberext.NewResource(http.MethodGet, "/ok", func(c fiberext.Context) error {
			return c.SendString("alive")
		}))

	srv := startServer(t, cfg)

	t.Run("panic is recovered and returns 500", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/panic", nil))
		assertStatus(t, resp, http.StatusInternalServerError)
		if !recovered.Load() {
			t.Fatal("expected recovery middleware to have run")
		}
	})

	t.Run("server handles next request normally after panic", func(t *testing.T) {
		resp := do(t, srv, newRequest(t, http.MethodGet, "/ok", nil))
		assertStatus(t, resp, http.StatusOK)
		body, _ := io.ReadAll(resp.Body)
		if string(body) != "alive" {
			t.Fatalf("want 'alive', got %q", body)
		}
	})
}
