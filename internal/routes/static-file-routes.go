package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MostajeranMohammad/static-file-server/internal/controller"
	"github.com/MostajeranMohammad/static-file-server/pkg/guards"
)

func NewStaticFileRouter(controller controller.StaticFile, jwtGuard guards.JWT) *fiber.App {
	router := fiber.New()

	router.Get("/download/:file_name", jwtGuard.GetOptionalJWTGuard(), controller.Download)
	router.Post("/upload/:bucket_name", jwtGuard.GetStrictJWTGuard(), controller.Upload)

	return router
}
