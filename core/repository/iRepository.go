package corerepository

import (
	"context"
	"gorm.io/gorm"
)

type IRepository[T any] interface {
	Select(tx *gorm.DB, ctx context.Context, param QueryInfo) ([]T, int32, int32, int32, error)
	Find(tx *gorm.DB, ctx context.Context, id int64) (T, error)
	Create(tx *gorm.DB, ctx context.Context, entity T) (T, error)
	Update(tx *gorm.DB, ctx context.Context, entity T) (T, error)
	Delete(tx *gorm.DB, ctx context.Context, id int64) error
}
