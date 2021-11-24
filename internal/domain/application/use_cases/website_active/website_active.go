package websiteactive

type WebsiteActiveUseCase interface {
	Execute(command WebsiteActiveCommand) (*WebsiteActiveResponse, error)
}

type WebsiteActiveCommand struct{}
type WebsiteActiveResponse struct{}
