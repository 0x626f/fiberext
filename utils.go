package fiberext

import (
	"github.com/gofiber/fiber/v3"
)

// FromParams parses URL parameters into T using the "uri" struct tag.
func FromParams[T any](ctx Context) (T, error) {
	var obj T
	err := ctx.Bind().URI(&obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

// FromBody decodes the request body into T based on Content-Type.
func FromBody[T any](ctx Context) (T, error) {
	var obj T

	if !ctx.HasBody() {
		return obj, nil
	}

	err := ctx.Bind().Body(&obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

// FromQuery parses query string parameters into T using the "query" struct tag.
func FromQuery[T any](ctx Context) (T, error) {
	var obj T
	err := ctx.Bind().Query(&obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

// GetParam returns the URL parameter for key, falling back to def if absent.
func GetParam(ctx Context, key string, def ...string) string {
	return ctx.Params(key, def...)
}

// GetQueryArg returns the query argument for key, falling back to def if absent.
func GetQueryArg(ctx Context, key string, def ...string) string {
	return ctx.Query(key, def...)
}

// Respond writes status and optionally serializes obj as JSON.
func Respond(ctx Context, code int, obj ...any) error {
	ctx.Status(code)

	if len(obj) > 0 {
		return ctx.JSON(obj[0])
	}

	return nil
}

// RespondError writes status, optionally serializes obj, and returns a fiber.Error.
func RespondError(ctx Context, code int, obj ...any) error {
	ctx.Status(code)
	if len(obj) > 0 {
		if err := ctx.JSON(obj[0]); err != nil {
			return err
		}
	}

	return fiber.NewError(code)
}
