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

func (mc *fileMetaDataController) GetFilesMetaData(c *fiber.Ctx) error {
	skip, limit, uploaderId, bucketName :=
		c.QueryInt("skip"), c.QueryInt("limit"), c.QueryInt("uploader_id"), c.Query("bucket-name")

	data, err := mc.metaDataUsecase.GetAllFilesMetaDataByAdmin(c.Context(), bucketName, limit, skip, uploaderId)
	if err != nil {
		return err
	}

	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}

func (mc *fileMetaDataController) GetFileMetaDataByFileName(c *fiber.Ctx) error {
	fileName := c.Params("file_name")
	data, err := mc.metaDataUsecase.GetFileMetaDataByFileName(c.Context(), fileName)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}

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

	bodyMap, err := utils.ParseBodyToMap(c.Body())
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	data, err := mc.metaDataUsecase.UpdateByFileName(c.Context(), fileName, bodyMap)
	if err != nil {
		return err
	}

	return c.JSON(entity.ResponseModel{Successful: true, Data: data})
}
