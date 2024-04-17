package handlers

type IService interface {
	IWalletService
}

type Handlers struct {
	Services IService
}

func New(services IService) *Handlers {
	return &Handlers{
		Services: services,
	}
}
