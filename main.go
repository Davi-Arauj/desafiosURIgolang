package main

import (
	"fmt"
)

func main() {
	var distancia,litros float64
	
	fmt.Scanf("%f", &distancia)
	fmt.Scanf("%f", &litros)
	consumo := (distancia) / (litros)
	fmt.Printf("%.3f km/l\n", consumo)

}
