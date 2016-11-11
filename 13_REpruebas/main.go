package main

import (
	"fmt"
	"regexp"
)
//y termine con numero o letra, y al final tenga un "@" para correo.
func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^(amor)+$`)

	//ValidID = valida una palabra que termine un numero, si empieza con letra no compila.
	fmt.Println(validID.MatchString(""))


}