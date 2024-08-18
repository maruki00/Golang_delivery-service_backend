package auth_infrastructure_models

type ResetPassword struct {
	Id        int
	Email     string
	Pin       int
	Token     string
	CreatedAt string
	UpdatedAt string
}
