package auth_infrastructure_models

type TwoFactory struct {
	Pin int `json:"pin"`
}

func (o *TwoFactory) GetPin() int {
	return o.Pin
}
