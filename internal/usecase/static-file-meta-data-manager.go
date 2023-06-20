package usecase

import (
	"context"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"github.com/MostajeranMohammad/static-file-server/internal/repo"
	"github.com/MostajeranMohammad/static-file-server/pkg/utils"
	"gorm.io/gorm/clause"
)

type StaticFileMetaManagerUsecase struct {
	r repo.StaticFileMetaData
}

func NewStaticFileMetaManagerUsecase(r repo.StaticFileMetaData) StaticFileMetaDataManager {
	return &StaticFileMetaManagerUsecase{r}
}

func (sfm *StaticFileMetaManagerUsecase) SaveFileMetaDataOnDB(
	ctx context.Context, fileName string,
	oldFileName string, bucketName string,
	uploader uint, userIdsWhoAccessThisFile []int32,
) (entity.StaticFileMetaData, error) {
	return sfm.r.Create(ctx, entity.StaticFileMetaData{
		FileName:                 fileName,
		OriginalFileName:         oldFileName,
		BucketName:               bucketName,
		UploaderId:               uint(uploader),
		UserIdsWhoAccessThisFile: userIdsWhoAccessThisFile,
	})
}

func (sfm *StaticFileMetaManagerUsecase) GetFileMetaDataByFileName(ctx context.Context, fileName string) (entity.StaticFileMetaData, error) {
	return sfm.r.GetByFileName(ctx, fileName)
}

func (sfm *StaticFileMetaManagerUsecase) GetAllFilesMetaDataByAdmin(ctx context.Context, bucketName string, limit int, skip int, uploaderId int) ([]entity.StaticFileMetaData, error) {
	filter := clause.AndConditions{}
	if bucketName != "" {
		filter.Exprs = append(filter.Exprs, clause.Eq{Column: "bucket_name", Value: bucketName})
	}
	if uploaderId != 0 {
		filter.Exprs = append(filter.Exprs, clause.Eq{Column: "uploader_id", Value: uploaderId})
	}

	return sfm.r.GetAll(ctx, filter, skip, limit)
}

func (sfm *StaticFileMetaManagerUsecase) CheckFileMetaDataForAccess(ctx context.Context, objectName string, userId int32) (bool, error) {
	metaData, err := sfm.r.GetFileAccessData(ctx, objectName)
	if err != nil {
		return false, err
	}

	if int32(metaData.UploaderId) == userId ||
		utils.Int32SliceContains(metaData.UserIdsWhoAccessThisFile, userId) {
		return true, nil
	}
	return false, nil
}

func (sfm *StaticFileMetaManagerUsecase) CountUploadedImagesByAdmin(ctx context.Context, bucketName string, uploaderId int) (int64, error) {
	filter := clause.AndConditions{}
	if bucketName != "" {
		filter.Exprs = append(filter.Exprs, clause.Eq{Column: "bucket_name", Value: bucketName})
	}
	if uploaderId != 0 {
		filter.Exprs = append(filter.Exprs, clause.Eq{Column: "uploader_id", Value: uploaderId})
	}

	return sfm.r.CountFiles(ctx, filter)
}

func (sfm *StaticFileMetaManagerUsecase) UpdateByFileName(ctx context.Context, fileName string, ids []int32) (entity.StaticFileMetaData, error) {
	return sfm.r.UpdateByFileName(ctx, fileName, ids)
}

func (sfm *StaticFileMetaManagerUsecase) DeleteFile(ctx context.Context, fileName string) (entity.StaticFileMetaData, error) {
	return sfm.r.DeleteByFileName(ctx, fileName)
}
