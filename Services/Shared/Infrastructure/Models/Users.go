package shared_models

type User struct {
	Id        int    `json: "id"`
	Username  string `json: "user_name"`
	Fullname  string `json: "full_name"`
	Email     string `json: "email"`
	Address   string `json: "address"`
	Password  string `json: "password"`
	UserType  string `json: "user_type"`
	UserLevel string `json: "user_level"`
	IsOnline  string `json: "is_online"`
	IsLocked  string `json: "is_locked"`
	LastSeen  string `json: "last_seen"`
	CreatedAt string `json: "created_at"`
	UpdatedAt string `json: "updated_at"`
}
