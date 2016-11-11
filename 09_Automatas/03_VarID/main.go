package main

import (
	"fmt"
	"regexp"
)
//y termine con numero o letra, y al final tenga un "@" para correo.
func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-zA-Z][0-9a-zA-Z]+$`)

	//ValidID = valida una palabra que termine un numero, si empieza con letra no compila.
	fmt.Println(validID.MatchString("marco1223a2"))
	fmt.Println(validID.MatchString("ksksks"))
	fmt.Println(validID.MatchString("le1"))


}
// var validID = regexp.MustCompile(`^[a-z]+$`) esto es para ver si es nombre o una palabra "x"
//var validID = regexp.MustCompile(`^[a-z]+[0-9]+$`) indentifica una palabra seguida de un numero.
//var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
//var validID = regexp.MustCompile(`^[0-9]|[a-z]+[0-9]+[@]+$`) identifica una palabra que empiece con numero o letra
