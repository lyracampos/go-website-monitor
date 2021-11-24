package websitedeactive

type WebsiteDeactiveHandler interface {
	Execute(command WebsiteDeactiveCommand) (*WebsiteDeactiveResponse, error)
}

type WebsiteDeactiveCommand struct{}
type WebsiteDeactiveResponse struct{}
