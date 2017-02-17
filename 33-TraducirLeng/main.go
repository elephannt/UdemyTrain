package main

import "fmt"

var nom string
var edad int
var menu string

var n int
var suma int
var lado int

func main() {

	////Programa uno
	//
	//fmt.Println("Ingrese la nombre:")
	//fmt.Scan(&nom)
	//fmt.Println("Ingrese la edad:")
	//fmt.Scan(&edad)
	//
	//fmt.Println("Su nombre es: ",nom," Su edad es: ",edad)
	fmt.Println("Ingrese los lados: ")
	fmt.Scan(&n)
	fmt.Println("Ingrese los lados separados por espacios:")
	numList := map[int]int{}

	for i := 0; i <= n; i++ {

		fmt.Scanf("%d", &lado)
		numList[i] = lado
		suma = suma + lado

	}
	fmt.Println("\nEl perimetro del poligono es: ",suma)
}
