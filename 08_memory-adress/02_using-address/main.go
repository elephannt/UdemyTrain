package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam:")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meter is ", yards, "yards.")
	fmt.Println("Yards memory address - ",&yards)
	fmt.Printf("%d\n",&yards)
}