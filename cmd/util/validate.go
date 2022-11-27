package validate

import (
	"github.com/KaioMarxDEV/gofinance/cmd/model"
	"github.com/go-playground/validator"
)

// struct used to define vary types of errors by struct validate definition
type errorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// Validation by models struct definition (validate:)
var validate = validator.New()

func Struct(user model.User) []*errorResponse {
	var errors []*errorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
