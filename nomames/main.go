package main

import (
	"bufio" //Permite leer entre espacios
	"fmt"   //Instrucciones basicas
	"github.com/fatih/color"
	"io/ioutil"
	"os" //Permite leer entre espacios
	"path/filepath"
	"regexp" //Permite las expresiones regulares
	"strings"
	"time"
)

var username string
var password string
var mainlog string
var login string

func main() {
	k := color.New(color.FgRed).Add(color.Underline)
	c := color.New(color.FgHiRed).Add(color.Underline)
	e := color.New(color.FgHiGreen).Add(color.Underline)
	g := color.New(color.FgHiWhite).Add(color.Underline)
	b := color.New(color.FgBlue).Add(color.Underline)

	g.Println("\n\n                       Instituto Tecnologico De Tijuana.\n")
	b.Println(`
                                  \__    __/
                                  /_/ /\ \_\
                                 __ \ \/ / __
                                 \_\_\/\/_/_/
                             __/\___\_\/_/___/\__
                               \/ __/_/\_\__ \/
                                 /_/ /\/\ \_\
                                  __/ /\ \__
                                  \_\ \/ /_/
                                  /        \
	`)
	g.Println(`Materia: Administracion de base de datos.

	 Creators:

	 Lopez Valenzuela Marco Antonio 13211502.
	 Sandoval Lizarraga Christopher Francisco 13211481.

	`)
B:
	g.Println("\nWould you like to sign in?\n")
	g.Print("[Y/N]: ")
	fmt.Scan(&mainlog)
	if mainlog == "Y" || mainlog == "y" {
	J:
		g.Print("\nEnter username: ")
		fmt.Scan(&username)
		g.Print("\nEnter password: ")
		fmt.Scan(&password)

		if username == "Marco" || username == "marco" && password == "123" || username == "Cristopher" || username == "cristopher" && password == "123" {
			k.Print("\nWelcomeback: ", strings.ToUpper(username), "...")
			e.Print("\n\ninitializing please wait...\n")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println()
			c.Println("Incorrect credentials, would you like to try again?.\n")
		A:
			g.Print("[Y/N]: ")
			fmt.Scan(&login)
			if login == "Y" || login == "y" {
				goto J
			} else if login == "N" || login == "n" {
				g.Println("\nClosing program...")
				time.Sleep(2 * time.Second)
				os.Exit(0)
			} else {
				g.Println("\nIncorrect command: ", login)
				goto A
			}

		}
		g.Println("\nEnter command <gomanual> for a display of the basic commands.\n")

		//DIRECTORIO

		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
	L:
		k.Print("[GoSQL]-> ")

		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()

		if funCrearBD(scan.Text()) {
			scann := strings.Fields(scan.Text())
			g.Println("\nCreating database, please wait...")
			time.Sleep(2 * time.Second)
			//e.Println("\nDatabase sucessfuly CREATED.")
			//os.Mkdir("C:\\ProyectoBonifaz"+string(filepath.Separator)+scann[2], 077)
			err, _ := os.Stat("C:\\ProyectoBonifaz\\" + scann[2])
			if err != nil {
				b.Println("\nDatabase already exist:", scann[2])
				println()
				goto L
			} else {
				e.Println("\nDatabase sucessfuly CREATED.\n")
				os.Mkdir("C:\\ProyectoBonifaz"+string(filepath.Separator)+scann[2], 077)
				goto L
			}

		} else if funBorrarBD(scan.Text()) {
			var seguro string
			scann := strings.Fields(scan.Text())
			g.Println("\nAre you sure you want to do delete database: ", scann[2], "?")
			g.Print("\n[Y/N]: ")
			fmt.Scan(&seguro)
			if seguro == "Y" || seguro == "y" {
				g.Println("\nDeleting database, please wait...")
				time.Sleep(2 * time.Second)
				e.Println("\nDatabase sucessfuly DELETED.")
				fmt.Println()
				os.Remove("C:\\ProyectoBonifaz\\" + scann[2])
				goto L

			} else if seguro == "N" || seguro == "n" {
				goto L

			} else {
				fmt.Println("\nInvalid command", seguro)
			}

		} else if funUSE(scan.Text()){
			scann := strings.Fields(scan.Text())
			os.Chdir("C:\\ProyectoBonifaz" + scann[2])

		} else if funCrearTabla(scan.Text()) {
			//para crear tablas o docs de texto.

		} else if funSignOUT(scan.Text()) {
			goto B

		} else if funBorrarTBALL(scan.Text()) {

			//NO SALE COMPLETO.
			os.Remove("C:\\ProyectoBonifaz")

		} else if funSHOW(scan.Text()) {
			files, _ := ioutil.ReadDir("C:\\ProyectoBonifaz\\")
			fmt.Println()
			fmt.Println(`	+---------------+`)
			fmt.Println(`        |   DATABASES   |`)
			fmt.Println(`	+---------------+`)
			for _, f := range files {
				fmt.Println()
				fmt.Println(`	|`, f.Name(), `	|`)
				fmt.Println(`	-----------------`)
			}
			fmt.Println(`	+---------------+`)

		} else if scan.Text() == "gomanual" {
			g.Println(`
	Gocreate database <name>: command allows us to create a new database.
	Godelete database <name>: command allows us to delete a database.
	Gocreate table <name>: command will create a table.
	Godelete table <name>: command will delete a table.
	Gosignout: will sign in with a different username.
	Goshow databases: will show the databases that already exist.
	Goexit: will exit the program. `)

		} else if funBorrarTB(scan.Text()) {
			//para borrar tablas o docs de texto
		} else if scan.Text() == "" {

		} else if scan.Text() == "goexit" {
			fmt.Println("\nClosing program...")
			time.Sleep(2 * time.Second)
			os.Exit(0)
		} else {
			c.Println("\nError, command was unsucessful: ", scan.Text())
		}
		fmt.Println("\n")
		goto L

	} else {
		g.Println("\nClosing program...")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}
}
func funCrearBD(name string) bool {
	crearBD, _ := regexp.MatchString("^(gocreate database [a-zA-Z0-9]+)$", name)
	return crearBD
}
func funBorrarBD(name string) bool {
	borrarBD, _ := regexp.MatchString("^(godelete database [a-zA-Z0-9]+)$", name)
	return borrarBD
}
func funCrearTabla(name string) bool {
	crearTB, _ := regexp.MatchString("^(gocreate table [a-zA-Z0-9]+)$", name)
	return crearTB
}
func funBorrarTB(name string) bool {
	borrarTB, _ := regexp.MatchString("^(godelete table [a-zA-Z0-9]+)$", name)
	return borrarTB
}
func funBorrarTBALL(name string) bool {
	borrarTB, _ := regexp.MatchString("^(godelete all databases+)$", name)
	return borrarTB
}
func funSignOUT(name string) bool {
	signout, _ := regexp.MatchString("^(gosignout+)$", name)
	return signout
}
func funSHOW(name string) bool {
	show, _ := regexp.MatchString("^(goshow databases+)$", name)
	return show
}
func funUSE(name string) bool {
	use, _ := regexp.MatchString("^(gouse database [a-zA-Z0-9]+)$", name)
	return use
}
func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
