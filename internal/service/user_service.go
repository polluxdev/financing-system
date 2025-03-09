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

type UserService struct {
	logger     logger.Interface
	db         *postgres.Postgres
	repository interfaces.UserRepository
}

func NewUserService(
	logger logger.Interface,
	db *postgres.Postgres,
	repository interfaces.UserRepository,
) interfaces.UserService {
	return &UserService{
		logger:     logger,
		db:         db,
		repository: repository,
	}
}

func (s *UserService) Create(ctx context.Context, request *http.Request, data web.CreateUserRequest) error {
	newUser := domain.User{
		ID:          uuid.NewString(),
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
	}

	return s.repository.Create(ctx, s.db.DB, newUser)
}
