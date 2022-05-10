package main

import "fmt"

func main() {
	var a,b float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	var notaA = a*3.5
	var notaB = b*7.5 
	media := (notaA+notaB)/11

	fmt.Printf("MEDIA = %.5f\n", media)
}
