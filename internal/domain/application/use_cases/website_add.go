package usecases

type WebsiteAddHandler interface {
	Execute(command WebsiteAddCommand) (*WebsiteAddResponse, error)
}

type WebsiteAddCommand struct{}
type WebsiteAddResponse struct{}
