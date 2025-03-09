package global

const (
	DATE_FORMAT               = "2006-01-02"
	TIME_FORMAT               = "15:04:05"
	DATE_TIME_FORMAT          = "2006-01-02 15:04:05"
	DATE_TIME_MST_FORMAT      = "2006-01-02 15:04:05 -0700 MST"
	DATE_TIME_ISO_FORMAT      = "2006-01-02T15:04:05Z"
	DATE_TIME_ISO_8601_FORMAT = "2006-01-02 15:04:05.000"
)

const (
	SUCCESS = "SUCCESS"

	BAD_REQUEST_ERROR     = "BAD_REQUEST_ERROR"
	CONFLICT_ERROR        = "CONFLICT_ERROR"
	NOT_FOUND_ERROR       = "NOT_FOUND_ERROR"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	INVALID_DATA          = "INVALID_DATA"

	BAD_REQUEST_MESSAGE           = "bad request"
	INVALID_REQUEST_MESSAGE       = "invalid request"
	INTERNAL_SERVER_ERROR_MESSAGE = "internal server error"
	NOT_FOUND_MESSAGE             = "data not found"
	SUCCESS_CREATE_MESSAGE        = "create data successfully"
	SUCCESS_READ_MESSAGE          = "read data successfully"
	SUCCESS_UPDATE_MESSAGE        = "update data successfully"
	SUCCESS_DELETE_MESSAGE        = "delete data successfully"
)
