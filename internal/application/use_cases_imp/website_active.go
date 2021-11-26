package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteActiveUseCase struct {
	data data.WebsiteData
}

func NewWebsiteActiveUseCase(d data.WebsiteData) usecases.WebsiteActiveUseCase {
	return &websiteActiveUseCase{
		data: d,
	}
}

func (w *websiteActiveUseCase) Execute(command usecases.WebsiteActiveCommand) (*usecases.WebsiteActiveResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	websiteDb.Active()
	//todo: validate

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteActiveResponse{Success: true}
	return &response, nil
}
