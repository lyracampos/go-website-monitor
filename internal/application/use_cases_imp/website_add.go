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

	//todo: validar entidade
	webistedb, err := w.data.Add(*website)
	if err != nil {
		return nil, errors.New("an exception was occurred")
	}

	response := usecases.WebsiteAddResponse{Id: webistedb.Id, Name: webistedb.Name, Url: webistedb.Url, Status: webistedb.Status}
	return &response, nil
}
