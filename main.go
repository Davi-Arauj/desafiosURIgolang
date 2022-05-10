package main

import "fmt"

func main() {
	var a,b,c,d int64

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)

	DIFERENCA := (a * b - c * d)
 
	fmt.Printf("DIFERENCA = %d\n", DIFERENCA)
}
