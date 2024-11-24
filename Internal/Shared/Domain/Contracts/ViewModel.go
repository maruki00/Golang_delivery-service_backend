package shared_domain_contracts

import shared_models "delivery/Internal/Shared/Infrastructure/Models"

type ViewModel interface {
	GetResponse() shared_models.ResponseModel
	String() string
}
