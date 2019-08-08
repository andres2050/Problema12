package dist

import "math/rand"

func stormy(ven *Venta) {
	max := 25
	x := rand.Intn(max)
	if x <= 0 {
		ven.Ventas = 0
	} else if x >= max {
		ven.Ventas = 25
	} else {
		density := x / max
		ven.Ventas = int(x * density)
	}
}
