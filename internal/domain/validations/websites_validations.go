package validations

import (
	"github.com/go-playground/validator"
	"net/url"
)

func UrlValidate(fl validator.FieldLevel) bool {
	_, err := url.ParseRequestURI(fl.Field().String())
	if err != nil {
		return false
	}

	return true
}