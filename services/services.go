package services

type IRepository interface {
	IWalletRepository
	ICategoryRepository
}

type Services struct {
	Repository IRepository
}

func New(repository IRepository) *Services {
	return &Services{
		Repository: repository,
	}
}
