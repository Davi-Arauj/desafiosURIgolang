package main

import "fmt"

func main() {
	
	var cod1,num1,cod2,num2 int64
	var valor1,valor2 float64

	fmt.Scanf("%d %d %f",&cod1,&num1,&valor1)
	fmt.Scanf("%d %d %f",&cod2,&num2,&valor2)

	total := (valor1 * float64(num1)) + (valor2 * float64(num2))

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)

}
