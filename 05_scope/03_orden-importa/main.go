package main

import "fmt"

func main() {
	fmt.Println(x)
	x := 42 //no lo puede usar, porque no esta ordenado
	fmt.Println(y)
}
var y  = 42 //como esta afuera si lo acepta dentro de los braces, es global.