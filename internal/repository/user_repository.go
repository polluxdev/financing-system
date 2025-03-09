package repository

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepo() interfaces.UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Create(ctx context.Context, db *gorm.DB, user domain.User) error {
	return db.Create(&user).Error
}
