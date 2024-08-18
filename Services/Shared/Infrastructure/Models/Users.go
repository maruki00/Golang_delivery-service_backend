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

func (obj *User) GetId() int {
	return obj.Id
}
func (obj *User) GetUsername() string {
	return obj.Username
}
func (obj *User) GetFullname() string {
	return obj.Fullname
}
func (obj *User) GetEmail() string {
	return obj.Email
}
func (obj *User) GetAddress() string {
	return obj.Address
}
func (obj *User) GetPassword() string {
	return obj.Password
}
func (obj *User) GetUserType() string {
	return obj.UserType
}
func (obj *User) GetUserLevel() string {
	return obj.UserLevel
}
func (obj *User) GetIsOnline() string {
	return obj.IsOnline
}
func (obj *User) GetIsLocked() string {
	return obj.IsLocked
}
func (obj *User) GetLastSeen() string {
	return obj.LastSeen
}
func (obj *User) GetCreatedAt() string {
	return obj.CreatedAt
}
func (obj *User) GetUpdatedAt() string {
	return obj.UpdatedAt
}
