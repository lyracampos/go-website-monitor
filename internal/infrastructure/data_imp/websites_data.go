package dataimp

import (
	"context"
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"
)

type webSiteData struct {
}

func NewWebsiteData() data.WebsiteData {
	return &webSiteData{}
}

func (w *webSiteData) List(ctx context.Context) (*entities.WebSites, error) {
	return nil, nil
}
func (w *webSiteData) Get(ctx context.Context, id int) (*entities.WebSite, error) {
	return nil, nil
}
func (w *webSiteData) Add(ctx context.Context, website entities.WebSite) (*entities.WebSite, error) {
	return nil, nil
}
func (w *webSiteData) Edit(ctx context.Context, website entities.WebSite) (*entities.WebSite, error) {
	return nil, nil
}
func (w *webSiteData) Delete(ctx context.Context, id int) (bool, error) {
	return false, nil
}
