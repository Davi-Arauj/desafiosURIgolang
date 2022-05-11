package main

import (
	"fmt"
)

func main() {
	var valor int64
	var notas = [7]int64{100, 50, 20, 10, 5, 2, 1}
	fmt.Scanf("%d", &valor)
	fmt.Println(valor)

	for _, nota := range notas {
		qtdNotas := (valor / nota)
		fmt.Printf("%d nota(s) de R$ %d,00\n", qtdNotas, nota)
		valor = valor % nota
	}

}
