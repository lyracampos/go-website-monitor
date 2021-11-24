package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"

	"github.com/gorilla/mux"
)

type websiteHandler struct {
	log        *log.Logger
	fooUseCase usecases.WebsiteFooUseCase
}

func NewWebisteHandler(l *log.Logger, d *data.WebsiteData, f usecases.WebsiteFooUseCase) *websiteHandler {
	return &websiteHandler{
		log:        l,
		fooUseCase: f,
	}
}

func (h *websiteHandler) Active(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	command := usecases.WebsiteFooCommand{Id: id}
	response, err := h.fooUseCase.Execute(command)
	// command := usecases.WebsiteFooCommand{Id: id}
	// response, err := usecases.WebsiteFooUseCase.Execute(usecases.WebsiteFooUseCase,command)
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
