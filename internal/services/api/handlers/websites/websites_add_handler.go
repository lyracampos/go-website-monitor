package websites

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	useCases "website-monitor/internal/domain/application/use_cases"
)

type websitesAddHandler struct {
	log     *log.Logger
	useCase useCases.WebsiteAddUseCase
}

func NewWebsiteAddHandler(l *log.Logger, u useCases.WebsiteAddUseCase) *websitesAddHandler {
	return &websitesAddHandler{
		log:     l,
		useCase: u,
	}
}

func (h *websitesAddHandler) Add(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var command useCases.WebsiteAddCommand
	json.Unmarshal(requestBody, &command)

	response, err := h.useCase.Execute(command)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)
}
