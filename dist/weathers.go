package dist

import (
	"crypto/rand"
	"math/big"
)

type Venta struct {
	Ventas int
	Clima  string
	Dia    int
}

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
			asignarVentas(&day)
			ven = append(ven, day)
		}
		if rnd >= 30 && rnd < 70 {
			ven = append(ven, Venta{
				Clima: "Nebuloso",
			})
		}
		if rnd >= 70 && rnd < 90 {
			ven = append(ven, Venta{
				Clima: "Tormentoso",
			})
		}
		if rnd >= 90 && rnd < 100 {
			ven = append(ven, Venta{
				Clima: "Espantoso",
			})
		}
	}

	return ven
}

func asignarVentas(ven *Venta) {
	ven.Ventas = 1
}
