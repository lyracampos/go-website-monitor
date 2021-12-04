package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/entities"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteAddUseCase struct {
	data data.WebsiteData
}

func NewWebsiteAddUseCase(d data.WebsiteData) usecases.WebsiteAddUseCase {
	return &websiteAddUseCase{
		data: d,
	}
}

func (w *websiteAddUseCase) Execute(command usecases.WebsiteAddCommand) (*usecases.WebsiteAddResponse, error) {
	website := entities.NewWebsite(command.Name, command.Url)

	err := website.Validate()
	if err != nil {
		return nil, err
	}
	websiteCurrent, err := w.data.Insert(*website)
	if err != nil || websiteCurrent == nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteAddResponse{Id: websiteCurrent.Id, Name: websiteCurrent.Name, Url: websiteCurrent.Url, Status: websiteCurrent.Status}
	return &response, nil
}
