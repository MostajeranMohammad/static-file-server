package routes

import (
	"github.com/MostajeranMohammad/static-file-server/internal/controller"
	"github.com/MostajeranMohammad/static-file-server/pkg/guards"
	"github.com/gofiber/fiber/v2"
)

func NewFileMetaDataRoutes(controller controller.FileMetaData, jwtGuard guards.JWT) *fiber.App {
	router := fiber.New()

	router.Get("/", jwtGuard.GetStrictJWTGuard(), controller.GetFilesMetaData)
	router.Get("/:file_name", jwtGuard.GetStrictJWTGuard(), controller.GetFileMetaDataByFileName)
	router.Put("/update-file-access/:file_name", jwtGuard.GetStrictJWTGuard(), controller.UpdateFileAccess)
	router.Delete("/delete-file/:file_name", jwtGuard.GetStrictJWTGuard(), controller.DeleteFile)

	return router
}
