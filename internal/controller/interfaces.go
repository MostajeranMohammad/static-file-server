package controller

import "github.com/gofiber/fiber/v2"

type (
	StaticFile interface {
		Upload(c *fiber.Ctx) error
		Download(c *fiber.Ctx) error
	}

	FileMetaData interface {
		GetFilesMetaData(c *fiber.Ctx) error
		GetFileMetaDataByFileName(c *fiber.Ctx) error
		UpdateFileAccess(c *fiber.Ctx) error
		DeleteFile(c *fiber.Ctx) error
	}
)
