package main

import (
	"fmt"
	"regexp"
)

func validateDigits(digits string) bool {
	Re := regexp.MustCompile(`^[0-9]+$`)
	return Re.MatchString(digits)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {
	digits := "13920a"
	if !validateDigits(digits) {
		fmt.Println("digits is invalid")
	} else {
		fmt.Println("digits is VALID")
	}
	digits = "a1"
	if !validateDigits(digits) {
		fmt.Println("digits is invalid")
	} else {
		fmt.Println("digits is VALID")
	}
}