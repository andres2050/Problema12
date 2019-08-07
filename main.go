package main

import (
	"crypto/rand"
	"math/big"
)

type ventas struct {
	Ventas int
	Clima  string
}

func distClima(ndays int) []ventas {

	ven := make([]ventas, 0)
	for i := 0; i < ndays; i++ {

		rndBig, _ := rand.Int(rand.Reader, big.NewInt(100))
		rnd := rndBig.Int64()

		if rnd >= 0 && rnd < 30 {
			ven = append(ven, ventas{
				Ventas: 1,
				Clima:  "Caluroso",
			})
		}
		if rnd >= 30 && rnd < 70 {
			ven = append(ven, ventas{
				Ventas: 1,
				Clima:  "Nebuloso",
			})
		}
		if rnd >= 70 && rnd < 90 {
			ven = append(ven, ventas{
				Ventas: 1,
				Clima:  "Tormentoso",
			})
		}
		if rnd >= 90 && rnd < 100 {
			ven = append(ven, ventas{
				Ventas: 1,
				Clima:  "Espantoso",
			})
		}
	}

	return ven
}

func main() {

	distClima(90)

}
