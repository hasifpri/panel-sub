package infrastructure

import "github.com/go-playground/validator/v10"

var VALIDATOR *validator.Validate

func InitializeValidator() {
	VALIDATOR = validator.New()
}
