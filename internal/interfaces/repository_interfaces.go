package interfaces

import (
	"context"

	"github.com/polluxdev/financing-system/internal/entity/domain"
	"gorm.io/gorm"
)

type (
	TenorRepository interface {
		CreateBulk(ctx context.Context, db *gorm.DB, data []domain.Tenor) error
		FindAll(ctx context.Context, db *gorm.DB, fields []string, conditions string, args []interface{}, offset, limit int) ([]domain.Tenor, error)
	}

	UserFacilityDetailRepository interface {
		CreateBulk(ctx context.Context, db *gorm.DB, data []domain.UserFacilityDetail) error
	}

	UserFacilityLimitRepository interface {
		FindByColumn(ctx context.Context, db *gorm.DB, conditions string, args []interface{}) (*domain.UserFacilityLimit, error)
	}

	UserFacilityRepository interface {
		Create(ctx context.Context, db *gorm.DB, user domain.UserFacility) error
	}

	UserRepository interface {
		Create(ctx context.Context, db *gorm.DB, data domain.User) error
	}
)
