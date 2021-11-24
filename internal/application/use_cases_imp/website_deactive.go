package usecasesimp

import (
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
	var response = usecases.WebsiteDeactiveResponse{}
	return &response, nil
}
