package controller

import (
	"bytes"
	"fmt"
	"io"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/MostajeranMohammad/static-file-server/internal/usecase"
	"github.com/MostajeranMohammad/static-file-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type staticFileController struct {
	useCase usecase.StaticFileManager
}

func NewStaticFileController(useCase usecase.StaticFileManager) StaticFile {
	return &staticFileController{
		useCase: useCase,
	}
}

func (s *staticFileController) Upload(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint((claims["user_id"]).(float64))
	bucketName := c.Params("bucket_name")

	if form, err := c.MultipartForm(); err == nil {
		err = utils.ValidateAllParamsAreIntString(form.Value["user_ids_who_access_this_file"])
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		UserIdsWhoAccessThisFile := utils.ConvertStringArrayToIntArray(form.Value["user_ids_who_access_this_file"])

		formFiles := form.File["file"]
		fmt.Println(map[string]interface{}{"name": formFiles[0].Filename, "size": formFiles[0].Size})
		if len(formFiles) < 1 {
			return fiber.NewError(fiber.StatusBadRequest, "files not found on request.")
		}

		files := []entity.FormFile{}

		for _, f := range formFiles {
			multipartFile, err := f.Open()
			if err != nil {
				return fiber.NewError(500, err.Error())
			}
			defer multipartFile.Close()

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, multipartFile); err != nil {
				return fiber.NewError(500, "error in coping file buffer")
			}

			files = append(files, entity.FormFile{FileName: f.Filename, FileSize: f.Size, Buffer: buf.Bytes()})
		}

		info, err := s.useCase.SaveFile(c.Context(), bucketName, uint(userId), files, UserIdsWhoAccessThisFile)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(entity.ResponseModel{
			Successful: true,
			Message:    "success",
			Data:       info,
			Meta:       nil,
		})
	} else {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
}

func (s *staticFileController) Download(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fileName := c.Params("file_name")

	file, err := s.useCase.GetFile(c.Context(), fileName, claims)
	if err != nil {
		return err
	}

	buffer := []byte{}
	file.Read(buffer)
	_, err = c.Write(buffer)
	return err
}
