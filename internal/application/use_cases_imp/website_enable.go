package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteActivateUseCase struct {
	data data.WebsiteData
}

func NewWebsiteActivateUseCase(d data.WebsiteData) usecases.WebsiteEnableUseCase {
	return &websiteActivateUseCase{
		data: d,
	}
}

func (w *websiteActivateUseCase) Execute(command usecases.WebsiteEnableCommand) (*usecases.WebsiteEnableResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	websiteDb.Enable()
	//todo: validate

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteEnableResponse{Success: true}
	return &response, nil
}
