package main

import (
	"os"
	"fmt"


	"io/ioutil"
	"bytes"
	"bufio"
	"strings"
)



var nombreTL string
var path string = "C:\\MainDB\\testing.txt"
func main() {
	//file, err := os.Create("C:\\MainDB\\testing.txt")
	//if err != nil {
	//	return
	//}
	//defer file.Close()
	//
	//fmt.Println("Escribe algo: ")
	//scan := bufio.NewScanner(os.Stdin)
	//scan.Scan()
	//file.WriteString(scan.Text())
	fmt.Print(" | GoSQL]-> ")

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	scann := strings.Fields(scan.Text())
	nombreTL = scann[2]
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte(nombreTL), []byte(""), -1)

	if err = ioutil.WriteFile(path, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}



}
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}