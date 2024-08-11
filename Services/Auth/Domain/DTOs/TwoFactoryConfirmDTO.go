package auth_domain_dtos

type TwoFactoryConfirmDTO struct {
	email string `json: email`
	pin   string `json: pin`
}

func (obj *TwoFactoryConfirmDTO) GetPin() string {
	return obj.pin
}

func (obj *TwoFactoryConfirmDTO) SetPin(pin string) {
	obj.pin = pin
}

func (obj *TwoFactoryConfirmDTO) GetEmail() string {
	return obj.email
}

func (obj *TwoFactoryConfirmDTO) SetEmail(email string) {
	obj.email = email
}
