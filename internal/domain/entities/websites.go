package entities

import (
	"errors"
	"github.com/go-playground/validator"
	"time"
	"website-monitor/internal/domain/validations"
)

//todo: criar const/enum para status
type WebSite struct {
	Id          int
	Name        string  `validate:"required"`
	Url         string  `validate:"required,url"`
	Status      int
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

func (w *WebSite) Activate() error {
	if w.Status == 1 {
		return errors.New("website already are activated")
	}
	w.Status = 1
	w.LastUpdated = time.Now()
	return nil
}

func (w *WebSite) Deactivate() error {
	if w.Status == 0 {
		return errors.New("this website already are deactivated")
	}
	w.Status = 0
	w.LastUpdated = time.Now()
	return nil
}

func (w *WebSite) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("url", validations.UrlValidate)
	return validate.Struct(w)
}
