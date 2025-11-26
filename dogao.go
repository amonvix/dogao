package dogao

import "github.com/amonvix/dog"

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
