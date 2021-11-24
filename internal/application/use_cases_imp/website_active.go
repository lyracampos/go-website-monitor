package usecasesimp

import (
	usecases "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteActiveUseCase struct {
	websiteData data.WebsiteData
}

func NewWebsiteActiveUseCase(d data.WebsiteData) usecases.WebsiteActiveUseCase {
	return &websiteActiveUseCase{
		websiteData: d,
	}
}

func (w *websiteActiveUseCase) Execute(command usecases.WebsiteActiveCommand) (*usecases.WebsiteActiveResponse, error) {
	var response = usecases.WebsiteActiveResponse{}
	return &response, nil
}
