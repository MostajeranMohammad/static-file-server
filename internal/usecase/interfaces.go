package usecase

import (
	"context"
	"io"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	minio "github.com/minio/minio-go/v7"
)

type (
	StaticFileManager interface {
		GetFile(ctx context.Context, fileName string, userClaims jwt.MapClaims) (io.Reader, func() error, error)
		SaveFile(
			ctx context.Context,
			bucketName string,
			uploaderId uint,
			files []entity.FormFile,
			UserIdsWhoAccessThisFile []int32,
		) ([]entity.StaticFileMetaData, error)
		DeleteFile(ctx context.Context, fileName string, userClaims jwt.MapClaims) (entity.StaticFileMetaData, error)
	}

	ObjectStorageManager interface {
		InitBuckets() error
		GetObject(ctx context.Context, bucketName string, objectName string) (io.Reader, func() error, error)
		SaveObject(ctx context.Context, bucketName string, fileName string, file io.Reader, fileSize int64) (minio.UploadInfo, error)
		DeleteObject(ctx context.Context, bucketName string, fileName string) error
	}

	StaticFileMetaDataManager interface {
		SaveFileMetaDataOnDB(ctx context.Context, fileName string, oldFileName string, bucketName string, uploader uint, userIdsWhoAccessThisFile []int32) (entity.StaticFileMetaData, error)
		GetFileMetaDataByFileName(ctx context.Context, fileName string) (entity.StaticFileMetaData, error)
		GetAllFilesMetaDataByAdmin(ctx context.Context, bucketName string, limit int, skip int, uploaderId int) ([]entity.StaticFileMetaData, error)
		CheckFileMetaDataForAccess(ctx context.Context, objectName string, userId int32) (bool, error)
		CountUploadedImagesByAdmin(ctx context.Context, bucketName string, uploaderId int) (int64, error)
		UpdateByFileName(ctx context.Context, fileName string, ids []int32) (entity.StaticFileMetaData, error)
		DeleteFile(ctx context.Context, fileName string) (entity.StaticFileMetaData, error)
	}

	ImageOptimizer interface {
		CheckImageNeedsOptimization(buffer []byte, imageSize int) (bool, error)
		OptimizeImage(buffer []byte, imageSize int, outputFileType int) ([]byte, error)
	}
)
