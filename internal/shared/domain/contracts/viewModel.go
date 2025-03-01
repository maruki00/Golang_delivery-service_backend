package shared_contracts

import shared_models "delivery/internal/shared/infra/models"

type ViewModel interface {
	GetResponse() shared_models.ResponseModel
	String() string
}
