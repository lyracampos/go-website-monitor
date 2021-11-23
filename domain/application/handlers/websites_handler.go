package handlers

import (
	"website-monitor/domain/application/models/requests"
	"website-monitor/domain/application/models/responses"
)

type WebsitesHandler interface {
	List() (*responses.WebsiteListResponse, error)
	Add(request requests.WebsiteAddRequest) (*responses.WebsiteAddResponse, error)
	Edit(request requests.WebsiteEditRequest) (*responses.WebsiteEditResponse, error)
	Delete(request requests.WebsiteDeleteRequest) (*responses.WebsiteDeleteResponse, error)
	Active(request requests.WebsiteActiveRequest) (*responses.WebsiteActiveResponse, error)
	Deactive(request requests.WebsiteDeactiveRequest) (*responses.WebsiteDeactiveResponse, error)
	ToMonitor() error
}
