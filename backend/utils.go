package main

import "github.com/go-playground/validator/v10"

func BuildErrorMessage(err error) map[string]string {
	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		message := err.Tag()
		errors[field] = message
	}
	return errors
}
