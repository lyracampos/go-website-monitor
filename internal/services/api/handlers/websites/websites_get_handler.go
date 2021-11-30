package websites

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"

	"github.com/gorilla/mux"
)

type websiteGetHandler struct {
	log  *log.Logger
	data data.WebsiteData
}

// swagger:route GET /websites/1 websites getWebsite
// Return a list of products from the database
// responses:
//	200: websiteResponse
func NewWebsiteGetHandler(l *log.Logger, d data.WebsiteData) *websiteGetHandler {
	return &websiteGetHandler{
		log:  l,
		data: d,
	}
}

func (h *websiteGetHandler) Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("unable to convert id"))
		return
	}
	h.log.Println("getting website ")

	website, err := h.data.Get(id)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("could not found website with id"))
		return
	}

	json.NewEncoder(rw).Encode(website)
}

// Data structure representing a  websites
// swagger:response websiteResponse
type websiteResponseWrapper struct {
	// Newly created product
	// in: body
	Body entities.WebSite
}