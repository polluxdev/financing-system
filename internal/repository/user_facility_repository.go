package repository

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"gorm.io/gorm"
)

type UserFacilityRepository struct{}

func NewUserFacilityRepository() interfaces.UserFacilityRepository {
	return &UserFacilityRepository{}
}

func (u *UserFacilityRepository) Create(ctx context.Context, db *gorm.DB, user domain.UserFacility) error {
	return db.Create(&user).Error
}
