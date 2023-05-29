package utils

import (
	"errors"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/MostajeranMohammad/static-file-server/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func FiberErrorHandler(logger logger.Interface) func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}
		logger.Error(err.Error(), err)

		if code != fiber.StatusInternalServerError {
			return ctx.Status(code).JSON(entity.ResponseModel{Successful: false, Message: "Internal Server Error"})
		}
		return ctx.Status(code).JSON(entity.ResponseModel{Successful: false, Message: err.Error()})
	}
}
