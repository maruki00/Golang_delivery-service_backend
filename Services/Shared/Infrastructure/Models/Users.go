package shared_models

type User struct {
	Id        int
	Username  string
	Fullname  string
	Email     string
	Address   string
	Password  string
	UserType  string
	UserLevel string
	IsOnline  string
	IsLocked  string
	LastSeen  string
	CreatedAt string
	UpdatedAt string
}
