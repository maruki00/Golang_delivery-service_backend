package ValueObject

type EmailValueObject struct {
	value string
}

func (obj *EmailValueObject) SetEmail(email string) {
	obj.value = email
}

func (obj *EmailValueObject) GetEmail() string {
	return obj.value
}
