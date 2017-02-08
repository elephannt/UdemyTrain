package main

import "fmt"

var contador = 0
var sumatoria = 0

func main() {

	fmt.Println("Lenguajes de programacion.\n")
	for i := 0; i < 100; i++ {
		contador = contador + 1
		sumatoria = sumatoria + contador
		fmt.Println("Sumatorias:", sumatoria , "\nEl contador va en: ", contador)

	}
	fmt.Println("\nEl valor de la sumatoria[1 + 2 + 3...+100 es: ]", sumatoria)
}
