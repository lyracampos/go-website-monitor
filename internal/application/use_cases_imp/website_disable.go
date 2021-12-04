package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteDisableUseCase struct {
	data data.WebsiteData
}

func NewWebsiteDisableUseCase(d data.WebsiteData) usecases.WebsiteDisableUseCase {
	return &websiteDisableUseCase{
		data: d,
	}
}

func (w *websiteDisableUseCase) Execute(command usecases.WebsiteDisableCommand) (*usecases.WebsiteDisableResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	websiteDb.Disable()
	err = websiteDb.Validate()
	if err != nil {
		return nil, err
	}

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteDisableResponse{Success: true}
	return &response, nil
}
