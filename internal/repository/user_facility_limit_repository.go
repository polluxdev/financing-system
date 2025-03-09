package repository

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"gorm.io/gorm"
)

type UserFacilityLimitRepository struct{}

func NewUserFacilityLimitRepository() interfaces.UserFacilityLimitRepository {
	return &UserFacilityLimitRepository{}
}

func (u *UserFacilityLimitRepository) FindByColumn(ctx context.Context, db *gorm.DB, conditions string, args []interface{}) (*domain.UserFacilityLimit, error) {
	var result domain.UserFacilityLimit
	return &result, db.Where(conditions, args...).Find(&result).Error
}
