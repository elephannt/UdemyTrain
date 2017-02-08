package main

import (
	"fmt"
	"github.com/howeyc/gopass"
)

func main() {
	var username string
	fmt.Print("\nEnter username: ")
	fmt.Scan(&username)

	//fmt.Printf("Enter silent password: ")
	silentPassword, _ := gopass.GetPasswd() // Silent
	fmt.Print(string(silentPassword))

	fmt.Printf("Enter masked password: ")
	maskedPassword, _ := gopass.GetPasswdMasked() // Masked
	if username == "marco" && string(maskedPassword) == "123" {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}
