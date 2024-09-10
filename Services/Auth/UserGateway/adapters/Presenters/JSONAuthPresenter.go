package auth_usergateway_presenters

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	auth_adapters_viewmodels "delivery/Services/Auth/UserGateway/adapters/ViewModels"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
)

type JSONAuthPresenter struct {
}

func (obj *JSONAuthPresenter) Success(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return auth_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *JSONAuthPresenter) Error(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return auth_adapters_viewmodels.NewJsonViewModel(data)
}
