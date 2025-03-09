package service

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/polluxdev/financing-system/global"
	"github.com/polluxdev/financing-system/helper"
	"github.com/polluxdev/financing-system/internal/entity/domain"
	"github.com/polluxdev/financing-system/internal/entity/web"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"github.com/polluxdev/financing-system/pkg/app"
	"github.com/polluxdev/financing-system/pkg/logger"
	"github.com/polluxdev/financing-system/pkg/postgres"
)

type FinanceService struct {
	logger                       logger.Interface
	db                           *postgres.Postgres
	tenorRepository              interfaces.TenorRepository
	userFacilityDetailRepository interfaces.UserFacilityDetailRepository
	userFacilityLimitRepository  interfaces.UserFacilityLimitRepository
	userFacilityRepository       interfaces.UserFacilityRepository
}

func NewFinanceService(
	logger logger.Interface,
	db *postgres.Postgres,
	tenorRepository interfaces.TenorRepository,
	userFacilityDetailRepository interfaces.UserFacilityDetailRepository,
	userFacilityLimitRepository interfaces.UserFacilityLimitRepository,
	userFacilityRepository interfaces.UserFacilityRepository,
) interfaces.FinanceService {
	return &FinanceService{
		logger:                       logger,
		db:                           db,
		tenorRepository:              tenorRepository,
		userFacilityDetailRepository: userFacilityDetailRepository,
		userFacilityLimitRepository:  userFacilityLimitRepository,
		userFacilityRepository:       userFacilityRepository,
	}
}

func (s *FinanceService) CalculateInstallment(ctx context.Context, request *http.Request, data web.CalculateInstallmentRequest) ([]web.CalculateInstallmentDTO, error) {
	// Fetch tenor data
	conditions, args := helper.ConstructConditionalClause(nil)
	tenors, err := s.tenorRepository.FindAll(ctx, s.db.DB, nil, conditions, args, 0, 0)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	// Check if data not empty
	if len(tenors) == 0 {
		err = app.ConstructNotFoundError(global.NOT_FOUND_ERROR, global.NOT_FOUND_MESSAGE)
		s.logger.Error(err)
		return nil, err
	}

	// Construct dto result
	result := make([]web.CalculateInstallmentDTO, 0)
	for _, item := range tenors {
		// Calculate with margin flat 20% per year
		totalMargin := (data.Amount * helper.CalculateFlatMargin()) / float64(item.Value)
		totalPayment := data.Amount + totalMargin
		monthlyInstallment := totalPayment / float64(item.Value)

		resultItem := web.ToCalculateInstallmentDTO(
			item.Value,
			data.Amount,
			totalMargin,
			totalPayment,
			monthlyInstallment,
		)

		result = append(result, resultItem)
	}

	return result, nil
}

func (s *FinanceService) SubmitFinancing(ctx context.Context, request *http.Request, data web.SubmitFinancingRequest) error {
	// Fetch tenor data
	conditions, args := helper.ConstructConditionalClause(nil)
	tenors, err := s.tenorRepository.FindAll(ctx, s.db.DB, nil, conditions, args, 0, 0)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	// Check if data not empty
	if len(tenors) == 0 {
		err = app.ConstructNotFoundError(global.NOT_FOUND_ERROR, global.NOT_FOUND_MESSAGE)
		s.logger.Error(err)
		return err
	}

	// Extract tenor values
	tenorValues := make([]uint8, 0)
	for _, item := range tenors {
		tenorValues = append(tenorValues, item.Value)
	}

	// Check if tenor data is valid
	if !helper.ContainUint8(tenorValues, data.Tenor) {
		err = app.ConstructBadRequestError(global.BAD_REQUEST_ERROR, "tenor value is invalid")
		s.logger.Error(err)
		return err
	}

	// Fetch user facility limit data
	userID := "6173a3eb-3fd1-4289-9d72-4b81ab9198bc"
	builder := []helper.ConditionalBuilder{
		{
			Column:   "user_id",
			Value:    userID,
			Logical:  "=",
			Operator: "AND",
		},
	}

	conditions, args = helper.ConstructConditionalClause(builder)

	userFacilityLimit, err := s.userFacilityLimitRepository.FindByColumn(ctx, s.db.DB, conditions, args)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	// Check if user facility limit is found
	if userFacilityLimit.ID == "" {
		err = app.ConstructNotFoundError(global.NOT_FOUND_ERROR, global.NOT_FOUND_MESSAGE)
		s.logger.Error(err)
		return err
	}

	// Check if user facility limit amount is sufficient
	if userFacilityLimit.LimitAmount < data.Amount {
		err = app.ConstructBadRequestError(global.BAD_REQUEST_ERROR, "insufficient limit amount")
		s.logger.Error(err)
		return err
	}

	// Calculate with margin flat 20% per year
	totalMargin := (data.Amount * helper.CalculateFlatMargin()) / float64(data.Tenor)
	totalPayment := data.Amount + totalMargin
	monthlyInstallment := totalPayment / float64(data.Tenor)

	// Begin transaction
	tx := s.db.DB.Begin()
	defer helper.CommitAndRollback(tx, &err)()

	// Create new user facility
	startDate, err := helper.ParseStringToTime(global.DATE_FORMAT, data.StartDate, nil)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	newUserFacility := domain.UserFacility{
		ID:                  uuid.NewString(),
		UserID:              userID,
		UserFacilityLimitID: userFacilityLimit.ID,
		Amount:              data.Amount,
		Tenor:               data.Tenor,
		StartDate:           *startDate,
		MonthlyInstallment:  monthlyInstallment,
		TotalMargin:         totalMargin,
		TotalPayment:        totalPayment,
	}

	err = s.userFacilityRepository.Create(ctx, tx, newUserFacility)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	// Create user facility details
	newUserFacilityDetails := make([]domain.UserFacilityDetail, 0)
	for i := range data.Tenor {
		dueDate := startDate.AddDate(0, int(i+1), 0)

		newUserFacilityDetail := domain.UserFacilityDetail{
			ID:                uuid.NewString(),
			UserFacilityID:    newUserFacility.ID,
			DueDate:           dueDate,
			InstallmentAmount: monthlyInstallment,
		}

		newUserFacilityDetails = append(newUserFacilityDetails, newUserFacilityDetail)
	}

	err = s.userFacilityDetailRepository.CreateBulk(ctx, tx, newUserFacilityDetails)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}
