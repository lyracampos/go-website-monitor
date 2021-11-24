package usecases

type WebsiteDeactiveUseCase interface {
	Execute(command WebsiteDeactiveCommand) (*WebsiteDeactiveResponse, error)
}

type WebsiteDeactiveCommand struct {
	Id int
}
type WebsiteDeactiveResponse struct{}
