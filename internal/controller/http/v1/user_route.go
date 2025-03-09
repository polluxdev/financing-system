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

type userRoutes struct {
	logger    logger.Interface
	config    *config.Config
	validator validator.Validator
	service   interfaces.UserService
}

func newUserRoutes(
	router *gin.RouterGroup,
	logger logger.Interface,
	config *config.Config,
	validator validator.Validator,
	service interfaces.UserService,
) {
	route := &userRoutes{logger, config, validator, service}

	handler := router.Group("/users")
	{
		handler.POST("/", route.createUser)
	}
}

// @Summary     Create user
// @Description Create a new user
// @ID          create-user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body web.CreateUserRequest true "Construct new user"
// @Success     201 {object} app.JSONResponse
// @Failure     400 {object} app.JSONResponse
// @Failure     500 {object} app.JSONResponse
// @Router      /users [post]
func (r *userRoutes) createUser(ctx *gin.Context) {
	requestId := ctx.GetString("requestId")

	var request web.CreateUserRequest

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
