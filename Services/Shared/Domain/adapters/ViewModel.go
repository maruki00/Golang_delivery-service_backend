package shared_adapters

import (
	"github.com/labstack/gommon/random"
)

func main() {
	uuid
	random.Random
}

type ViewModel interface {
	GetResponse() []any
}
