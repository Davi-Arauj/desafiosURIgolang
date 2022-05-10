package main

import "fmt"

func main() {
	var salario,vendas float64
	var nome string

	fmt.Scanf("%s",&nome)
	fmt.Scanf("%f",&salario)
	fmt.Scanf("%f",&vendas)

	comissao := (vendas*0.15)+salario

	fmt.Printf("TOTAL = R$ %.2f\n",comissao)
}
