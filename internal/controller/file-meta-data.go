package controller

import (
	"github.com/MostajeranMohammad/static-file-server/internal/dto"
	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/MostajeranMohammad/static-file-server/internal/usecase"
	"github.com/MostajeranMohammad/static-file-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type fileMetaDataController struct {
	metaDataUsecase          usecase.StaticFileMetaDataManager
	staticFileManagerUsecase usecase.StaticFileManager
}

func NewFileMetaDataController(metaDataUsecase usecase.StaticFileMetaDataManager,
	staticFileManagerUsecase usecase.StaticFileManager) FileMetaData {
	return &fileMetaDataController{metaDataUsecase, staticFileManagerUsecase}
}

// @Security BearerAuth
// @Produce      json
// @Param        limit  query  int  false "no comment"
// @Param        skip  query  int  false "no comment"
// @Param        uploader_id  query  int  false "no comment"
// @Param        bucket_name  query  string  false "no comment"
// @Router       /file-meta-data/ [get]
func (mc *fileMetaDataController) GetFilesMetaData(c *fiber.Ctx) error {
	skip, limit, uploaderId, bucketName :=
		c.QueryInt("skip"), c.QueryInt("limit"), c.QueryInt("uploader_id"), c.Query("bucket_name")

	data, err := mc.metaDataUsecase.GetAllFilesMetaDataByAdmin(c.Context(), bucketName, limit, skip, uploaderId)
	if err != nil {
		return err
	}

	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}

// @Security BearerAuth
// @Produce      json
// @Param        file_name  path  string  true "no comment"
// @Router       /file-meta-data/{file_name} [get]
func (mc *fileMetaDataController) GetFileMetaDataByFileName(c *fiber.Ctx) error {
	fileName := c.Params("file_name")
	data, err := mc.metaDataUsecase.GetFileMetaDataByFileName(c.Context(), fileName)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}

// @Security BearerAuth
// @Produce      json
// @Param        file_name  path  string  true "no comment"
// @Param        body  body  dto.StaticFileMetaDataUpdateAccessDto  true "no comment"
// @Router       /file-meta-data/update-file-access/{file_name} [put]
func (mc *fileMetaDataController) UpdateFileAccess(c *fiber.Ctx) error {
	fileName := c.Params("file_name")

	body := dto.StaticFileMetaDataUpdateAccessDto{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = utils.ValidateDto(body)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	data, err := mc.metaDataUsecase.UpdateByFileName(c.Context(), fileName, body.UserIdsWhoAccessThisFile)
	if err != nil {
		return err
	}

	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}

// @Security BearerAuth
// @Produce      json
// @Param        file_name  path  string  true "no comment"
// @Router       /file-meta-data/delete-file/{file_name} [delete]
func (mc *fileMetaDataController) DeleteFile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fileName := c.Params("file_name")

	data, err := mc.staticFileManagerUsecase.DeleteFile(c.Context(), fileName, claims)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}
