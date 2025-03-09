package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/polluxdev/financing-system/global"
	"github.com/polluxdev/financing-system/pkg/app"
)

func RecoverError(ctx *gin.Context, err interface{}) {
	if err, ok := err.(error); ok {
		statusCode, response := app.ConstructJSONError(ctx.GetString("requestId"), err)
		app.ToJSON(ctx, statusCode, response)
	} else {
		err = app.ConstructInternalServerError(global.INTERNAL_SERVER_ERROR, global.INTERNAL_SERVER_ERROR_MESSAGE)
		statusCode, response := app.ConstructJSONError(ctx.GetString("requestId"), err)
		app.ToJSON(ctx, statusCode, response)
	}
}
