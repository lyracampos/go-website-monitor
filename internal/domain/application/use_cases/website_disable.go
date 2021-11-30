package usecases

type WebsiteDeactivateUseCase interface {
	Execute(command WebsiteDeactivateCommand) (*WebsiteDeactivateResponse, error)
}

type WebsiteDeactivateCommand struct {
	Id int
}
type WebsiteDeactivateResponse struct {
	Success bool
}
