package auth_domain_dto

type FrogetPasswordDTO struct {
	email string `json: email`
}

func (obj *FrogetPasswordDTO) GetEmail() string {
	return obj.email
}

func (obj *FrogetPasswordDTO) SetEmail(email string) {
	obj.email = email
}
