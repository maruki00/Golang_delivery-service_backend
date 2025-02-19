package app

import "delivery/internal/auth/infra/repositories"

type App struct {
	Repo repositories.AuthRepository
}

func InitApp() (*App, error) {

	return nil, nil
}
