package dto

import (
	"fmt"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/go-playground/validator/v10"
)

type StaticFileMetaDataOutputDto struct {
	ID                       uint
	FileName                 string
	BucketName               string
	Url                      string
	UploaderId               uint
	UserIdsWhoAccessThisFile []int32
	CreatedAt                string
	UpdatedAt                string
}

func NewStaticFileMetaDataOutputDto(e entity.StaticFileMetaData) StaticFileMetaDataOutputDto {
	return StaticFileMetaDataOutputDto{
		ID:                       e.ID,
		FileName:                 e.FileName,
		BucketName:               e.BucketName,
		Url:                      fmt.Sprintf("/download/:%s", e.FileName),
		UploaderId:               e.UploaderId,
		UserIdsWhoAccessThisFile: []int32(e.UserIdsWhoAccessThisFile),
		CreatedAt:                e.CreatedAt.String(),
		UpdatedAt:                e.UpdatedAt.String(),
	}
}

type StaticFileMetaDataUpdateAccessDto struct {
	UserIdsWhoAccessThisFile []int32 `validate:"required,dive,gt=0"`
}

func ValidateStaticFileMetaDataUpdateAccessDto(d StaticFileMetaDataUpdateAccessDto) error {
	validate := validator.New()
	return validate.Struct(d)
}
