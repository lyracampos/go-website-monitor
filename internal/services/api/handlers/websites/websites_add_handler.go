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

//	swagger:route POST /websites websites AddWebsite
//	Add new website in application
//
//	responses:
//	 201: websiteAddResponse
//	 400: badRequestResponse
//   501: internalServerErrorResponse
func (h *websitesAddHandler) Create(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var command useCases.WebsiteAddCommand
	json.Unmarshal(requestBody, &command)

	err := command.Validate()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	response, err := h.useCase.Execute(command)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)
}

// Data structure representing website added
// swagger:response websiteAddResponse
type websiteAddResponseWrapper struct {
	// Newly added website
	// in: body
	Body useCases.WebsiteAddResponse
}

// swagger:parameters AddWebsite
type websiteAddCommandWrapper struct {
	// Payload to add new website in application
	// in: body
	// required: true
	Body useCases.WebsiteAddCommand
}