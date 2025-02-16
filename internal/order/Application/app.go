package order_app

import (
	order_application_services "delivery/internal/order/Application/Services"
	order_domain_contracts "delivery/internal/order/Domain/Contracts"
	order_domain_ports "delivery/internal/order/Domain/Ports"
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

// func (app *App) Worder(ctx context.Context, )
