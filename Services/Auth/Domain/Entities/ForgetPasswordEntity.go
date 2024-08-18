package auth_domain_entities

type ForgetPasswordEntity interface {
	GetId() int
	GetEmail() string
	GetToken() string
}
