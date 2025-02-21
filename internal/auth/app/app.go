package app

import (
	"context"
	"delivery/internal/auth/app/services"
	"delivery/internal/auth/domain/contracts"
	"delivery/internal/auth/domain/ports"
	"delivery/internal/auth/infra/repositories"
	"delivery/internal/auth/userGateWay/adapters/presenters"

	"github.com/go-playground/validator"
	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	Repo      *contracts.IAuthRepository
	Svc       *services.AuthService
	Validate  *validator.Validate
	InputPort ports.AuthInputPort
}

func InitApp() (*App, error) {

	repo := repositories.NewAuthRepository(nil)
	presenter := &presenters.JsonAuthPresenter{}
	return &App{
		Repo: repo,
	}, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

}
