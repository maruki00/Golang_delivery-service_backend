package auth_domain_login_ports

type LoginOutputPort interface {
	InvalidCredentials() string
	Success() string
	Error() string
}
