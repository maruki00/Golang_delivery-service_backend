package ValueObject

type AddressValueObject struct {
	value string
}

func (obj *AddressValueObject) SetAddress(address string) {
	obj.value = address
}

func (obj *AddressValueObject) GetAddress() string {
	return obj.value
}
