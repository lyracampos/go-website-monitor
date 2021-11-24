package usecasesimp

import (
	usecaseimp "website-monitor/internal/domain/application/use_cases"
	"website-monitor/internal/domain/infrastructure/data"
)

type websiteFooUseCase struct {
	websiteData data.WebsiteData
}

func NewWebsiteFooUseCase(d data.WebsiteData) usecaseimp.WebsiteFooUseCase {
	return &websiteFooUseCase{
		websiteData: d,
	}
}

func (w *websiteFooUseCase) Execute(command usecaseimp.WebsiteFooCommand) (*usecaseimp.WebsiteFooResponse, error) {
	var response = usecaseimp.WebsiteFooResponse{}
	return &response, nil
}
