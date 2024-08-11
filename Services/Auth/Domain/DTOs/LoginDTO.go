package auth_domain_dto

type LoginDTO struct {
	login    string `json: login`
	password string `json: password`
}

func (obj *LoginDTO) GetLogin() string {
	return obj.login
}

func (obj *LoginDTO) GetPassword() string {
	return obj.password
}

func (obj *LoginDTO) SetLogin(login string) {
	obj.login = login
}

func (obj *LoginDTO) SetPassword(password string) {
	obj.password = password
}

func (obj *LoginDTO) Failed() bool {
	return obj.login == "" || obj.password == ""
}
