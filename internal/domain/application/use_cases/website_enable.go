package usecases

type WebsiteEnableUseCase interface {
	Execute(command WebsiteEnableCommand) (*WebsiteEnableResponse, error)
}

type WebsiteEnableCommand struct {
	Id int
}
type WebsiteEnableResponse struct {
	Success bool
}
