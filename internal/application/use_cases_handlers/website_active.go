package usecaseshandlers

import (
	websiteactive "website-monitor/internal/domain/application/use_cases/website_active"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteActiveUseCase struct {
	websiteData data.WebsiteData
}

func NewWebsiteActiveUseCase(d data.WebsiteData) websiteactive.WebsiteActiveUseCase {
	return &websiteActiveUseCase{
		websiteData: d,
	}
}

func (w *websiteActiveUseCase) Execute(command websiteactive.WebsiteActiveCommand) (*websiteactive.WebsiteActiveResponse, error) {
	var response = websiteactive.WebsiteActiveResponse{}
	return &response, nil
}
