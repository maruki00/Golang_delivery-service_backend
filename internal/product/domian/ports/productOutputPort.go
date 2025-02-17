package product_domain_ports

import (
	shared_domain_contracts "delivery/internal/shared/Domain/Contracts"
	shared_models "delivery/internal/shared/infra/Models"
)

type ProductOutputPort interface {
	Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
	Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
}
