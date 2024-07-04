package DTOs

type UserDTO struct {
	Id       int    `json:id`
	UserName string `json:user_name`
	FullName string `json:full_name`
	Email    string `json:email`
	Password string `json:password`
	Address  string `json:address`
	UserType string `json:user_type`
}

func (obj *UserDTO) SetId(Id int) {
	obj.Id = Id
}
func (obj *UserDTO) SetUserName(UserName string) {
	obj.UserName = UserName
}
func (obj *UserDTO) SetFullName(FullName string) {
	obj.FullName = FullName
}
func (obj *UserDTO) SetEmail(Email string) {
	obj.Email = Email
}
func (obj *UserDTO) SetPassword(Password string) {
	obj.Password = Password
}
func (obj *UserDTO) SetAddress(Address string) {
	obj.Address = Address
}
func (obj *UserDTO) SetType(Type string) {
	obj.UserType = Type
}

// hello world
func (obj UserDTO) GetId() int {
	return obj.Id
}
func (obj UserDTO) GetUserName() string {
	return obj.UserName
}
func (obj UserDTO) GetFullName() string {
	return obj.FullName
}
func (obj UserDTO) GetEmail() string {
	return obj.Email
}
func (obj UserDTO) GetPassword() string {
	return obj.Password
}
func (obj UserDTO) GetAddress() string {
	return obj.Address
}
func (obj UserDTO) GetType() string {
	return obj.UserType
}
