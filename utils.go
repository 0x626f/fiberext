package fiberext

import (
	"github.com/gofiber/fiber/v2"
)

func FromParams[T any](ctx Context) (T, error) {
	var obj T
	err := ctx.ParamsParser(&obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func FromBody[T any](ctx Context) (T, error) {
	var obj T

	if len(ctx.Body()) <= 0 {
		return obj, nil
	}

	err := ctx.BodyParser(&obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func FromQuery[T any](ctx Context) (T, error) {
	var obj T
	err := ctx.QueryParser(&obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func GetParam(ctx Context, key string, def ...string) string {
	return ctx.Params(key, def...)
}

func GetQueryArg(ctx Context, key string, def ...string) string {
	return ctx.Query(key, def...)
}

func Respond(ctx Context, code int, obj ...any) error {
	ctx.Status(code)

	if len(obj) > 0 {
		return ctx.JSON(obj[0])
	}

	return nil
}

func RespondError(ctx Context, code int, obj ...any) error {
	ctx.Status(code)
	if len(obj) > 0 {
		if err := ctx.JSON(obj[0]); err != nil {
			return err
		}
	}

	return fiber.NewError(code)
}
