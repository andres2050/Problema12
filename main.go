package main

import (
	"encoding/json"
	"fmt"

	"github.com/andres2050/Problema12/dist"
)

func main() {

	ven := dist.DistClima(90)
	str, _ := json.Marshal(ven)
	fmt.Println(string(str))
}
