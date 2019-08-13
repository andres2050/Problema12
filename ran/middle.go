package ran

import (
	"crypto/rand"
	"math/big"
)

func randCreator(n int) []int64 {
	rands := []int64{}
	for i := n; i < n; i++ {
		rndBig, _ := rand.Int(rand.Reader, big.NewInt(100))
		rnd := rndBig.Int64()
		rands = append(rands, rnd)
	}
	return rands
}

func media(rands []int64) int64 {
	sum := int64(0)
	for i := 0; i < len(rands); i++ {
		sum = sum + rands[i]
	}
	return sum
}
