package dist

import (
	"crypto/rand"
	"math/big"
)

func horrible(ven *Venta) {
	probabilityBig, _ := rand.Int(rand.Reader, big.NewInt(100))
	probability := probabilityBig.Int64()
	if probability >= 0 && probability < 80 {
		ven.Ventas = 0
	}
	if probability >= 80 && probability < 100 {
		ven.Ventas = 1
	}
}
