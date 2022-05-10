package main

import "fmt"

func main() {
	
    var a,b,c float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)
	var notaA = a*2.0
	var notaB = b*3.0
	var notaC = c*5.0 
 
	media := (notaA+notaB+notaC)/10

	fmt.Printf("MEDIA = %.1f\n", media)
}
