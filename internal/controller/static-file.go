package controller

import (
	"bytes"
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

// @Security BearerAuth
// @Produce      json
// @Param        bucket_name  path  string  true "no comment"
// @Param        user_ids_who_access_this_file  formData  integer  false "no comment"
// @Param        file  formData  file  true "no comment"
// @Router       /static-file/upload/{bucket_name} [post]
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

// @Produce      json
// @Param        authorization  header  string  false "for download private objects put you jwt here."
// @Param        file_name  path  string  true  "no comment"
// @Router       /static-file/download/{file_name} [get]
func (s *staticFileController) Download(c *fiber.Ctx) error {
	var user *jwt.Token
	var claims jwt.MapClaims
	if c.Locals("user") != nil {
		user = c.Locals("user").(*jwt.Token)
		claims = user.Claims.(jwt.MapClaims)
	}
	fileName := c.Params("file_name")

	file, closeObjectFn, err := s.useCase.GetFile(c.Context(), fileName, claims)
	defer closeObjectFn()
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename="+fileName)
	_, err = io.Copy(c, file)
	return err
}
