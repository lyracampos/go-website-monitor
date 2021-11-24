package usecases

type WebsiteMonitoringHandler interface {
	Execute(command WebsiteMonitoringCommand) (*WebsiteMonitoringResponse, error)
}

type WebsiteMonitoringCommand struct{}
type WebsiteMonitoringResponse struct{}
