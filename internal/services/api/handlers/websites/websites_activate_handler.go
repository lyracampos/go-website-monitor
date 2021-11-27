package websites

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	useCases "website-monitor/internal/domain/application/use_cases"
)

type websiteActiveHandler struct {
	log     *log.Logger
	useCase useCases.WebsiteActivateUseCase
}

func NewWebsiteActiveHandler(l *log.Logger, u useCases.WebsiteActivateUseCase) *websiteActiveHandler {
	return &websiteActiveHandler{
		log:     l,
		useCase: u,
	}
}

func (h *websiteActiveHandler) Active(rw http.ResponseWriter, r *http.Request) {
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
