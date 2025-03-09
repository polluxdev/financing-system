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

type tenorRoutes struct {
	logger    logger.Interface
	config    *config.Config
	validator validator.Validator
	service   interfaces.TenorService
}

func newTenorRoutes(
	router *gin.RouterGroup,
	logger logger.Interface,
	config *config.Config,
	validator validator.Validator,
	service interfaces.TenorService,
) {
	route := &tenorRoutes{logger, config, validator, service}

	handler := router.Group("/tenors")
	{
		handler.POST("/", route.createTenor)
	}
}

// @Summary     Create tenor
// @Description Create a new tenor
// @ID          create-tenor
// @Tags  	    tenor
// @Accept      json
// @Produce     json
// @Param       request body web.CreateTenorRequest true "Construct new tenor"
// @Success     201 {object} app.JSONResponse
// @Failure     400 {object} app.JSONResponse
// @Failure     500 {object} app.JSONResponse
// @Router      /tenors [post]
func (r *tenorRoutes) createTenor(ctx *gin.Context) {
	requestId := ctx.GetString("requestId")

	var request web.CreateTenorRequest

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

	err = r.service.Create(ctx, ctx.Request, request)
	if err != nil {
		statusCode, response := app.ConstructJSONError(requestId, err)
		app.ToJSON(ctx, statusCode, response)
		return
	}

	response := app.ConstructJSONResponse(requestId, global.SUCCESS, global.SUCCESS_CREATE_MESSAGE, nil, nil)
	app.ToJSON(ctx, http.StatusCreated, response)
}
