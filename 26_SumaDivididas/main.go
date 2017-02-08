package main

import (
	"fmt"

)

var dividiendo float32 = 0
var divisor  float32 = 0
var suma float32 = 0
var n  int = 0
var contador = 0

func main() {

	fmt.Println("Terms that you would like to use?: ")
	fmt.Scan(&n)
	dividiendo = 1
	divisor = 1

	for contador = 0; contador < n; contador++ {
		fmt.Println("Dividiendo: ", dividiendo, "Divisor", divisor)

		suma += dividiendo  / divisor
		dividiendo = divisor
		divisor++
	}

	fmt.Println("La suma es: ", suma)

}
