// Package classification Website Monitor API.
//
// the purpose of this api is to provide a monitoring service for websites
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package websites

import (
	"encoding/json"
	"log"
	"net/http"
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteListHandler struct {
	log  *log.Logger
	data data.WebsiteData
}

func NewWebsiteListHandler(l *log.Logger, d data.WebsiteData) *websiteListHandler {
	return &websiteListHandler{
		log:  l,
		data: d,
	}
}

// swagger:route GET /websites websites listWebsites
// Return a list of products from the database
// responses:
//	200: websitesResponse
func (h *websiteListHandler) List(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("listing - websites")
	rw.Header().Set("Content-Type", "application/json")

	websites, err := h.data.List()
	if err != nil {
		rw.Write([]byte("no websites founded"))
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(rw).Encode(websites)
}

// Data structure representing a list of websites
// swagger:response websitesResponse
type websitesResponseWrapper struct {
	// Newly created product
	// in: body
	Body []entities.WebSite
}