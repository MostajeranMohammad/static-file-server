package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/MostajeranMohammad/static-file-server/internal/contracts/consts"
	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/MostajeranMohammad/static-file-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type StaticFileManagerUsecase struct {
	objectStorageManagerUseCase      ObjectStorageManager
	staticFileMetaDataManagerUseCase StaticFileMetaDataManager
	imageOptimizerUseCase            ImageOptimizer
}

func NewStaticFileManagerUsecase(
	objectStorageManagerUseCase ObjectStorageManager,
	staticFileMetaDataManagerUseCase StaticFileMetaDataManager,
	imageOptimizerUseCase ImageOptimizer,
) StaticFileManager {
	return &StaticFileManagerUsecase{objectStorageManagerUseCase, staticFileMetaDataManagerUseCase, imageOptimizerUseCase}
}

func (sf StaticFileManagerUsecase) GetFile(ctx context.Context, fileName string, userClaims jwt.MapClaims) (io.Reader, func() error, error) {
	bucketName, err := utils.GetBucketNameFromFileName(fileName)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	if consts.PrivateBuckets[bucketName] {
		access, err := sf.staticFileMetaDataManagerUseCase.CheckFileMetaDataForAccess(ctx, fileName, userClaims["user_id"].(int32))
		if err != nil {
			return nil, func() error { return nil }, err
		}
		if !access {
			return nil, func() error { return nil }, fiber.NewError(403, "Access denied")
		}
	}
	return sf.objectStorageManagerUseCase.GetObject(ctx, bucketName, fileName)
}

func (sf StaticFileManagerUsecase) SaveFile(
	ctx context.Context,
	bucketName string,
	uploaderId uint,
	files []entity.FormFile,
	UserIdsWhoAccessThisFile []int32,
) ([]entity.StaticFileMetaData, error) {

	type insertResultType struct {
		InsertedFileMetaData entity.StaticFileMetaData
		err                  error
	}
	insertResult := make(chan insertResultType, len(files))

	for _, f := range files {
		file := &f
		oldFileName := file.FileName
		newFileName := utils.GenerateFileName(bucketName, consts.WEBP)
		go func() {
			needOptimization, err := sf.imageOptimizerUseCase.CheckImageNeedsOptimization(file.Buffer, consts.MediumImage)
			if err != nil {
				insertResult <- insertResultType{err: err}
				return
			}
			if needOptimization {
				if err != nil {
					insertResult <- insertResultType{err: err}
					return
				}
				file.Buffer, err = sf.imageOptimizerUseCase.OptimizeImage(file.Buffer, consts.MediumImage, consts.WEBP)
				if err != nil {
					insertResult <- insertResultType{err: err}
					return
				}
			}

			_, err = sf.objectStorageManagerUseCase.SaveObject(ctx, bucketName, newFileName, bytes.NewReader(file.Buffer), int64(bytes.NewReader(file.Buffer).Len()))
			if err != nil {
				insertResult <- insertResultType{err: err}
				return
			}

			metaData, err := sf.staticFileMetaDataManagerUseCase.SaveFileMetaDataOnDB(ctx, newFileName, oldFileName, bucketName, uploaderId, UserIdsWhoAccessThisFile)
			if err != nil {
				insertResult <- insertResultType{err: err}
				return
			}

			insertResult <- insertResultType{err: nil, InsertedFileMetaData: metaData}
		}()
	}
	defer close(insertResult)

	finalResult := make([]entity.StaticFileMetaData, len(files))
	finalError := ""
	for i := 0; i < len(files); i++ {
		goroutineResult := <-insertResult
		if goroutineResult.err != nil {
			finalError += fmt.Sprintf("%s\n", goroutineResult.err.Error())
		} else {
			finalResult[i] = goroutineResult.InsertedFileMetaData
		}
	}
	var err error = nil
	if finalError != "" {
		err = errors.New(finalError)
	}
	return finalResult, err
}

func (sf StaticFileManagerUsecase) DeleteFile(ctx context.Context, fileName string, userClaims jwt.MapClaims) (entity.StaticFileMetaData, error) {
	bucketName, err := utils.GetBucketNameFromFileName(fileName)
	if err != nil {
		return entity.StaticFileMetaData{}, err
	}

	err = sf.objectStorageManagerUseCase.DeleteObject(ctx, bucketName, fileName)
	if err != nil {
		return entity.StaticFileMetaData{}, err
	}

	meta, err := sf.staticFileMetaDataManagerUseCase.DeleteFile(ctx, fileName)
	if err != nil {
		return entity.StaticFileMetaData{}, err
	}
	return meta, nil
}
