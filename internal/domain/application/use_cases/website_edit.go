package usecases

type WebsiteEditUseCase interface {
	Execute(command WebsiteEditCommand) (*WebsiteEditResponse, error)
}

type WebsiteEditCommand struct {
	Id   int
	Name string
	Url  string
}
type WebsiteEditResponse struct {
	Id     int
	Name   string
	Url    string
	Status int
}
