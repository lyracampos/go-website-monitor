package repositories

import "website-monitor/domain/entities"

type WebSiteRepository interface {
	List() []entities.WebSite
	Get(id int) entities.WebSite
	Add(website entities.WebSite)
	Edit(website entities.WebSite)
	Delete(id int)
}
