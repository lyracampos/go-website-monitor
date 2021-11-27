package websites

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	useCases "website-monitor/internal/domain/application/use_cases"
)

type websitesEditHandler struct {
	log     *log.Logger
	useCase useCases.WebsiteEditUseCase
}

func NewWebsitesEditHandler(l *log.Logger, u useCases.WebsiteEditUseCase) *websitesEditHandler {
	return &websitesEditHandler{
		log:     l,
		useCase: u,
	}
}

func (h *websitesEditHandler) Edit(rw http.ResponseWriter, r *http.Request) {
	//todo: revisar esse handler
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unable to convert id"))
		return
	}

	requestBody, _ := ioutil.ReadAll(r.Body)
	var command useCases.WebsiteEditCommand
	json.Unmarshal(requestBody, &command)
	command.Id = id

	response, err := h.useCase.Execute(command)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
	}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}
