package dataimp

import (
	"time"
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"
)

var portalG1 = entities.WebSite{Id: 1, Name: "G1.com", Url: "http://www.g1.com.br", Status: 1, LastCheck: time.Now()}
var uol = entities.WebSite{Id: 1, Name: "UOL", Url: "http://www.uol.com.br", Status: 1, LastCheck: time.Now().Add(-2)}
var youtube = &entities.WebSite{Id: 1, Name: "Youtube", Url: "http://www.youtube.com.br", Status: 1, LastCheck: time.Now()}
var websites = []*entities.WebSite{&portalG1, &uol}

type webSiteData struct {
}

func NewWebsiteData() data.WebsiteData {
	return &webSiteData{}
}

func (w *webSiteData) List() ([]*entities.WebSite, error) {
	return websites, nil
}
func (w *webSiteData) Get(id int) (*entities.WebSite, error) {
	return nil, nil
}
func (w *webSiteData) Add(website entities.WebSite) (*entities.WebSite, error) {
	return nil, nil
}
func (w *webSiteData) Update(website entities.WebSite) (*entities.WebSite, error) {
	return nil, nil
}
func (w *webSiteData) Delete(id int) (bool, error) {
	return false, nil
}
