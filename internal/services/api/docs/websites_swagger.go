package docs

import (
	useCases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/entities"
)

// Data structure representing website added
// swagger:response websiteAddResponse
type websiteAddResponseWrapper struct {
	// in: body
	Body useCases.WebsiteAddResponse
}

// swagger:parameters AddWebsite
type websiteAddCommandWrapper struct {
	// Payload to add new website in application
	// in: body
	// required: true
	Body useCases.WebsiteAddCommand
}

// Data structure representing a list of website
// swagger:response websiteListResponse
type websitesListResponseWrapper struct {
	// Newly created product
	// in: body
	Body []entities.WebSite
}

// Data structure representing a  websites
// swagger:response websiteGetResponse
type websiteGetResponseWrapper struct {
	// in: body
	Body entities.WebSite
}

// swagger:parameters WebsiteDelete WebsiteEnable EditWebsite WebsiteDisable
type websiteIdParamsWrapper struct {
	// the id of website for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// Data structure representing website edited
// swagger:response websiteEditResponse
type websiteEditResponseWrapper struct {
	// in: body
	Body useCases.WebsiteEditResponse
}

// swagger:parameters EditWebsite
type websiteEditCommandWrapper struct {
	// Payload to edit a website in application
	// in: body
	// required: true
	Body useCases.WebsiteEditCommand
}

// Data structure representing website enabled_
// swagger:response websiteEnableResponse
type websiteEnableResponseWrapper struct {
	// in: body
	Body useCases.WebsiteEnableResponse
}

// Data structure representing website disabled
// swagger:response websiteDisableResponse
type websiteDisableResponseWrapper struct {
	// in: body
	Body useCases.WebsiteDisableResponse
}
