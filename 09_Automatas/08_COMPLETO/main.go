package main

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
	return Re.MatchString(email)
}

func validateURL(url string) bool {
	//Re := regexp.MustCompile(`^[www.|http://|https://]*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	Re := regexp.MustCompile(`^(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	//[A-Za-z0-9./?=_%+\-]* al final para youtube.com/WATCH_03
	return Re.MatchString(url)
}
func validateNUM(num string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z]+[0-9]+$`)
	return Re.MatchString(num)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateDigits(digits string) bool {
	Re := regexp.MustCompile(`^[0-9]+$`)
	return Re.MatchString(digits)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {
	var global string
	fmt.Print("Ingresa dato: ")
	fmt.Scan(&global)
	 if validateEmail(global) {
		fmt.Println("Es un Email")
	} else if validateURL(global) {
		fmt.Println("Es un URL")
	} else if validateNUM(global){
		fmt.Println("Es un ID")
	} else if validateDigits(global){
		fmt.Println("Es un Numero")
	} else {
		 fmt.Println("Error, caracter indefinido")
	 }

}
