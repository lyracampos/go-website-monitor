package websites

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	useCases "website-monitor/internal/domain/application/use_cases"
)

type websitesDeleteHandler struct {
	log *log.Logger
	useCase useCases.WebsiteDeleteUseCase
}

func NewWebsiteDeleteHandler(l *log.Logger, u useCases.WebsiteDeleteUseCase) *websitesDeleteHandler {
	return &websitesDeleteHandler{
		log:l,
		useCase: u,
	}
}

// swagger:route DELETE /websites/{id} websites WebsiteDelete
// Remove website from the application
//
// responses:
//  204: noContentResponse
//  404: badRequestResponse
//  501: internalServerErrorResponse
func (h *websitesDeleteHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	command := useCases.WebsiteDeleteCommand{Id: id}
	response, err := h.useCase.Execute(command)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred"))
		return
	}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}
}