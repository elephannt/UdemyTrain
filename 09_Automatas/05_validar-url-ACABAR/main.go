package main

import (
	"fmt"
	"regexp"
)

func validateURL(url string) bool {
	Re := regexp.MustCompile(`^[http|https]+://+[www]+\.[A-Za-z0-9._%+\-]+\.[com|org|edu]{3}$`)
	return Re.MatchString(url)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {
	url := "https://www.facebook.com"
	if !validateURL(url) {
		fmt.Println("URL address is invalid")
	} else {
		fmt.Println("URL address is VALID")
	}
	url = "www.google.org"
	if !validateURL(url) {
		fmt.Println("URL address is invalid")
	} else {
		fmt.Println("URL address is VALID")
	}
}
