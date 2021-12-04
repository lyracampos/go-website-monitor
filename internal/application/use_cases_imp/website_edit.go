package usecasesimp

import (
	"errors"
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteEditUseCase struct {
	data data.WebsiteData
}

func NewWebsiteEditUseCase(d data.WebsiteData) usecases.WebsiteEditUseCase {
	return &websiteEditUseCase{
		data: d,
	}
}

func (w *websiteEditUseCase) Execute(command usecases.WebsiteEditCommand) (*usecases.WebsiteEditResponse, error) {
	websiteDb, err := w.data.Get(command.Id)
	if err != nil || websiteDb == nil {
		return nil, errors.New("could not found website with id")
	}

	websiteDb.Edit(command.Name, command.Url)
	err = websiteDb.Validate()
	if err != nil {
		return nil, err
	}

	_, err = w.data.Update(*websiteDb)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteEditResponse{Id: websiteDb.Id, Name: websiteDb.Name, Url: websiteDb.Url, Status: websiteDb.Status}
	return &response, nil
}
