package main

import "fmt"

func main() {
	
    var a float64
	n := 3.14159
	var A float64

	fmt.Scanf("%f", &a)

	A = (a * a) * n

	fmt.Printf("A=%.4f\n", A)
}
