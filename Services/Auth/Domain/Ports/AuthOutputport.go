package auth_domain_ports

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
)

type AuthOutputPort interface {
	Success(response any) domain_auth_contracts.ViewModel
	Error(response any) domain_auth_contracts.ViewModel
}
