package auth_domain_dtos

type RegisterDTO struct {
	UserName  string
	FullName  string
	Email     string
	Address   string
	Password  string
	UserType  string
	UserLevel string
	IsOnline  string
	IsLocked  string
	LastSeen  string
}
