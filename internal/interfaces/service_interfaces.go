package interfaces

import (
	"context"
	"net/http"

	"github.com/polluxdev/financing-system/internal/entity/web"
)

type (
	FinanceService interface {
		CalculateInstallment(ctx context.Context, request *http.Request, data web.CalculateInstallmentRequest) ([]web.CalculateInstallmentDTO, error)
		SubmitFinancing(ctx context.Context, request *http.Request, data web.SubmitFinancingRequest) error
	}

	TenorService interface {
		Create(ctx context.Context, request *http.Request, data web.CreateTenorRequest) error
	}

	UserService interface {
		Create(ctx context.Context, request *http.Request, data web.CreateUserRequest) error
	}
)
