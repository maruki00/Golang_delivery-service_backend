package DTOs

type UserDTO struct {
	Id       int    `json:id`
	UserName string `json:user_name`
	FullName string `json:full_name`
	Email    string `json:email`
	Password string `json:password`
	Country  string `json:country`
	City     string `json:city`
	Street   string `json:street`
	House    string `json:house`
	Flat     string `json:flat`
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
func (obj *UserDTO) SetHouse(House string) {
	obj.House = House
}
func (obj *UserDTO) SetFlat(Flat string) {
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
func (obj *UserDTO) GetHouse() string {
	return obj.House
}
func (obj *UserDTO) GetFlat() string {
	return obj.Flat
}
