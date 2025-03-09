package app

import (
	"errors"

	"github.com/polluxdev/financing-system/global"
	"gorm.io/gorm"
)

func WrapErrorMap(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ConstructNotFoundError(global.NOT_FOUND_ERROR, "error not found")
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ConstructConflictError(global.CONFLICT_ERROR, "error duplicate")
	case errors.Is(err, gorm.ErrInvalidData):
		return ConstructBadRequestError(global.BAD_REQUEST_ERROR, "invalid data provided")
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return ConstructBadRequestError(global.BAD_REQUEST_ERROR, "primary key required")
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return ConstructConflictError(global.CONFLICT_ERROR, "foreign key constraint violated")
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return ConstructInternalServerError(global.INTERNAL_SERVER_ERROR, "invalid transaction")
	case errors.Is(err, gorm.ErrNotImplemented):
		return ConstructInternalServerError(global.INTERNAL_SERVER_ERROR, "operation not implemented")
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return ConstructBadRequestError(global.BAD_REQUEST_ERROR, "missing WHERE clause")
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return ConstructInternalServerError(global.INTERNAL_SERVER_ERROR, "dry-run mode unsupported")
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return ConstructInternalServerError(global.INTERNAL_SERVER_ERROR, "unsupported driver")
	case errors.Is(err, gorm.ErrPreloadNotAllowed):
		return ConstructBadRequestError(global.BAD_REQUEST_ERROR, "preload not allowed")
	case errors.Is(err, gorm.ErrInvalidField):
		return ConstructBadRequestError(global.BAD_REQUEST_ERROR, "invalid field")
	case errors.Is(err, gorm.ErrInvalidValue):
		return ConstructBadRequestError(global.BAD_REQUEST_ERROR, "invalid field value")
	case errors.Is(err, gorm.ErrInvalidDB):
		return ConstructInternalServerError(global.INTERNAL_SERVER_ERROR, "invalid database connection")
	default:
		return err
	}
}
