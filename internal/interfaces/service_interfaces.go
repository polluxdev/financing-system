package interfaces

import (
	"context"
	"net/http"

	"github.com/polluxdev/financing-system/internal/entity/web"
)

type (
	TenorService interface {
		Create(ctx context.Context, request *http.Request, data web.CreateTenorRequest) error
	}

	UserService interface {
		Create(ctx context.Context, request *http.Request, data web.CreateUserRequest) error
	}
)
