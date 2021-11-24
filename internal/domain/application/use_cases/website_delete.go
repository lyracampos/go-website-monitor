package usecases

type WebsiteDeleteHandler interface {
	Execute(command WebsiteDeleteCommand) (*WebsiteDeleteResponse, error)
}

type WebsiteDeleteCommand struct{}
type WebsiteDeleteResponse struct{}
