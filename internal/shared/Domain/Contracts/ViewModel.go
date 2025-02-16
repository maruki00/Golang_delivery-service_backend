package shared_domain_contracts

import shared_models "delivery/internal/shared/Infrastructure/Models"

type ViewModel interface {
	GetResponse() shared_models.ResponseModel
	String() string
}
