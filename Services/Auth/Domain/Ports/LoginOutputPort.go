package authports

import adapters_viewmodels "delivery/app/adapters/ViewModels"

type LoginOutputPort interface {
	GetResponse() adapters_viewmodels.JsonViewModel
}
