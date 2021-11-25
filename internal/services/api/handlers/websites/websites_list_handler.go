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

func (h *websiteListHandler) List(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("listing - websites")

	websites, err := h.data.List()
	if err != nil {
		rw.Write([]byte("no websites founded"))
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(websites)
}
