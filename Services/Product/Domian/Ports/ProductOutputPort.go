package product_domain_ports

import (
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
)

type ProductOutputPort interface {
	Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
	Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
}
