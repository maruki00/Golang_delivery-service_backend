package auth_usergateway_presenters

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	auth_adapters_viewmodels "delivery/Services/Auth/UserGateway/adapters/ViewModels"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
)

type JSONAuthPresenter struct {
}

func (obj *JSONAuthPresenter) Success(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return auth_adapters_viewmodels.NewJsonViewModel(shared_models.ResponseModel{
		Status:  200,
		Message: "Operation Done with success",
		Error:   "",
		Data:    data,
	})
}

func (obj *JSONAuthPresenter) Error(data any) domain_auth_contracts.ViewModel {
	return auth_adapters_viewmodels.NewJsonViewModel(shared_models.ResponseModel{
		Status:  400,
		Message: "Operation was failed",
		Error:   data,
		Data:    nil,
	})
}
