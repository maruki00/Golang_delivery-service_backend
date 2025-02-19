package presenters

import shared_models "delivery/internal/shared/infra/models"

type JsonAuthPresenter struct {
}

func NewJSONAuthPresenter() JsonAuthPresenter {
	return JsonAuthPresenter{}
}

func (obj *JsonAuthPresenter) Success(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_viewmodels.NewJsonViewModel(data)
}

func (obj *JsonAuthPresenter) Error(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_viewmodels.NewJsonViewModel(data)
}
