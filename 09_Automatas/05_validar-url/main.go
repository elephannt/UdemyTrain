package main

import (
	"fmt"
	"regexp"
)

func validateURL(url string) bool {
	//Re := regexp.MustCompile(`^[www.|http://|https://]*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	Re := regexp.MustCompile(`^(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	//[A-Za-z0-9./?=_%+\-]* al final para youtube.com/WATCH_03
	return Re.MatchString(url)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}

func main() {
	url := "https://wwww.youtube.com"
	if !validateURL(url) {
		fmt.Println("URL address is invalid")
	} else {
		fmt.Println("URL address is VALID")
	}
	url = "www.google"
	if !validateURL(url) {
		fmt.Println("URL address is invalid")
	} else {
		fmt.Println("URL address is VALID")
	}
}
