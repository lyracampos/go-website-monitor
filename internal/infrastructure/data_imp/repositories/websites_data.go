package repositories

import (
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"
	"website-monitor/internal/infrastructure/data_imp/data_context"
)

type webSiteData struct {
}

func NewWebsiteData() data.WebsiteData {
	return &webSiteData{}
}

func (w *webSiteData) List() ([]*entities.WebSite, error) {
	data_context.Connect()
	var websites []*entities.WebSite
	data_context.DB.Find(&websites)
	return websites, nil
}

func (w *webSiteData) Get(id int) (*entities.WebSite, error) {
	data_context.Connect()
	var website *entities.WebSite
	data_context.DB.First(&website, id)
	return website, nil
}
func (w *webSiteData) Insert(website entities.WebSite) (*entities.WebSite, error) {
	data_context.Connect()
	data_context.DB.Create(&website)
	return &website, nil
}

func (w *webSiteData) Update(website entities.WebSite) (*entities.WebSite, error) {
	data_context.Connect()
	data_context.DB.Save(&website)
	return &website, nil
}

func (w *webSiteData) Delete(id int) (bool, error) {
	data_context.Connect()
	var website *entities.WebSite
	data_context.DB.Delete(&website, id)
	return true, nil
}
