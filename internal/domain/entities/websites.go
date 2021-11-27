package entities

import (
	"errors"
	"time"
)

//todo: criar const/enum para status
type WebSite struct {
	Id          int
	Name        string //obrigatóio, 3+
	Url         string //obrigatório, válido, 10+
	Status      int    //obrigatório
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

func (w *WebSite) Active() error {
	if w.Status == 1 {
		return errors.New("this website already are activated")
	}
	w.Status = 1
	w.LastUpdated = time.Now()
	return nil
}

func (w *WebSite) Deactive() error {
	if w.Status == 0 {
		return errors.New("this website already are deactivated")
	}
	w.Status = 0
	w.LastUpdated = time.Now()
	return nil
}
