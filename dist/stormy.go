package dist

import (
	"crypto/rand"
	"math"
	"math/big"
)

func stormy(ven *Venta) {
	max := int64(25)
	probabilityBig, _ := rand.Int(rand.Reader, big.NewInt(100))
	x := probabilityBig.Int64()
	if x <= 0 {
		ven.Ventas = 0
	} else if x >= max {
		ven.Ventas = 25
	} else {
		ven.Ventas = int(math.Pow(float64(x), 2) / float64(max))
	}
}
