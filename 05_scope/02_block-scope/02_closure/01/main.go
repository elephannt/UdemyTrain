package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "El credito pertenece al que esta dentro del ring"
		fmt.Println(y)
	}
	//fmt.println(y) // afuera de la visibilidad de y
}
