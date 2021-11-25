package websites

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	usecases "website-monitor/internal/domain/application/use_cases"

	"github.com/gorilla/mux"
)

type websiteDeactiveHandler struct {
	log     *log.Logger
	usecase usecases.WebsiteDeactiveUseCase
}

func NewWebsiteDeactiveHandler(l *log.Logger, u usecases.WebsiteDeactiveUseCase) *websiteDeactiveHandler {
	return &websiteDeactiveHandler{
		log:     l,
		usecase: u,
	}
}

func (h *websiteDeactiveHandler) Deactive(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	command := usecases.WebsiteDeactiveCommand{Id: id}
	response, err := h.usecase.Execute(command)
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