package DTOs

type LoginDTO struct {
	login    string `json: login`
	password string `json: password`
}

func (obj LoginDTO) getLogin() string {
	return obj.login
}

func (obj LoginDTO) getPassword() string {
	return obj.password
}

func (obj *LoginDTO) setLogin(login string) {
	obj.login = login
}

func (obj *LoginDTO) setPassword(login string) {
	obj.password = obj.password
}
