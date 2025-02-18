package auth_usergateway_presenters

import (
	shared_domain_contracts "delivery/internal/shared/Domain/Contracts"
	shared_adapters_viewmodels "delivery/internal/shared/UserGateway/Adapters/ViewModels"
	shared_models "delivery/internal/shared/infra/Models"
)

type JsonAuthPresenter struct {
}

func NewJSONAuthPresenter() JsonAuthPresenter {
	return JsonAuthPresenter{}
}

func (obj *JsonAuthPresenter) Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *JsonAuthPresenter) Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}
