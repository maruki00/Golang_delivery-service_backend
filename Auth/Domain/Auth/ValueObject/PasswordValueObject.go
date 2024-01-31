package ValueObject

type PasswordValueObject struct {
	value string
}

func (obj *PasswordValueObject) SetPassword(password string) {
	obj.value = password
}

func (obj *PasswordValueObject) ToString() string {
	return obj.value
}
