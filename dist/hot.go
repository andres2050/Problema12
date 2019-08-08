package dist

import "math/rand"

func hot(ven *Venta) {
	probability := rand.Intn(100)
	if probability >= 0 && probability < 20 {
		ven.Ventas = 20
	}
	if probability >= 20 && probability < 50 {
		ven.Ventas = 30
	}
	if probability >= 50 && probability < 100 {
		ven.Ventas = 40
	}
}
