package order_app

import (
	"context"
	order_application_services "delivery/Services/Order/Application/Services"
	order_domain_contracts "delivery/Services/Order/Domain/Contracts"
	order_domain_ports "delivery/Services/Order/Domain/Ports"
)

type App struct {
	Service    *order_application_services.OrderService
	Repository order_domain_contracts.IOrderRepository
	InputPort  order_domain_ports.OrderInputPort
	OutputPort order_domain_ports.OrderOutputPort
}

func NewApp(
	Service *order_application_services.OrderService,
	Repository order_domain_contracts.IOrderRepository,
	InputPort order_domain_ports.OrderInputPort,
	OutputPort order_domain_ports.OrderOutputPort,
) *App {
	return &App{
		Service:    Service,
		Repository: Repository,
		InputPort:  InputPort,
		OutputPort: OutputPort,
	}
}

func (app *App) Worder(ctx context.Context, )
