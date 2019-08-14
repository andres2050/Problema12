package ran

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

//RandCreator aas
func RandCreator(n int) []float64 {
	rands := []float64{}
	for i := 0; i < n; i++ {
		rndBig, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
		rnd := float64(rndBig.Int64()) / float64((n))
		rands = append(rands, rnd)
	}

	return rands
}

//Media asas
func Media(rands []float64) float64 {
	sum := float64(0)

	for i := 0; i < len(rands); i++ {
		sum = sum + rands[i]
	}
	if len(rands) > 0 {
		k := sum / float64(len(rands))
		return k
	}

	return 0
}

//Varianza varianza
func Varianza(rands []float64, middle float64) float64 {
	x := middle
	sum := 0.0
	for i := 0; i < len(rands); i++ {
		sum = math.Pow(rands[i]+x, 2)
	}

	return sum / float64(len(rands))
}

//PruebaDeMedias realiza la prueba de medidas para el generador usado.
func PruebaDeMedias(numbers int) {
	var rands = RandCreator(numbers)
	var middle = Media(rands)
	varian := Varianza(rands, middle)
	limiteSuperior := middle + 1.96*(varian/(math.Sqrt(float64(numbers))))
	limiteInferior := middle - 1.96*(varian/(math.Sqrt(float64(numbers))))
	if middle < limiteSuperior {
		if middle > limiteInferior {
			fmt.Printf("Prueba media superada: %f > %f > %f\n", limiteSuperior, middle, limiteInferior)
		}
	} else {
		fmt.Printf("Prueba media no superada: %f > %f > %f\n", limiteSuperior, middle, limiteInferior)
	}
	fmt.Printf("Limite Superior: %f\n", limiteSuperior)
	fmt.Printf("Media: %f\n", middle)
	fmt.Printf("Limite Inferior: %f\n", limiteInferior)
}
