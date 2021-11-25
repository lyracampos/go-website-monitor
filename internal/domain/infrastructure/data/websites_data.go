package data

import (
	"website-monitor/internal/domain/entities"
)

type WebsiteData interface {
	List() ([]*entities.WebSite, error)
	Get(id int) (*entities.WebSite, error)
	Add(website entities.WebSite) (*entities.WebSite, error)
	Update(website entities.WebSite) (*entities.WebSite, error)
	Delete(id int) (bool, error)
}
