package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	//"time"
	"time"
)

var menu string
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func validateCreateBD(bd string) bool {
	Re := regexp.MustCompile(`(^mccreate$)`)
	return Re.MatchString(bd)
}
func validatedatabse(bds string) bool {
	Re := regexp.MustCompile(`(^database$)`)
	return Re.MatchString(bds)
}
func validateUseBD(usebd string) bool {
	Re := regexp.MustCompile(`(^mcuse$)`)
	return Re.MatchString(usebd)
}
func validateCreateTable(crttbl string) bool {
	Re := regexp.MustCompile(`(^mccreatetable$)`)
	return Re.MatchString(crttbl)
}
func validateOperadores(op string) bool {
	re := regexp.MustCompile(`^[\[\]+\-*/.(){}=]+$`)
	return re.MatchString(op)
}
func validateDigits(digits string) bool {
	Re := regexp.MustCompile(`^[0-9]+$`)
	return Re.MatchString(digits)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateID(id string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z]+[0-9]+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validatedeci(deci string) bool {
	Re := regexp.MustCompile(`^[0-9]\.[0-9]+$`)
	return Re.MatchString(deci)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validatestring(str string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z]`)
	return Re.MatchString(str)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {

	contador := 1
	c := color.New(color.FgHiRed).Add(color.Underline)
	//d := color.New(color.FgHiYellow).Add(color.Underline)
	e := color.New(color.FgHiGreen).Add(color.Underline)
	g := color.New(color.FgHiWhite).Add(color.Underline)
	g.Println("\n\n                        Instituto Tecnologico De Tijuana.\n")
	g.Println(`Materia: Administracion de base de datos.

	Creators:

	 Lopez Valenzuela Marco Antonio 13211502.
	 Sandoval Lizarraga Christopher Francisco 13211481.

	`)
L:
	g.Println("WELCOME to the commands menu.\n")
	e.Print("               in   =    ")
	g.Println("it will run the program.")
	e.Print("               man  =    ")
	g.Println("it will generate a manual for the tokens.")
	e.Print("               quit =    ")
	g.Println("it will close the program.\n")
	g.Print("Enter a valid command: ")
	fmt.Scan(&menu)

	f, _ := ioutil.ReadFile("C://Download//Pokemon//bonifaz.txt")
	s := string(f[:])
	//r := strings.NewReplacer("{", " { ", "}", " } ", "(", " ( ", ")", " ) ", "+", " + ", "=", " = ", ".", " . ", "[", " ] ", "-", " - ", "*", " * ", "/", " / ", "++", " ++ ", "--", " -- ", "==", " == ", "<", " < ", ">", " > ", ">=", " >= ", "<=", " <= ")
	//z := r.Replace(s)
	//line:=strings.Fields(z)
	line := strings.Fields(s)
	if menu == "in" {
		fmt.Println("\n")
		//fmt.Println("\n")
		fmt.Println("It will run the program in 3 seconds...")
		time.Sleep(3 * time.Second)
		fmt.Println("\n")
		for i := range line {

			color.Set(color.FgHiWhite).Add(color.Underline)

			if validateCreateBD(line[i]) {
				fmt.Println(contador, "- Creating database...", line[i])
				contador++
			} else if validatedatabse(line[i]) {
				fmt.Println(contador, "- Creating database...", line[i])
				contador++
			} else if validateUseBD(line[i]) {
				fmt.Println(contador, "- Using database...", line[i])
				contador++
			} else if validateCreateTable(line[i]) {
				fmt.Println(contador, "- Creating table...", line[i])
				contador++
			} else if validateDigits(line[i]) {
				fmt.Println(contador, "- Es un digito: ", line[i])
				contador++
			} else if validateOperadores(line[i]) {
				fmt.Println(contador, "- Es un operador/agrupador: ", line[i])
				contador++
			} else if validatedeci(line[i]) {
				fmt.Println(contador, "- Es un numero tipo decimal: ", line[i])
				contador++
			} else if validateID(line[i]) {
				fmt.Println(contador, "- Es un tipo alphanumeric: ", line[i])
				contador++
			} else if validatestring(line[i]) {
				fmt.Println(contador, "- Es un numero tipo cadena: ", line[i])
				contador++

			} else {
				c.Println("Error invalid entry: (", line[i], ") on  line : ", contador)
				contador++
			}
			color.Unset()
		}
	} else if menu == "man" {
		/*si es posible abrir el doc de word directo*/
		fmt.Println("\n")
		fmt.Println("It will bring the manual in 5 seconds...")
		time.Sleep(5 * time.Second)
		fmt.Println("\n")
		b, _ := ioutil.ReadFile("C://Download//Pokemon//manualbonifaz.txt") // just pass the file name
		g.Println(string(b))

		//} else if menu == "clear" {
		//	CallClear()

	} else if menu == "quit" {
		g.Print("\n                         THANKS FOR USING THE PROGRAM.\n")
		os.Exit(0)
	} else {
		c.Println("\n                         INVALID ENTRY,PLEASE RE-TRY:", menu)
	}
	//main()
	fmt.Println("\n")
	contador = 1
	goto L
}
