package websites

import (
	"encoding/json"
	"log"
	"net/http"
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

// swagger:route GET /websites websites ListWebsites
// Return a list of website from the application
// responses:
//	200: websiteListResponse
//	501: internalServerErrorResponse
func (h *websiteListHandler) List(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("listing - web sites")
	rw.Header().Set("Content-Type", "application/json")

	websites, err := h.data.List()
	if err != nil {
		rw.Write([]byte("no websites founded"))
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(rw).Encode(websites)
}
