package main

import "golang.org/x/crypto/ssh/terminal"
import "fmt"

func main() {
	fmt.Println("Enter password: ")
	password, err := terminal.ReadPassword(0)
	if password {

	} else {

	}
}