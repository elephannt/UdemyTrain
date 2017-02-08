package main

import (
	"fmt"
	"math/rand"
	"os"
)

var coin = 0
var amount = 0
var menu string

func main() {

	fmt.Println("Welmcome to the best system, would you like to place a bet?:")
	fmt.Scan(&menu)
	if menu == "Yes" || menu == "yes" {
	L:
		fmt.Println("Enter amount to bet:")
		fmt.Scan(&amount)
		coin = rand.Int() % 2

		if coin == 0 {
			fmt.Println("Heads, you won double")
			fmt.Println("You have won: \n", amount*2)

		} else {
			fmt.Println("You lose")
		}
		fmt.Println("Would you like to keep betting?:")
		fmt.Scan(&menu)
		if menu == "Yes" || menu == "yes" {
			goto L
		} else {
			os.Exit(0)
		}
	} else {

		fmt.Println("Closing program...")
	}

}
