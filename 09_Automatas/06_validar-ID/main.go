package main

import (
	"fmt"
	"regexp"
)

func validateNUM(num string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z]+[0-9]+$`)
	return Re.MatchString(num)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {
	num := "marco1"
	if !validateNUM(num) {
		fmt.Println("Number is invalid")
	} else {
		fmt.Println("Number is valid")
	}
	num = "marco1"
	if !validateNUM(num) {
		fmt.Println("Number is invalid")
	} else {
		fmt.Println("Number is valid")
	}
}
