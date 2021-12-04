package websites

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	usecases "website-monitor/internal/domain/application/use_cases"

	"github.com/gorilla/mux"
)

type websiteDisableHandler struct {
	log     *log.Logger
	useCase usecases.WebsiteDisableUseCase
}

func NewWebsiteDisableHandler(l *log.Logger, u usecases.WebsiteDisableUseCase) *websiteDisableHandler {
	return &websiteDisableHandler{
		log:     l,
		useCase: u,
	}
}

func (h *websiteDisableHandler) Disable(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	command := usecases.WebsiteDisableCommand{Id: id}
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
	rw.WriteHeader(http.StatusOK)
}
