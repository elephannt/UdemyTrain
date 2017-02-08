package main

import "fmt"

func main() {

	var n int
	var num int
	var result int

	fmt.Println("Cuantos valores quieres calcular?")
	fmt.Scanf("%d\n", &n)
	fmt.Println("Ingrese los valores separados por espacios:")
	numList := map[int]int{}

	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &num)
		numList[i] = num
	}
	//Este resultado es para valores de 2
	result = gcd(numList[1], numList[2])

	//Este ciclo es para valores de 2 +
	for j := 3; j <= n; j++ {
		result = gcd(result, numList[j])
	}

	fmt.Println("El maximo comun divisor de los valores es: ", result)
}

//Funccion para impementar el algoritmo de Euclid Algo
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}