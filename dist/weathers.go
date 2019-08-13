package dist

import (
	"crypto/rand"
	"math/big"
)

//Venta asas
type Venta struct {
	Ventas int
	Clima  string
	Dia    int
}

//DistClima asas
func DistClima(ndays int) []Venta {

	ven := make([]Venta, 0)
	for i := 0; i < ndays; i++ {

		rndBig, _ := rand.Int(rand.Reader, big.NewInt(100))
		rnd := rndBig.Int64()

		if rnd >= 0 && rnd < 30 {
			day := Venta{
				Clima: "Caluroso",
				Dia:   i,
			}
			hot(&day)
			ven = append(ven, day)
		}
		if rnd >= 30 && rnd < 70 {
			day := Venta{
				Clima: "Nebuloso",
				Dia:   i,
			}
			nebulous(&day)
			ven = append(ven, day)
		}
		if rnd >= 70 && rnd < 90 {
			day := Venta{
				Clima: "Tormentoso",
				Dia:   i,
			}
			stormy(&day)
			ven = append(ven, day)
		}
		if rnd >= 90 && rnd < 100 {
			day := Venta{
				Clima: "Espantoso",
				Dia:   i,
			}
			horrible(&day)
			ven = append(ven, day)
		}
	}

	return ven
}
