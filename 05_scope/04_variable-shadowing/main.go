package main

import "fmt"

func max(x int) int {
	return 42 + x

}
func main() {
	max := max(7)
	fmt.Println(max)//max es el resultado ahora, no la funcion
}
//no hagan esto, es codigo malo para shadow variables