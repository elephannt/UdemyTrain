package main
import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"regexp"
)
func checkID (palabra string) bool{
	patronID := regexp.MustCompile("^[A-Za-z][A-Za-z0-9]+$")
	return patronID.MatchString(palabra)
}
func checkNum (numero string) bool{
	patronNumerico := regexp.MustCompile("^[0-9]+$")
	return patronNumerico.MatchString(numero)
}

func main() {
	file, _ := os.Open("C:/Users/cesar/Documents/Escuela/7mo Semestre/Lenguajes Automatas/Go/Lexico/lexico.txt")
	scanner := bufio.NewScanner(file)
	//parts := strings.Split(line," ")
	for scanner.Scan(){
		line := scanner.Text()
		if checkID(line){
			fmt.Println(line, " - Es id" )
		}
		if checkNum(line){
			fmt.Println(line, " - Es un numero")
		}
	}
}