package database

import (
	"context"

	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/entity"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewDbRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (repository *Repository) Save(ctx context.Context, exchange *entity.Exchange) error {
	return repository.DB.WithContext(ctx).Create(exchange).Error
}
