package repo

import (
	"context"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StaticFileMetaDataRepo struct {
	db *gorm.DB
}

func NewStaticFileMetaRepo(db *gorm.DB) StaticFileMetaData {
	return &StaticFileMetaDataRepo{db}
}

func (sr *StaticFileMetaDataRepo) Create(ctx context.Context, data entity.StaticFileMetaData) (entity.StaticFileMetaData, error) {
	result := sr.db.WithContext(ctx).Clauses(clause.Returning{}).Create(&data)
	return data, result.Error
}

func (sr *StaticFileMetaDataRepo) GetByFileName(ctx context.Context, fileName string) (entity.StaticFileMetaData, error) {
	fileMeta := entity.StaticFileMetaData{}
	result := sr.db.WithContext(ctx).Where("FileName = ?", fileName).First(&fileMeta)
	return fileMeta, result.Error
}

func (sr *StaticFileMetaDataRepo) GetAll(ctx context.Context, filter clause.AndConditions, skip int, limit int) ([]entity.StaticFileMetaData, error) {
	filesMeta := []entity.StaticFileMetaData{}
	query := sr.db.Limit(limit).Offset(skip)
	if len(filter.Exprs) > 0 {
		query.Where(filter)
	}
	result := query.Find(&filesMeta)

	return filesMeta, result.Error
}

func (sr *StaticFileMetaDataRepo) GetFileAccessData(ctx context.Context, fileName string) (entity.StaticFileMetaData, error) {
	fileMeta := entity.StaticFileMetaData{}
	result := sr.db.WithContext(ctx).Where("FileName = ?", fileName).Select("UploaderId", "UserIdsWhoAccessThisFile").First(&fileMeta)
	return fileMeta, result.Error
}

func (sr *StaticFileMetaDataRepo) CountFiles(ctx context.Context, filter clause.AndConditions) (int64, error) {
	var count int64
	query := sr.db.WithContext(ctx)
	if len(filter.Exprs) > 0 {
		query.Where(filter)
	}

	result := query.Count(&count)
	return count, result.Error
}

func (sr *StaticFileMetaDataRepo) UpdateByFileName(ctx context.Context, fileName string, data map[string]interface{}) (entity.StaticFileMetaData, error) {
	updatedRecord := entity.StaticFileMetaData{}
	result := sr.db.WithContext(ctx).Model(&updatedRecord).Clauses(clause.Returning{}).Where("FileName = ?", fileName).Updates(data)
	return updatedRecord, result.Error
}

func (sr *StaticFileMetaDataRepo) DeleteByFileName(ctx context.Context, fileName string) (entity.StaticFileMetaData, error) {
	deletedRecord := entity.StaticFileMetaData{}
	result := sr.db.WithContext(ctx).Where("FileName = ?", fileName).Delete(&deletedRecord)
	return deletedRecord, result.Error
}
