package auth_domain_ports

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
)

type AuthOutputPort interface {
	Success(data shared_models.ResponseModel) domain_auth_contracts.ViewModel
	Error(data shared_models.ResponseModel) domain_auth_contracts.ViewModel
}
