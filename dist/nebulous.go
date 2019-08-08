package dist

import (
	"crypto/rand"
	"math/big"
)

func nebulous(ven *Venta) {
	probabilityBig, _ := rand.Int(rand.Reader, big.NewInt(100))
	probability := probabilityBig.Int64()
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
