package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteDeactiveUseCase struct {
	data data.WebsiteData
}

func NewWebsiteDeactiveUseCase(d data.WebsiteData) usecases.WebsiteDeactiveUseCase {
	return &websiteDeactiveUseCase{
		data: d,
	}
}

func (w *websiteDeactiveUseCase) Execute(command usecases.WebsiteDeactiveCommand) (*usecases.WebsiteDeactiveResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found this website")
	}

	websiteDb.Deactive()
	//todo: validate

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteDeactiveResponse{Success: true}
	return &response, nil
}
