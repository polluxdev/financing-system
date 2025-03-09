package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/financing-system/config"
	"github.com/polluxdev/financing-system/global"
	"github.com/polluxdev/financing-system/internal/entity/web"
	"github.com/polluxdev/financing-system/internal/interfaces"
	"github.com/polluxdev/financing-system/pkg/app"
	"github.com/polluxdev/financing-system/pkg/logger"
	"github.com/polluxdev/financing-system/pkg/validator"
)

type financeRoutes struct {
	logger    logger.Interface
	config    *config.Config
	validator validator.Validator
	service   interfaces.FinanceService
}

func newFinanceRoutes(
	router *gin.RouterGroup,
	logger logger.Interface,
	config *config.Config,
	validator validator.Validator,
	service interfaces.FinanceService,
) {
	route := &financeRoutes{logger, config, validator, service}
	router.POST("/calculate-installments", route.calculateInstallment)
	router.POST("/submit-financing", route.submitFinancing)
}

// @Summary     Calculate installments
// @Description Calculate installments
// @ID          calculate-installments
// @Tags  	    finance
// @Accept      json
// @Produce     json
// @Param       request body web.CalculateInstallmentRequest true "Construct calculate installment request"
// @Success     200 {object} app.JSONResponse
// @Failure     400 {object} app.JSONResponse
// @Failure     500 {object} app.JSONResponse
// @Router      /calculate-installments [post]
func (r *financeRoutes) calculateInstallment(ctx *gin.Context) {
	requestId := ctx.GetString("requestId")

	var request web.CalculateInstallmentRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		r.logger.Error(err.Error())
		response := app.ConstructJSONResponse(requestId, global.BAD_REQUEST_ERROR, global.INVALID_REQUEST_MESSAGE, nil, err)
		app.ToJSON(ctx, http.StatusBadRequest, response)
		return
	}

	err = r.validator.Validate(&request)
	if err != nil {
		r.logger.Error(err.Error())
		errMsgs := r.validator.ParseErrors(err)
		response := app.ConstructJSONResponse(requestId, global.BAD_REQUEST_ERROR, global.BAD_REQUEST_MESSAGE, errMsgs, err)
		app.ToJSON(ctx, http.StatusBadRequest, response)
		return
	}

	result, err := r.service.CalculateInstallment(ctx, ctx.Request, request)
	if err != nil {
		statusCode, response := app.ConstructJSONError(requestId, err)
		app.ToJSON(ctx, statusCode, response)
		return
	}

	response := app.ConstructJSONResponse(requestId, global.SUCCESS, global.SUCCESS_CREATE_MESSAGE, result, nil)
	app.ToJSON(ctx, http.StatusOK, response)
}

// @Summary     Submit Financing
// @Description Submit Financing
// @ID          submit-financing
// @Tags  	    finance
// @Accept      json
// @Produce     json
// @Param       request body web.SubmitFinancingRequest true "Construct submit financing request"
// @Success     200 {object} app.JSONResponse
// @Failure     400 {object} app.JSONResponse
// @Failure     500 {object} app.JSONResponse
// @Router      /submit-financing [post]
func (r *financeRoutes) submitFinancing(ctx *gin.Context) {
	requestId := ctx.GetString("requestId")

	var request web.SubmitFinancingRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		r.logger.Error(err.Error())
		response := app.ConstructJSONResponse(requestId, global.BAD_REQUEST_ERROR, global.INVALID_REQUEST_MESSAGE, nil, err)
		app.ToJSON(ctx, http.StatusBadRequest, response)
		return
	}

	err = r.validator.Validate(&request)
	if err != nil {
		r.logger.Error(err.Error())
		errMsgs := r.validator.ParseErrors(err)
		response := app.ConstructJSONResponse(requestId, global.BAD_REQUEST_ERROR, global.BAD_REQUEST_MESSAGE, errMsgs, err)
		app.ToJSON(ctx, http.StatusBadRequest, response)
		return
	}

	err = r.service.SubmitFinancing(ctx, ctx.Request, request)
	if err != nil {
		statusCode, response := app.ConstructJSONError(requestId, err)
		app.ToJSON(ctx, statusCode, response)
		return
	}

	response := app.ConstructJSONResponse(requestId, global.SUCCESS, global.SUCCESS_CREATE_MESSAGE, nil, nil)
	app.ToJSON(ctx, http.StatusOK, response)
}
