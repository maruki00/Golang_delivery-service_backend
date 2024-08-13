package auth_adapters

type LoginPresenter struct {
}

func (obj *LoginPresenter) InvalidCredentials() string {
	return "helloworld"
}
func (obj *LoginPresenter) Success() string {
	return "helloworld"
}
func (obj *LoginPresenter) Error() string {
	return "helloworld"
}
