package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello"
	count := 0
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		count++
		fmt.Printf("%c %v\n", r, size)

		str = str[:len(str)-size]
	}
	fmt.Println("count:",count)
}
