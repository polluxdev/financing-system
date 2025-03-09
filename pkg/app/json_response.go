package app

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/financing-system/global"
	"github.com/polluxdev/financing-system/internal/entity/web"
)

type JSONResponse struct {
	RequestID string      `json:"request_id"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error,omitempty"`
}

type JSONPaginationResponse struct {
	JSONResponse
	*web.PaginationDTO
}

func ConstructJSONResponse(requestID, code, msg string, data, err interface{}) JSONResponse {
	return JSONResponse{
		RequestID: requestID,
		Code:      code,
		Message:   msg,
		Data:      data,
		Error:     err,
	}
}

func ConstructJSONPaginationResponse(requestID, code, msg string, data interface{}, pagination *web.PaginationDTO, err error) JSONPaginationResponse {
	return JSONPaginationResponse{
		JSONResponse:  ConstructJSONResponse(requestID, code, msg, data, err),
		PaginationDTO: pagination,
	}
}

func ConstructJSONError(requestID string, err error) (int, JSONResponse) {
	var statusCode int
	var response JSONResponse

	if appErr, ok := err.(*AppError); ok {
		statusCode = appErr.Status
		response = ConstructJSONResponse(requestID, appErr.Code, appErr.Message, nil, err)
	} else {
		statusCode = http.StatusInternalServerError
		response = ConstructJSONResponse(requestID, global.INTERNAL_SERVER_ERROR, global.INTERNAL_SERVER_ERROR_MESSAGE, nil, err)
	}

	return statusCode, response
}

func ToJSON(ctx interface{}, code int, data interface{}) {
	switch v := ctx.(type) {
	case *gin.Context:
		v.JSON(code, data)
	case http.ResponseWriter:
		v.WriteHeader(code)
		json.NewEncoder(v).Encode(data)
	default:
	}
}
