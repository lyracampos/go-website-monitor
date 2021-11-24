package usecases

type WebsiteFooUseCase interface {
	Execute(command WebsiteFooCommand) (*WebsiteFooResponse, error)
}

type WebsiteFooCommand struct {
	Id int
}
type WebsiteFooResponse struct{}
