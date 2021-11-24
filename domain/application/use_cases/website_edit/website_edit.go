package websiteedit

type WebsiteEditHandler interface {
	Execute(command WebsiteEditCommand) (*WebsiteEditResponse, error)
}

type WebsiteEditCommand struct{}
type WebsiteEditResponse struct{}
