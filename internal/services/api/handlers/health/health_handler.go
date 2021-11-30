

package health

import (
	"log"
	"net/http"
	"website-monitor/internal/domain/infrastructure/data"
)

type healthHandler struct {
	log  *log.Logger
	data data.WebsiteData
}

func NewHealthHandler(l *log.Logger, d data.WebsiteData) *healthHandler {
	return &healthHandler{
		log:  l,
		data: d,
	}
}

func (h *healthHandler) Check(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("api is working :D"))
}
