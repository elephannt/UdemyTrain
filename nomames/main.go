package main

import (
	"bufio" //Permite leer entre espacios
	"fmt"   //Instrucciones basicas
	"github.com/fatih/color"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"log"
	"os" //Permite leer entre espacios
	"path/filepath"
	"regexp" //Permite las expresiones regulares
	"strings"
	"time"

	"os/exec"
	"bytes"
)

var username string
var password string
var mainlog string
var login string
var seguro string

var nombreBD string
var nombreTL string
var nombreDT string
var nombreAT string

var pathcambiando string = "C:\\MainDB\\"
var path string = "C:\\MainDB\\"

func main() {

	var (
		newFile *os.File
		err     error
	)

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
	if mainlog == "N" || mainlog == "n" {
		g.Println("\nClosing program...")
		time.Sleep(2 * time.Second)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		os.Exit(0)
	}
	if mainlog == "Y" || mainlog == "y" {

	J:
		//PARA LIMPIAR LAPANTALLA
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		c.Println(`
   +-----------------------------------+
   |   ____        ____    ___   _     |
   |  / ___|  ___ / ___|  / _ \ | |    |
   | | |  _  / _ \\___ \ | | | || |    |
   | | |_| || (_) |___) || |_| || |___ |
   |  \____| \___/|____/  \__\_\|_____||
   |                                   |
   +-----------------------------------+
		`)
		g.Print(`          Enter username: `)
		fmt.Scan(&username)

		// PONER ESTO QUITA EL ERROR REPETIDO, SIN ESTO NO FUNCIONAEL PASSWORD, CHECAR TESTTING... WTF.

		silentPassword, _ := gopass.GetPasswd() // Silent
		fmt.Print(string(silentPassword))
		g.Print(`          Enter password: `)
		maskedPassword, _ := gopass.GetPasswdMasked() // Masked

		if username == "Marco" || username == "marco" && string(maskedPassword) == "1" || username == "Cristopher" || username == "cristopher" && string(maskedPassword) == "123456789" {
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
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
				os.Exit(0)
			} else {
				g.Println("\nIncorrect command: ", login)
				goto A
			}

		}

E:
		g.Println("\nEnter command <gomanual> for a display of the basic commands.\n")

		//DIRECTORIO

		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
	L:
		k.Print(pathcambiando, " | GoSQL]-> ")

		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()

		if funCrearBD(scan.Text()) {
			scann := strings.Fields(scan.Text())
			g.Println("\nCreating database, please wait...")
			time.Sleep(2 * time.Second)
			//e.Println("\nDatabase sucessfuly CREATED.")
			//os.Mkdir("C:\\MainDB"+string(filepath.Separator)+scann[2], 077)
			err, _ := os.Stat("C:\\MainDB\\" + scann[2])
			if err != nil {
				c.Println("\nDatabase already exist:", scann[2])
				println()
				goto L
			} else {
				e.Println("\nDatabase sucessfuly CREATED.\n")
				os.Mkdir("C:\\MainDB"+string(filepath.Separator)+scann[2], 077)
				goto L
			}

		} else if funBorrarBD(scan.Text()) {
			scann := strings.Fields(scan.Text())
			g.Println("\nAre you sure you want to do delete database: ", scann[2], "?")
			g.Print("\n[Y/N]: ")
			fmt.Scan(&seguro)
			if seguro == "Y" || seguro == "y" {
				g.Println("\nDeleting database, please wait...")
				time.Sleep(2 * time.Second)
				e.Println("\nDatabase sucessfuly DELETED.")
				fmt.Println()
				os.RemoveAll("C:\\MainDB\\" + scann[2])
				goto L

			} else if seguro == "N" || seguro == "n" {
				goto L

			} else {
				fmt.Println("\nInvalid command", seguro)
			}

		} else if funGOBACK(scan.Text()) {
			_, err := os.Stat(pathcambiando)
			if err == nil {
				os.Chdir(path)
				g.Println("\nReturning to main database...")
				time.Sleep(2 * time.Second)
				pathcambiando = path

			}
			if os.IsNotExist(err) {
				g.Println("\nDatabase doesnt exist...")
			}

		} else if funUSE(scan.Text()) {
			scann := strings.Fields(scan.Text())
			nombreBD = scann[2]

			_, err := os.Stat(pathcambiando + nombreBD)
			if err == nil {
				os.Chdir(pathcambiando + nombreBD)
				g.Println("\nUsing database base please wait...", nombreBD)
				time.Sleep(2 * time.Second)
				pathcambiando = pathcambiando + nombreBD
			}
			if os.IsNotExist(err) {
				g.Println("\nDatabase doesnt exist...", nombreBD)
			}

		} else if funUSETL(scan.Text()){


		} else if funCreateFile(scan.Text()) {

			scann := strings.Fields(scan.Text())
			nombreTL = scann[2]
			newFile, err = os.Create(path + nombreBD + "\\" + nombreTL + ".txt")
			if err != nil {
				g.Println("\nTable already exists...", nombreTL)
			}
			g.Println("\nCreating new table, please be patient...")
			time.Sleep(2 * time.Second)
			nombreTL =""
			fmt.Println()
			log.Println()
			newFile.Close()

			//err, _ := os.Stat("C:\\MainDB\\" + scann[2])
			//if err != nil {
			//	c.Println("\nDatabase already exist:", scann[2])
			//	println()
			//	goto L
			//} else {
			//	e.Println("\nDatabase sucessfuly CREATED.\n")
			//	os.Mkdir("C:\\MainDB"+string(filepath.Separator)+scann[2], 077)
			//	goto L
			//}

		} else if funCrearTabla(scan.Text()) {
			//para crear tablas o docs de texto.

		} else if funSignOUT(scan.Text()) {
			goto B

		} else if funBorrarTBALL(scan.Text()) {

			//NO SALE COMPLETO.
			os.Remove("C:\\MainDB")

		} else if funSHOW(scan.Text()) {
			files, _ := ioutil.ReadDir("C:\\MainDB\\")
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

		}else if funshowtl(scan.Text()){
			g.Println("\nRetriving, please be patient...")
			time.Sleep(2 * time.Second)
			files, _ := ioutil.ReadDir(path+nombreBD+nombreTL)
			fmt.Println()
			fmt.Println(`	+---------------+`)
			fmt.Println(`        |   TABLES      |`)
			fmt.Println(`	+---------------+`)
			for _, f := range files {
				fmt.Println()
				fmt.Println(`	|`, f.Name(), `	|`)
				fmt.Println(`	-----------------`)
			}
			fmt.Println(`	+---------------+`)
			//pathcambiando=path

		} else if scan.Text() == "gomanual" {
			g.Println(`
	Gocreate database <name>: command allows us to create a new database.
	Godelete database <name>: command allows us to delete a database.
	Gocreate table <name>: command will create a table.
	Godrop table <name>: command will delete a table.
	Gosignout: will sign in with a different username.
	Goshow databases: will show the databases that already exist.
	Gouse database <name>: will use the database selected.
	Goback: will return to the main database.
	Goexit: will exit the program.
	Goclear screen: will clear the current screen.
	Goinsert <name> <data>:will insert data into desired table.
	Goread table <name>: will display the atribbutes.
        Godropattribute <name> <data>:will delete data from table.
	`)

		} else if funInsert(scan.Text()){
			scanns := strings.Fields(scan.Text())
			nombreTL = scanns [1]

			fileHandle, _ := os.OpenFile(path + nombreBD + "\\" + nombreTL+".txt", os.O_APPEND, 0666)
			writer := bufio.NewWriter(fileHandle)
			defer fileHandle.Close()

			//Nomas para ver donde estoy.

			//fmt.Println(path)
			//fmt.Println(nombreBD)
			//fmt.Println(nombreTL)

			scann := strings.Fields(scan.Text())
			g.Println("\nInserting data on table, please wait...")
			time.Sleep(2 * time.Second)
			e.Println("\nData was sucessfuly added.")

			nombreDT = scann [2]
			fmt.Fprintln(writer, nombreDT)
			writer.Flush()
			nombreTL=""

		} else if funreadtable(scan.Text()){

			scanns := strings.Fields(scan.Text())
			nombreTL = scanns [2]
			g.Println("\nRetriving data, please wait...")
			time.Sleep(2 * time.Second)
			data, err := ioutil.ReadFile(path + nombreBD + "\\" + nombreTL+".txt")
			if err != nil {
				g.Println("\nError, table doesnt exist: ",nombreTL)
				fmt.Println()
				goto L
			}
			s := string(data[:])
			line:=strings.Split(s,",")
			fmt.Println()
			fmt.Println(`	+-----------------+`)
			fmt.Println(`        |    ATTRIBUTES   |`)
			fmt.Println(`	+-----------------+`)
			for i := range line {
				fmt.Println()
				g.Println(`	|`,line[i],`	          `)
				fmt.Println(`	-------------------`)
			}
			fmt.Println(`	+-----------------+`)
                        nombreTL=""

		} else if funborrartl(scan.Text()) {
			scann := strings.Fields(scan.Text())
			nombreTL = scann[2]
			g.Println("\nRemoving table, please wait...")
			time.Sleep(2 * time.Second)
			os.Remove(path + nombreBD + "\\" +nombreTL + ".txt")
			e.Println("\nTable was sucessfuly removed.")



		} else if scan.Text() == "" {

		} else if funborrarattr(scan.Text()){

			scann := strings.Fields(scan.Text())
			nombreTL = scann[1]

			input, err := ioutil.ReadFile(path + nombreBD + "\\" + nombreTL+".txt")
			if err != nil {
				g.Println("\nError 1",err)

			}
			g.Println("\nDeleting data, please wait...")
			time.Sleep(2 * time.Second)
			nombreAT = scann [2]
			g.Println("\nData sucessfuly deleted.")
			output := bytes.Replace(input, []byte(nombreAT), []byte(""), -1)
nombreTL=""
			if err = ioutil.WriteFile(path + nombreBD + "\\" + nombreTL+".txt", output, 0666); err != nil {
				g.Println("\nError 1",err)
			}

		} else if funClearScreen(scan.Text()){
			g.Println("\nClearing the screen...")
			time.Sleep(2 * time.Second)
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			goto E

		} else if scan.Text() == "goexit" {
			g.Println("\nClosing program...")
			time.Sleep(2 * time.Second)
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
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
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
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
func funGOBACK(name string) bool {
	goback, _ := regexp.MatchString("^(goback+)$", name)
	return goback
}
func funCrearTabla(name string) bool {
	crearTB, _ := regexp.MatchString("^(gocreate table [a-zA-Z0-9]+)$", name)
	return crearTB
}
func funBorrarTB(name string) bool {
	borrarTB, _ := regexp.MatchString("^(godelete lksdlksl [a-zA-Z0-9]+)$", name)
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
func funUSETL(name string) bool {
	usetl, _ := regexp.MatchString("^(gouse table [a-zA-Z0-9]+)$", name)
	return usetl
}
func funCreateFile(name string) bool {
	file, _ := regexp.MatchString("^(gocreate table [a-zA-Z0-9]+)$", name)
	return file
}
func funClearScreen(name string) bool {
	screen, _ := regexp.MatchString("^(goclear screen+)$", name)
	return screen
}
func funInsert(name string) bool {
	insert, _ := regexp.MatchString("^(goinsert [a-zA-Z0-9]+ [a-zA-Z0-9,)(=\\s]+)$", name)
	return insert
}
func funshowtl(name string) bool {
	showt, _ := regexp.MatchString("^(goshow tables [a-zA-Z0-9]+)$", name)
	return showt
}
func funreadtable(name string) bool {
	readt, _ := regexp.MatchString("^(goread table [a-zA-Z0-9]+)$", name)
	return readt
}
func funborrartl(name string) bool {
	borrartl, _ := regexp.MatchString("^(godrop table [a-zA-Z0-9]+)$", name)
	return borrartl
}
func funborrarattr(name string) bool {
	attr, _ := regexp.MatchString("^(godropattribute [a-zA-Z0-9]+ [a-zA-Z0-9,)(=\\s]+)$", name)
	return attr
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
