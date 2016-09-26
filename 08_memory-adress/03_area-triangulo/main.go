package main

import "fmt"

func main() {
	var base float64
	var height float64
	var area float64
	fmt.Print("Entra la base de tu triangulo: ")
	fmt.Scan(&base)
	fmt.Print("Entra la altura de tu traingulo: ")
	fmt.Scan(&height)
	fmt.Print("El area del triangulo es: ")
	area = base * height / 2
	fmt.Println(base, "es la base", height, "es la altura", area, "es el area")
	fmt.Println("La memoria de area es:",&area)
	fmt.Printf("%d\n",&area)
}
