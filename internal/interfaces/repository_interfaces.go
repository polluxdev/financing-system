package interfaces

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"gorm.io/gorm"
)

type (
	TenorRepository interface {
		CreateBulk(ctx context.Context, db *gorm.DB, data []domain.Tenor) error
	}

	UserRepository interface {
		Create(ctx context.Context, db *gorm.DB, data domain.User) error
	}
)
