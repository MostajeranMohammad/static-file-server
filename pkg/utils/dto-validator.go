package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate
var initialized bool

func ValidateDto(d interface{}) error {
	if !initialized {
		validate = validator.New()
		initialized = true
	}
	return validate.Struct(d)
}
