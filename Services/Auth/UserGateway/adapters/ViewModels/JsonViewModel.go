package auth_adapters_viewmodels

import (
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"encoding/json"
)

type JsonViewModel struct {
	Data shared_models.ResponseModel
}

func NewJsonViewModel(data shared_models.ResponseModel) JsonViewModel {
	return JsonViewModel{
		Data: data,
	}
}

// func (obj *JsonViewModel) GetResponse() *shared_models.ResponseModel {
// 	return obj.data
// }

// func (obj *JsonViewModel) String() string {
// 	d, err := json.Marshal(obj.data)
// 	if err != nil {
// 		return "error data could not be serialized"
// 	}
// 	return string(d)
// }

func (obj JsonViewModel) GetResponse() shared_models.ResponseModel {
	return obj.Data
}
func (obj JsonViewModel) String() string {
	d, err := json.Marshal(obj.Data)
	if err != nil {
		return "error data could not be serialized"
	}
	return string(d)
}