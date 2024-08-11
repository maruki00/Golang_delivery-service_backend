package DTOs

type ForgetPasswordDTO struct {
	email string `json: email`
}

func (obj *ForgetPasswordDTO) GetEmail() string {
	return obj.email
}

func (obj *ForgetPasswordDTO) SetLogin(email string) {
	obj.email = email
}
