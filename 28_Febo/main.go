package main
import "fmt"
var a1 int = 1
var a2 int = 1
var cont int = 0
var n = 0

func main() {

	fmt.Println("Ingrese el numero de terminos para la susecion: ")
	fmt.Scan(&n)
	for cont = 0; cont<=n; cont++{

		cont = a1 + a2
		a1 = a2
		a2 = cont
		fmt.Println("Serie: ", cont)
	}

}