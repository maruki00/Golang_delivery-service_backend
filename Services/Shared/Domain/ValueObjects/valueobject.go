package SharedValueObject

type ValueObject interface {
	SetValue(valu any)
	GetValue() any
	ToSTring() string
}
