package usecases

type WebsiteAddUseCase interface {
	Execute(command WebsiteAddCommand) (*WebsiteAddResponse, error)
}

type WebsiteAddCommand struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type WebsiteAddResponse struct {
	Id     int
	Name   string
	Url    string
	Status int
}
