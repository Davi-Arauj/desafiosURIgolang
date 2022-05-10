package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	pi := 3.14159
	fmt.Scanf("%f %f %f", &a, &b, &c)
	triangulo := (a * c) / 2
	circulo := pi * (math.Pow(c, 2))
	trapezio := ((a + b) * c) / 2
	quadrado := math.Pow(b, 2)
	retangulo := (a*b)

	fmt.Printf("TRIANGULO: %.3f\n", triangulo)
	fmt.Printf("CIRCULO: %.3f\n", circulo)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezio)
	fmt.Printf("QUADRADO: %.3f\n", quadrado)
	fmt.Printf("RETANGULO: %.3f\n", retangulo)


}
