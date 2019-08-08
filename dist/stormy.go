package dist

import (
	"math"
	"math/rand"
)

func stormy(ven *Venta) {
	max := 25
	x := rand.Intn(max)
	if x <= 0 {
		ven.Ventas = 0
	} else if x >= max {
		ven.Ventas = 25
	} else {
		ven.Ventas = int(math.Pow(float64(x), 2) / float64(max))
	}
}
