package dist

import "math/rand"

func nebulous(ven *Venta) {
	probability := rand.Intn(100)
	if probability >= 0 && probability < 10 {
		ven.Ventas = 10
	}
	if probability >= 10 && probability < 30 {
		ven.Ventas = 20
	}
	if probability >= 30 && probability < 70 {
		ven.Ventas = 30
	}
	if probability >= 70 && probability < 100 {
		ven.Ventas = 40
	}
}
