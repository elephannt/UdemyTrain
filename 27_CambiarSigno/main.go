package main

import "fmt"

var numero int = 0
var contador int = 0
var n int = 0

func main() {

	fmt.Println("Terms you would like to use:")
	fmt.Scan(&n)
	for contador = 1; contador <= n; contador++ {
		if contador%2 == 0 {
			fmt.Print(" + ", contador)
		} else {
			fmt.Print(" - ", contador)
		}

	}
}
