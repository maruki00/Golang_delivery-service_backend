package product_usergateway_adapters_presenters

import (
	shared_domain_contracts "delivery/Internal/Shared/Domain/Contracts"
	shared_models "delivery/Internal/Shared/Infrastructure/Models"
	shared_adapters_viewmodels "delivery/Internal/Shared/UserGateway/Adapters/ViewModels"
)

type ProductPresenter struct {
}

func (obj *ProductPresenter) Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *ProductPresenter) Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}
