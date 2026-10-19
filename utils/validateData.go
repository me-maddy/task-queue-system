package utils

import (
	"strings"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateData(data any) *string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	var messages []string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			messages = append(messages, err.Field()+" is required")
		case "email":
			messages = append(messages, err.Field()+" must be a valid email address")
		case "min":
			messages = append(messages, err.Field()+" must be at least "+err.Param()+" characters long")
		case "max":
			messages = append(messages, err.Field()+" must be at most "+err.Param()+" characters long")
		case "len":
			messages = append(messages, err.Field()+" must be exactly "+err.Param()+" characters long")
		default:
			messages = append(messages, err.Field()+" is invalid")
		}
	}
	errorMessage := strings.Join(messages, " || ")
	return &errorMessage
}
