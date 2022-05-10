package main

import "fmt"

func main() {
	
	var numFuncionario,numHoras int64
	var valorHora float64

	fmt.Scanf("%d",&numFuncionario)
	fmt.Scanf("%d",&numHoras)
	fmt.Scanf("%f",&valorHora)

	salary := valorHora*float64(numHoras)

	fmt.Printf("NUMBER = %d\n",numFuncionario)
	fmt.Printf("SALARY = U$ %.2f\n",salary)
}
