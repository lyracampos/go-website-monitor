package entities

import (
	"errors"
	"github.com/go-playground/validator"
	"time"
	"website-monitor/internal/domain/types"
	"website-monitor/internal/domain/validations"
)

type WebSite struct {
	Id          int
	Name        string  `validate:"required"`
	Url         string  `validate:"required,url"`
	Status      types.WebsiteStatus
	LastChecked time.Time
	LastUpdated time.Time
}

func NewWebsite(name string, url string) *WebSite {
	return &WebSite{
		Name:        name,
		Url:         url,
		Status:      1,
		LastChecked: time.Now(),
	}
}

func (w *WebSite) Edit(name string, url string) {
	w.Name = name
	w.Url = url
	w.LastUpdated = time.Now()
}

func (w *WebSite) Enable() error {
	if w.Status == types.Enabled {
		return errors.New("website already are activated")
	}
	w.Status = types.Enabled
	w.LastUpdated = time.Now()
	return nil
}

func (w *WebSite) Disable() error {
	if w.Status == types.Disabled {
		return errors.New("this website already are deactivated")
	}
	w.Status = types.Disabled
	w.LastUpdated = time.Now()
	return nil
}

func (w *WebSite) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("url", validations.UrlValidate)
	return validate.Struct(w)
}
