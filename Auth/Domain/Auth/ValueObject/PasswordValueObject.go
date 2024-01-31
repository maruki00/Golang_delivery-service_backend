package ValueObject

type PasswordValueObject struct {
	value string
}

func (obj *EmailValueObject) SetPassword(password string) {
	obj.value = password
}

func (obj *EmailValueObject) GetPassword() string {
	return obj.value
}
