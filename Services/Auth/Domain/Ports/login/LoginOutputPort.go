package authports

import (
	shared_adapters "delivery/Services/Shared/Domain/adapters"
)

type LoginOutputPort interface {
	GetResponse(responsemodel any) shared_adapters.ViewModel
}
