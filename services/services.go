package services

type IRepository interface {
	IWalletRepository
}

type Services struct {
	Repository IRepository
}

func New(repository IRepository) *Services {
	return &Services{
		Repository: repository,
	}
}
