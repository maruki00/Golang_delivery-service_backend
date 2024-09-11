package auth_usergateway_presenters

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"

	shared_models "delivery/Services/Shared/Infrastructure/Models"
	shared_adapters_viewmodels "delivery/Services/Shared/UserGateway/Adapters/ViewModels"
)

type JsonAuthPresenter struct {
}

func NewJSONAuthPresenter() JsonAuthPresenter {
	return JsonAuthPresenter{}
}

func (obj *JsonAuthPresenter) Success(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *JsonAuthPresenter) Error(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}
