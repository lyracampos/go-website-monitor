package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteEnableUseCase struct {
	data data.WebsiteData
}

func NewWebsiteEnableUseCase(d data.WebsiteData) usecases.WebsiteEnableUseCase {
	return &websiteEnableUseCase{
		data: d,
	}
}

func (w *websiteEnableUseCase) Execute(command usecases.WebsiteEnableCommand) (*usecases.WebsiteEnableResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	websiteDb.Enable()

	err = websiteDb.Validate()
	if err != nil {
		return nil, err
	}

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteEnableResponse{Success: true}
	return &response, nil
}
