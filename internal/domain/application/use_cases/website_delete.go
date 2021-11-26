package usecases

type WebsiteDeleteUseCase interface {
	Execute(command WebsiteDeleteCommand) (*WebsiteDeleteResponse, error)
}

type WebsiteDeleteCommand struct {
	Id int
}
type WebsiteDeleteResponse struct {
	Success bool
}
