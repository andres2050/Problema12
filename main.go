package main

import (
	"fmt"
	"io/ioutil"

	"github.com/andres2050/Problema12/dist"
	"github.com/andres2050/Problema12/ran"
)

func main() {
	days := 90
	ven := dist.DistClima(days)
	dataToExcel(ven)
	ran.PruebaDeMedias(days)
}

func dataToExcel(data []dist.Venta) {
	text := "Dia;Clima;Ventas\n"
	for _, v := range data {
		text += fmt.Sprintf("%d;%s;%d\n", v.Dia, v.Clima, v.Ventas)
	}
	err := ioutil.WriteFile("ventas.csv", []byte(text), 0755)
	if err != nil {
		fmt.Println("Cierre el archivo antes de generar otro\nError:", err.Error())
	}
}
