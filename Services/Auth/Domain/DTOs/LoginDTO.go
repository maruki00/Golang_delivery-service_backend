package auth_domain_dtos

type LoginDTO struct {
	Login    string `json: login`
	Password string `json: password`
}

func (obj *LoginDTO) GetLogin() string {
	return obj.Login
}

func (obj *LoginDTO) GetPassword() string {
	return obj.Password
}

func (obj *LoginDTO) SetLogin(login string) {
	obj.Login = login
}

func (obj *LoginDTO) SetPassword(password string) {
	obj.Password = password
}

func (obj *LoginDTO) Failed() bool {
	return obj.Login == "" || obj.Password == ""
}
