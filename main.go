package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func distClima(ndays int) {

	for i := 0; i < ndays; i++ {

		rndBig, _ := rand.Int(rand.Reader, big.NewInt(100))
		rnd := rndBig.Int64()

		if rnd >= 0 && rnd < 30 {
			fmt.Printf("Caluroso: %d\n", i+1)
		}
		if rnd >= 30 && rnd < 70 {
			fmt.Printf("Nebuloso: %d\n", i+1)
		}
		if rnd >= 70 && rnd < 90 {
			fmt.Printf("Tormentoso: %d\n", i+1)
		}
		if rnd >= 90 && rnd < 100 {
			fmt.Printf("Espantoso: %d\n", i+1)
		}
	}
}

func main() {

	distClima(90)

}
