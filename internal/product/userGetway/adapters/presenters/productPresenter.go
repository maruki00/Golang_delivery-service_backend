package product_usergateway_adapters_presenters

import (
	shared_domain_contracts "delivery/internal/shared/Domain/Contracts"
	shared_adapters_viewmodels "delivery/internal/shared/UserGateway/Adapters/ViewModels"
	shared_models "delivery/internal/shared/infra/Models"
)

type ProductPresenter struct {
}

func (obj *ProductPresenter) Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *ProductPresenter) Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}
