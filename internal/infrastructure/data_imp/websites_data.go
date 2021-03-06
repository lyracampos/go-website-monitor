package dataimp

import (
	"errors"
	"time"
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"
)

var portalG1 = entities.WebSite{Id: 1, Name: "G1.com", Url: "http://www.g1.com.br", Status: 1, LastChecked: time.Now()}
var uol = entities.WebSite{Id: 2, Name: "UOL", Url: "http://www.uol.com.br", Status: 1, LastChecked: time.Now().Add(-2)}
var youtube = entities.WebSite{Id: 3, Name: "Youtube", Url: "http://www.youtube.com.br", Status: 1, LastChecked: time.Now()}
var websites = []*entities.WebSite{&portalG1, &uol, &youtube}

type webSiteData struct {
}

func NewWebsiteData() data.WebsiteData {
	return &webSiteData{}
}

func (w *webSiteData) List() ([]*entities.WebSite, error) {
	return websites, nil
}
func (w *webSiteData) Get(id int) (*entities.WebSite, error) {
	for _, site := range websites {
		if site.Id == id {
			return site, nil
		}
	}
	return nil, errors.New("could not found website")
}
func (w *webSiteData) Insert(website entities.WebSite) (*entities.WebSite, error) {
	newId := 1
	lastWebsite := websites[len(websites)-1]

	if lastWebsite != nil {
		newId = lastWebsite.Id + 1
	}
	website.Id = newId
	websites = append(websites, &website)
	return &website, nil
}
func (w *webSiteData) Update(website entities.WebSite) (*entities.WebSite, error) {
	websites[website.Id-1].Name = website.Name
	websites[website.Id-1].Url = website.Url
	websites[website.Id-1].LastUpdated = website.LastUpdated
	return websites[website.Id-1], nil
}
func (w *webSiteData) Delete(id int) (bool, error) {
	copy(websites[id:], websites[id+1:])
	websites[len(websites)-1] = nil
	websites = websites[:len(websites)-1]

	return true, nil
}
