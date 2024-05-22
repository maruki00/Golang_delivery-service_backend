package DTOs

type UserDTO struct {
	Id       int
	UserName string
	FullName string
	Email    string
	Password string
	Country  string
	City     string
	Street   string
	House    int
	Flat     int
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
func (obj *UserDTO) SetCountry(Country string) {
	obj.Country = Country
}
func (obj *UserDTO) SetCity(City string) {
	obj.City = City
}
func (obj *UserDTO) SetStreet(Street string) {
	obj.Street = Street
}
func (obj *UserDTO) SetHouse(House int) {
	obj.House = House
}
func (obj *UserDTO) SetFlat(Flat int) {
	obj.Flat = Flat
}

func (obj *UserDTO) GetId() int {
	return obj.Id
}
func (obj *UserDTO) GetUserName() string {
	return obj.UserName
}
func (obj *UserDTO) GetFullName() string {
	return obj.FullName
}
func (obj *UserDTO) GetEmail() string {
	return obj.Email
}
func (obj *UserDTO) GetPassword() string {
	return obj.Password
}
func (obj *UserDTO) GetCountry() string {
	return obj.Country
}
func (obj *UserDTO) GetCity() string {
	return obj.City
}
func (obj *UserDTO) GetStreet() string {
	return obj.Street
}
func (obj *UserDTO) GetHouse() int {
	return obj.House
}
func (obj *UserDTO) GetFlat() int {
	return obj.Flat
}
