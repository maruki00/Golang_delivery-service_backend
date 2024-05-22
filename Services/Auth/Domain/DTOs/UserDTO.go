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


func (obj *UserDTO)  setId(setId int){
	obj.Id = Id
}
func (obj *UserDTO)  setUserName(setUserName string){
	obj.UserName = UserName
}
func (obj *UserDTO)  setFullName(setFullName string){
	obj.FullName = FullName
}
func (obj *UserDTO)  setEmail(setEmail string){
	obj.Email = Email
}
func (obj *UserDTO)  setPassword(setPassword string){
	obj.Password = Password
}
func (obj *UserDTO)  setCountry(setCountry string){
	obj.Country = Country
}
func (obj *UserDTO)  setCity(setCity string){
	obj.City = City
}
func (obj *UserDTO)  setStreet(setStreet string){
	obj.Street = Street
}
func (obj *UserDTO)  setHouse(setHouse int){
	obj.House = House
}
func (obj *UserDTO)  setFlat(setFlat int){
	obj.Flat = Flat
}

func (obj *UserDTO) getId()  int{
	return obj.Id
}
func (obj *UserDTO) getUserName()  string{
	return obj.UserName
}
func (obj *UserDTO) getFullName()  string{
	return obj.FullName
}
func (obj *UserDTO) getEmail()  string{
	return obj.Email
}
func (obj *UserDTO) getPassword() string{
	return obj.Password
}
func (obj *UserDTO) getCountry() string{
	return obj.Country
}
func (obj *UserDTO) getCity():  string{
	return obj.City
}
func (obj *UserDTO) getStreet()  string{
	return obj.Street
}
func (obj *UserDTO) getHouse() int{
	return obj.House
}
func (obj *UserDTO) getFlat()int{
	return obj.Flat
}