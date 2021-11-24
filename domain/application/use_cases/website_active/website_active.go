package websiteactive

type WebsiteActiveHandler interface {
	Execute(command WebsiteActiveCommand) (*WebsiteActiveResponse, error)
}

type WebsiteActiveCommand struct{}
type WebsiteActiveResponse struct{}
