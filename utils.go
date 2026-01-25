package fiberext

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FromParams[T any](c *fiber.Ctx) (*T, error) {
	obj := new(T)
	err := c.ParamsParser(obj)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func FromBody[T any](c *fiber.Ctx) (*T, error) {
	obj := new(T)

	if len(c.Body()) <= 0 {
		return obj, nil
	}

	err := c.BodyParser(obj)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func FromQuery[T any](c *fiber.Ctx) (*T, error) {
	obj := new(T)
	err := c.QueryParser(obj)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func RespondBadRequest(c *fiber.Ctx, errors ...string) error {
	var err error
	c.Status(fiber.StatusBadRequest)

	if len(errors) > 0 {
		err = c.JSON(fiber.Map{
			"error": strings.Join(errors, "."),
		})
	}
	return err
}

func RespondInternalError(c *fiber.Ctx, errors ...string) error {
	var err error
	c.Status(fiber.StatusInternalServerError)

	if len(errors) > 0 {
		err = c.JSON(fiber.Map{
			"error": strings.Join(errors, "."),
		})
	}
	return err
}

func RespondCreated(c *fiber.Ctx, obj interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(obj)
}

func RespondOK(c *fiber.Ctx, obj interface{}) error {
	return c.Status(fiber.StatusOK).JSON(obj)
}
