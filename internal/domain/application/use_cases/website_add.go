package usecases

import (
	"github.com/go-playground/validator"
	"website-monitor/internal/domain/types"
	"website-monitor/internal/domain/validations"
)

type WebsiteAddUseCase interface {
	Execute(command WebsiteAddCommand) (*WebsiteAddResponse, error)
}

type WebsiteAddCommand struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required,url"`
}

type WebsiteAddResponse struct {
	Id     int
	Name   string
	Url    string
	Status types.WebsiteStatus
}

func (w *WebsiteAddCommand) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("url", validations.UrlValidate)
	return validate.Struct(w)
}
