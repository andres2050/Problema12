package dist

import (
	"crypto/rand"
	"math/big"
)

func hot(ven *Venta) {
	probabilityBig, _ := rand.Int(rand.Reader, big.NewInt(100))
	probability := probabilityBig.Int64()
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
