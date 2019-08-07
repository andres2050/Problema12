package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

type ventas struct {
	Ventas int
	Clima  string
	Dia    int
}

func distClima(ndays int) []ventas {

	ven := make([]ventas, 0)
	for i := 0; i < ndays; i++ {

		rndBig, _ := rand.Int(rand.Reader, big.NewInt(100))
		rnd := rndBig.Int64()

		if rnd >= 0 && rnd < 30 {
			day := ventas{
				Clima: "Caluroso",
				Dia:   i,
			}
			asignarVentas(&day)

			ven = append(ven, day)
		}
		if rnd >= 30 && rnd < 70 {
			ven = append(ven, ventas{
				Clima: "Nebuloso",
			})
		}
		if rnd >= 70 && rnd < 90 {
			ven = append(ven, ventas{
				Clima: "Tormentoso",
			})
		}
		if rnd >= 90 && rnd < 100 {
			ven = append(ven, ventas{
				Clima: "Espantoso",
			})
		}
	}

	return ven
}

func asignarVentas(ven *ventas) {
	ven.Ventas = 1
}

func main() {

	ven := distClima(90)
	str, _ := json.Marshal(ven)
	fmt.Println(string(str))
}
