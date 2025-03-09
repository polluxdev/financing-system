package repository

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"gorm.io/gorm"
)

type UserFacilityDetailRepository struct{}

func NewUserFacilityDetailRepository() interfaces.UserFacilityDetailRepository {
	return &UserFacilityDetailRepository{}
}

func (s *UserFacilityDetailRepository) CreateBulk(ctx context.Context, db *gorm.DB, data []domain.UserFacilityDetail) error {
	return db.Create(&data).Error
}
