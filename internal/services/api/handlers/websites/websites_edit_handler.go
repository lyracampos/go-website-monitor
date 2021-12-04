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

// swagger:route PUT /websites/{id} websites WebsiteEdit
// Edit one website in application
//
// responses:
//  200: websiteEditResponse
//  404: notFoundResponse
//  501: internalServerErrorResponse
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

	err = command.Validate()
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

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("an exception was occurred."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}
