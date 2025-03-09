package repository

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"gorm.io/gorm"
)

type TenorRepository struct{}

func NewTenorRepository() interfaces.TenorRepository {
	return &TenorRepository{}
}

func (s *TenorRepository) CreateBulk(ctx context.Context, db *gorm.DB, data []domain.Tenor) error {
	return db.Create(&data).Error
}

func (u *TenorRepository) FindAll(ctx context.Context, db *gorm.DB, fields []string, conditions string, args []interface{}, offset, limit int) ([]domain.Tenor, error) {
	var result []domain.Tenor

	query := db.Where(conditions, args...)
	if len(fields) > 0 {
		query = query.Select(fields)
	}

	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}

	return result, query.Find(&result).Error
}
