package websites

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	useCases "website-monitor/internal/domain/application/use_cases"
)

type websiteActivateHandler struct {
	log     *log.Logger
	useCase useCases.WebsiteActivateUseCase
}

func NewWebsiteActivateHandler(l *log.Logger, u useCases.WebsiteActivateUseCase) *websiteActivateHandler {
	return &websiteActivateHandler{
		log:     l,
		useCase: u,
	}
}

func (h *websiteActivateHandler) Active(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	command := useCases.WebsiteActivateCommand{Id: id}
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
