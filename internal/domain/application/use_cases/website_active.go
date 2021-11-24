package usecases

type WebsiteActiveUseCase interface {
	Execute(command WebsiteActiveCommand) (*WebsiteActiveResponse, error)
}

type WebsiteActiveCommand struct {
	Id int
}
type WebsiteActiveResponse struct{}
