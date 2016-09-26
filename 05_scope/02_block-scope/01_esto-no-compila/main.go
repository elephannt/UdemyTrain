package main
import "fmt"
func main() {
	x:= 42
	fmt.Println(x)
	foo()
}
func foo(){
	// no tiene acceso a x, porque x esta en el bloque de arriba, tendria que estar en el main
	//no va compilar
	fmt.Println(x)
}