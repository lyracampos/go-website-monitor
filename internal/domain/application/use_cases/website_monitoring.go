package usecases

type WebsiteMonitoringUseCase interface {
	Execute(command WebsiteMonitoringCommand) (*WebsiteMonitoringResponse, error)
}

type WebsiteMonitoringCommand struct{}
type WebsiteMonitoringResponse struct{}
