package usecases

type WebsiteDisableUseCase interface {
	Execute(command WebsiteDisableCommand) (*WebsiteDisableResponse, error)
}

type WebsiteDisableCommand struct {
	Id int
}
type WebsiteDisableResponse struct {
	Success bool
}
