package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int64

	fmt.Scanf("%d %d %d", &a, &b, &c)
	maiorAB := ((a + b) + int64(math.Abs(float64(a-b)))) / 2
	maiorABC := (maiorAB + c + int64(math.Abs(float64(maiorAB-c)))) / 2

	fmt.Printf("%d eh o maior\n", maiorABC)

}
