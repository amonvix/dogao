package dogao

import (
	"fmt"

	"github.com/amonvix/dog"
)

func Latir() string {
	return "WOOF!"
}

func Latidos() string {
	return "WOOF! WOOF! WOOF!"
}

func LatidoAlto() string {
	return dog.QuandoCrescer(Latir())
}

func LatidoAltos() string {
	return dog.QuandoCrescer(Latidos())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}
func From12() {
	fmt.Println("I'm from version 1.2.0")

}
