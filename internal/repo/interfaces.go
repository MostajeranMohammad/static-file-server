package repo

import (
	"context"

	"github.com/MostajeranMohammad/static-file-server/internal/entity"
	"gorm.io/gorm/clause"
)

type (
	StaticFileMetaData interface {
		Create(ctx context.Context, data entity.StaticFileMetaData) (entity.StaticFileMetaData, error)
		GetByFileName(ctx context.Context, fileName string) (entity.StaticFileMetaData, error)
		GetAll(ctx context.Context, filter clause.AndConditions, skip int, limit int) ([]entity.StaticFileMetaData, error)
		GetFileAccessData(ctx context.Context, fileName string) (entity.StaticFileMetaData, error)
		CountFiles(ctx context.Context, filter clause.AndConditions) (int64, error)
		UpdateByFileName(ctx context.Context, fileName string, data map[string]interface{}) (entity.StaticFileMetaData, error)
		DeleteByFileName(ctx context.Context, fileName string) (entity.StaticFileMetaData, error)
	}
)
