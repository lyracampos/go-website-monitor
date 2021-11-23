package data

import (
	"context"
	"website-monitor/domain/core/entities"
)

type WebsiteRepository interface {
	List(ctx context.Context) (*entities.WebSites, error)
	Get(ctx context.Context, id int) (*entities.WebSite, error)
	Add(ctx context.Context, website entities.WebSite) (*entities.WebSite, error)
	Edit(ctx context.Context, website entities.WebSite) (*entities.WebSite, error)
	Delete(ctx context.Context, id int) (bool, error)
}
