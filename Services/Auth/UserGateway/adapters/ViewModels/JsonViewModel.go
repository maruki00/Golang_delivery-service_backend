package auth_adapters_viewmodels

import (
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"encoding/json"
)

type JsonViewModel struct {
	data *shared_models.ResponseModel
}

func NewJsonViewModel(data *shared_models.ResponseModel) JsonViewModel {
	return JsonViewModel{
		data: data,
	}
}

func (obj *JsonViewModel) GetResponse() *shared_models.ResponseModel {
	return obj.data
}

func (obj *JsonViewModel) String() string {
	d, err := json.Marshal(obj.data)
	if err != nil {
		return "error data could not be serialized"
	}
	return string(d)
}
