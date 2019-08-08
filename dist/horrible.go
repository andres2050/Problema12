package dist

import "math/rand"

func horrible(ven *Venta) {
	probability := rand.Intn(100)
	if probability >= 0 && probability < 80 {
		ven.Ventas = 0
	}
	if probability >= 80 && probability < 100 {
		ven.Ventas = 1
	}
}
