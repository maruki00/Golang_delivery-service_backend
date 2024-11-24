package auth_domain_ports

import (
	shared_domain_contracts "delivery/Internal/Shared/Domain/Contracts"
	shared_models "delivery/Internal/Shared/Infrastructure/Models"
)

type AuthOutputPort interface {
	Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
	Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
}
