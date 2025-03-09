package repository

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"gorm.io/gorm"
)

type TenorRepository struct{}

func NewTenorRepo() interfaces.TenorRepository {
	return &TenorRepository{}
}

func (s *TenorRepository) CreateBulk(ctx context.Context, db *gorm.DB, data []domain.Tenor) error {
	return db.Create(&data).Error
}
