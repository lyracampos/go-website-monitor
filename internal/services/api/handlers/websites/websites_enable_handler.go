package websites

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	useCases "website-monitor/internal/domain/application/use_cases"
)

type websiteEnableHandler struct {
	log     *log.Logger
	useCase useCases.WebsiteEnableUseCase
}

func NewWebsiteEnableHandler(l *log.Logger, u useCases.WebsiteEnableUseCase) *websiteEnableHandler {
	return &websiteEnableHandler{
		log:     l,
		useCase: u,
	}
}

// swagger:route PUT /websites/{id}/enable websites WebsiteEnable
// Enable monitoring to one website from the application
//
// responses:
//  200: websiteEnableResponse
//  404: notFoundResponse
//  501: internalServerErrorResponse
func (h *websiteEnableHandler) Enable(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	command := useCases.WebsiteEnableCommand{Id: id}
	response, err := h.useCase.Execute(command)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}
}
