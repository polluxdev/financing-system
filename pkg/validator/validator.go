package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/polluxdev/financing-system/pkg/logger"
)

type Validator interface {
	CustomFunctions() map[string]validator.Func
	RegisterCustomValidator()
	Validate(req interface{}) error
	ParseErrors(err interface{}) []string
}

type ValidatorImpl struct {
	logger   *logger.Logger
	validate *validator.Validate
}

func New(logger *logger.Logger) Validator {
	return &ValidatorImpl{
		logger:   logger,
		validate: validator.New(),
	}
}

func (v *ValidatorImpl) CustomFunctions() map[string]validator.Func {
	return map[string]validator.Func{
		"isValidEmail":       isValidEmail,
		"isValidPhoneNumber": isValidPhoneNumber,
	}
}

func (v *ValidatorImpl) RegisterCustomValidator() {
	for name, fn := range v.CustomFunctions() {
		err := v.validate.RegisterValidation(name, fn)
		if err != nil {
			v.logger.Fatal(err)
		}
	}
}

func (v *ValidatorImpl) Validate(req interface{}) error {
	return v.validate.Struct(req)
}
