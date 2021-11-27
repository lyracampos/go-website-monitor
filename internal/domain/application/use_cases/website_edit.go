package usecases

import (
	"github.com/go-playground/validator"
	"website-monitor/internal/domain/validations"
)

type WebsiteEditUseCase interface {
	Execute(command WebsiteEditCommand) (*WebsiteEditResponse, error)
}

type WebsiteEditCommand struct {
	Id   int
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required,url"`
}

type WebsiteEditResponse struct {
	Id     int
	Name   string
	Url    string
	Status int
}

func (w *WebsiteEditCommand) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("url", validations.ValidateURL)
	return validate.Struct(w)
}
