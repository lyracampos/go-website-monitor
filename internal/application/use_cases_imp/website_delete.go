package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteDeleteUseCase struct {
	data data.WebsiteData
}

func NewWebsiteDeleteUseCase(d data.WebsiteData) usecases.WebsiteDeleteUseCase {
	return &websiteDeleteUseCase{
		data: d,
	}
}

func (w *websiteDeleteUseCase) Execute(command usecases.WebsiteDeleteCommand) (*usecases.WebsiteDeleteResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	_, err = w.data.Delete(command.Id)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteDeleteResponse{Success: true}
	return &response, nil
}
