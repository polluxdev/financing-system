package service

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/entity/web"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"github.com/polluxdev/financing-system/pkg/logger"
	"github.com/polluxdev/financing-system/pkg/postgres"
)

type TenorService struct {
	logger     logger.Interface
	db         *postgres.Postgres
	repository interfaces.TenorRepository
}

func NewTenorService(
	logger logger.Interface,
	db *postgres.Postgres,
	repository interfaces.TenorRepository,
) interfaces.TenorService {
	return &TenorService{
		logger:     logger,
		db:         db,
		repository: repository,
	}
}

func (s *TenorService) Create(ctx context.Context, request *http.Request, data web.CreateTenorRequest) error {
	newTenors := make([]domain.Tenor, 0)
	for _, item := range data.Data {
		newTenor := domain.Tenor{
			ID:    uuid.NewString(),
			Value: item,
		}
		newTenors = append(newTenors, newTenor)
	}

	return s.repository.CreateBulk(ctx, s.db.DB, newTenors)
}
