package usecases

type WebsiteActivateUseCase interface {
	Execute(command WebsiteActivateCommand) (*WebsiteActivateResponse, error)
}

type WebsiteActivateCommand struct {
	Id int
}
type WebsiteActivateResponse struct {
	Success bool
}
