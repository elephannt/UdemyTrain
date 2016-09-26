package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string
	fmt.Print("Ingresa nombre: ")
	fmt.Scan(&name)
	count := 0
	for len(name) > 0 {
		r, size := utf8.DecodeLastRuneInString(name)
		count++
		fmt.Printf("%c %v\n", r, size)

		name = name[:len(name)-size]
	}
	fmt.Println("count:",count)
}
