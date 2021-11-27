package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteDeactivateUseCase struct {
	data data.WebsiteData
}

func NewWebsiteDeactivateUseCase(d data.WebsiteData) usecases.WebsiteDeactivateUseCase {
	return &websiteDeactivateUseCase{
		data: d,
	}
}

func (w *websiteDeactivateUseCase) Execute(command usecases.WebsiteDeactivateCommand) (*usecases.WebsiteDeactivateResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	websiteDb.Deactivate()
	//todo: validate

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteDeactivateResponse{Success: true}
	return &response, nil
}
