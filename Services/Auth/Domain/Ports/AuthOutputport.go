package auth_domain_ports

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
)

type AuthOutputPort interface {
	Success(response shared_models.ResponseModel) domain_auth_contracts.ViewModel
	Error(response shared_models.ResponseModel) domain_auth_contracts.ViewModel
}
