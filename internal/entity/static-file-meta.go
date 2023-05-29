package entity

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type StaticFileMetaData struct {
	gorm.Model
	FileName                 string `gorm:"index;unique;type:varchar(255)"`
	OriginalFileName         string `gorm:"type:varchar(255)"`
	BucketName               string `gorm:"type:varchar(255)"`
	UploaderId               uint
	UserIdsWhoAccessThisFile pq.Int32Array `gorm:"type:INTEGER[]"`
}
