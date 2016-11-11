package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
	return Re.MatchString(email)
}

func validateURL(url string) bool {
	//Re := regexp.MustCompile(`^[www.|http://|https://]*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	Re := regexp.MustCompile(`^(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	//[A-Za-z0-9./?=_%+\-]* al final para youtube.com/WATCH_03
	return Re.MatchString(url)
}

func validateDigits(digits string) bool {
	Re := regexp.MustCompile(`^[0-9]+$`)
	return Re.MatchString(digits)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}

/*DE AQUI PARA ABAJO EMPEZAMOS CON LAS PALABRAS RESERVADAS.*/

func validateCiclos(abs string) bool {
	re := regexp.MustCompile(`(^goif$)|(^goelse$)|(^gofor$)|(^goforeach$)|(^gowhile$)|(^godo$)|(^goswitch$)+$`)
	return re.MatchString(abs)
}
func validateEnterosConSigno(enteros string) bool {
	re := regexp.MustCompile(`(^gosbyte$)|(^goint32$)|(^goint16$)|(^goint64)+$`)
	return re.MatchString(enteros)
}
func validateEnterosSinSigno(enteros string) bool {
	re := regexp.MustCompile(`(^gobyte$)|(^gouint32$)|(^gouint16$)|(^gouint64$)+$`)
	return re.MatchString(enteros)
}
func validateFlotSimple(flot string) bool {
	re := regexp.MustCompile(`(^gofloat16$)|(^gofloat32$)|(^gofloat64)+$`)
	return re.MatchString(flot)
}
func validateFlotDoble(doble string) bool {
	re := regexp.MustCompile(`(^godouble$)+$`)
	return re.MatchString(doble)
}
func validateUnicodeChar(char string) bool {
	re := regexp.MustCompile(`(^gochar$)|(^gochar16$)|(^gochar32$)|(^gochar64$)+$`)
	return re.MatchString(char)
}
func validateBoleanoLogic(bol string) bool {
	re := regexp.MustCompile(`(^gobool$)+$`)
	return re.MatchString(bol)
}
func validateObject(obj string) bool {
	re := regexp.MustCompile(`(^goobject$)+$`)
	return re.MatchString(obj)
}
func validateSecChar(sec string) bool {
	re := regexp.MustCompile(`(^gostring$)+$`)
	return re.MatchString(sec)
}
func validateDeci(dec string) bool {
	re := regexp.MustCompile(`(^godecimal$)+$`)
	return re.MatchString(dec)
}
func validateOperadores(op string) bool {
	re := regexp.MustCompile(`^[\[\]+\-*/.(){}=]+$`)
	return re.MatchString(op)
}
func validateModAcceso(mod string) bool {
	re := regexp.MustCompile(`(^gounion$)|(^gonative$)|(^gosigned$)|(^goregister$)|(^gounsigned$)|(^gonative$)|(^goabstract$)|(^goasync$)|(^goconst$)|(^goevent$)|(^goextern$)|(^goin$)|(^goout$)|(^gooverride$)|(^goreadonly$)|(^goseal$)|(^gostatic$)|(^govirtual$)|(^govolatile$)+$`)
	return re.MatchString(mod)
}
func validateInstDeAlto(inst string) bool {
	re := regexp.MustCompile(`(^gobreak$)|(^gocontinue$)|(^gogoto$)|(^goreturn$)+$`)
	return re.MatchString(inst)
}
func validateTipDeDtsGral(dat string) bool {
	re := regexp.MustCompile(`(^golong$)|(^goshort$)|(^gostruct$)|(^gouint$)|(^goulong$)|(^goushort$)+$`)
	return re.MatchString(dat)
}
func validateExecpciones(exc string) bool {
	re := regexp.MustCompile(`(^gothrows$)|(^gothrow$)|(^gotry$)|(^gocatch$)|(^gotry$)|(^gofinally$)|(^gotry$)|(^gocatch$)|(^goexcept$)|(^goraise$)|(^goassert$)+$`)
	return re.MatchString(exc)
}
func validateAccesibilidad(acc string) bool {
	re := regexp.MustCompile(`(^gointernal$)|(^goprivate$)|(^goprotected$)|(^gopublic$)|(^gointer$)+$`)
	return re.MatchString(acc)
}
func validateSpacName(spc string) bool {
	re := regexp.MustCompile(`(^gonamespace$)|(^gousing$)|(^goexternalias$)|(^goimport$)|(^gofrom$)|(^golambda$)+$`)
	return re.MatchString(spc)
}
func validateConver(cnv string) bool {
	re := regexp.MustCompile(`(^goexplicit$)|(^goimplicit$)|(^gooperator$)+$`)
	return re.MatchString(cnv)
}

func validateParM(prm string) bool {
	re := regexp.MustCompile(`(^goparams$)|(^goref$)|(^goout$)|(^gorefout$)|(^gosuper$)|(^gosynch$)+$`)
	return re.MatchString(prm)
}
func validateClavAc(clva string) bool {
	re := regexp.MustCompile(`(^gobase$)|(^gothis$)|(^gothese$)|(^gothem$)|(^gothey$)+$`)
	return re.MatchString(clva)
}
func validateClavOP(clvo string) bool {
	re := regexp.MustCompile(`(^goas$)|(^goawait$)|(^gois$)|(^gonew$)|(^gosizeof$)|(^gotypeof$)|(^gotrue$)|(^gofalse$)|(^gostackalloc$)|(^gonameof)|(^goinstanceof$)+$`)
	return re.MatchString(clvo)
}
func validateTipRef(tpf string) bool {
	re := regexp.MustCompile(`(^govoid$)|(^govar$)+$`)
	return re.MatchString(tpf)
}
func validateClvIns(clvin string) bool {
	re := regexp.MustCompile(`(^gofixed$)|(^golock$)|(^gointerface$)+$`)
	return re.MatchString(clvin)
}
func validateTipRef2(clvio string) bool {
	re := regexp.MustCompile(`(^goclass$)|(^godelegate$)|(^godynamic$)|(^gointerface$)+$`)
	return re.MatchString(clvio)
}
func validateClvLite(lite string) bool {
	re := regexp.MustCompile(`(^gonull$)|(^godefault$)+$`)
	return re.MatchString(lite)
}
func validateClvConx(convx string) bool {
	re := regexp.MustCompile(`(^gopackage$)|(^goadd$)|(^godynamic$)|(^goglobal$)|(^gojoin$)|(^gopartial$)|(^goselect$)|(^govar$)|(^goyield$)|(^golet$)|(^goremove$)|(^gononlocal$)|(^goprint$)|(^goprintline$) |(^goprintint$)|(^goprintdouble$)|(^goprintbool$)|(^goprintchar$)|(^goprintstring$)|(^goimplements$)+$`)
	return re.MatchString(convx)
}
func validateOpLog(logx string) bool {
	re := regexp.MustCompile(`(^goall$)|(^goany$)|(^gobetween$)|(^goexists$)|(^golike$)|(^gosome$)|(^goand$)|(^goor$)|(^gonot$)|(^goxand$)|(^goxor$)|(^goxnot$)+$`)
	return re.MatchString(logx)
}
func validateMath(math string) bool {
	re := regexp.MustCompile(`(^gotan$)|(^gosin$)|(^gocos$)|(^gopi$)+$`)
	return re.MatchString(math)
}
func validateOprel(rel string) bool {
	re := regexp.MustCompile(`^(>|<|>=|<=|!=|==)+$`)
	return re.MatchString(rel)
}
func validateCMDctrl(cctrl string) bool {
	re := regexp.MustCompile(`(^gocheckpoint$)|(^gokill$)|(^goreconfigure$)|(^goshutdown$)|(^gokillstatsjob$)+$`)
	return re.MatchString(cctrl)
}
func validateSecStat(stst string) bool {
	re := regexp.MustCompile(`(^goaddsignature$)|(^goclosemasterkey$)|(^godeny$)|(^goexecuteas$)|(^gogrant$)|(^goopenmasterkey$)|(^gorevert$)|(^gorevoke$)|(^gosetuser$)|(^goazure$)+$`)
	return re.MatchString(stst)
}
func validateFuncOBD(obd string) bool {
	re := regexp.MustCompile(`(^goabsolute$)|(^goaction$)|(^goada$)|(^goallocate$)|(^goalter$)|(^goare$)|(^goasc$)|(^goat$)|(^goavg$)|(^gobegin$)|(^gobit$)|(^goforeign$)|(^gofortran$)|(^gofound$)|(^gofull$)|(^gogrant$)+$`)
	return re.MatchString(obd)
}
func validateID(id string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateSPACE(id string) bool {
	Re := regexp.MustCompile(`^[ ]+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {
	/*ABRIMOS NUESTRO ARCHIVO DE TEXTO*/
	c := color.New(color.FgHiRed).Add(color.Underline)
	d := color.New(color.FgHiYellow).Add(color.Underline)
	e := color.New(color.FgHiGreen).Add(color.Underline)
	g := color.New(color.FgHiWhite).Add(color.Underline)
	contador := 1
	/*INICIAMOS EL MENU DE COMANDOS*/
	var menu string
	d.Println("\n\n                        Instituto Tecnologico De Tijuana.\n")
	fmt.Println(`
                   ██╗     ███████╗██╗  ██╗██╗ ██████╗ ██████╗
                   ██║     ██╔════╝╚██╗██╔╝██║██╔════╝██╔═══██╗
                   ██║     █████╗   ╚███╔╝ ██║██║     ██║   ██║
                   ██║     ██╔══╝   ██╔██╗ ██║██║     ██║   ██║
                   ███████╗███████╗██╔╝ ██╗██║╚██████╗╚██████╔╝
                   ╚══════╝╚══════╝╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═════╝

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

	f, _ := ioutil.ReadFile("C://Download//Pokemon//pokemon.txt")
	s := string(f[:])
	r := strings.NewReplacer("{"," { ", "}"," } ", "("," ( ",")"," ) ","+"," + ","="," = ","."," . ","["," ] ","-"," - ","*"," * ","/"," / ","++"," ++ ","--"," -- ","=="," == ","<"," < ",">"," > ",">="," >= ","<="," <= ")
	z := r.Replace(s)

	line:=strings.Fields(z)
	if menu == "in" {
		for i := range line {

			color.Set(color.FgHiWhite).Add(color.Underline)
			/*ER tokens ya con sus palabras*/
			if validateCiclos(line[i]) {
				fmt.Println(contador, ("- Es un tipo ciclo: "), line[i])
				contador++
			} else if validateEnterosConSigno(line[i]) {
				fmt.Println(contador, "- Es un entero con signo: ", line[i])
				contador++
			} else if validateEnterosSinSigno(line[i]) {
				fmt.Println(contador, "- Es un entero sin signo: ", line[i])
				contador++
			} else if validateFlotSimple(line[i]) {
				fmt.Println(contador, "- Es un tipo flotante simple: ", line[i])
				contador++
			} else if validateFlotDoble(line[i]) {
				fmt.Println(contador, "- Es un tipo flotante double: ", line[i])
				contador++
			} else if validateUnicodeChar(line[i]) {
				fmt.Println(contador, "- Es un tipo unicode(char): ", line[i])
				contador++
			} else if validateBoleanoLogic(line[i]) {
				fmt.Println(contador, "- Es un tipo bolean logico: ", line[i])
				contador++
			} else if validateObject(line[i]) {
				fmt.Println(contador, "- Tipo base de todos los otros tipos: ", line[i])
				contador++
			} else if validateSecChar(line[i]) {
				fmt.Println(contador, "- Una secuencia de caracteres(string): ", line[i])
				contador++
			} else if validateDeci(line[i]) {
				fmt.Println(contador, "- Tipo preciso fraccionario o integral: ", line[i])
				contador++
			} else if validateOperadores(line[i]) {
				fmt.Println(contador, "- Son operadores: ", line[i])
				contador++
			} else if validateModAcceso(line[i]) {
				fmt.Println(contador, "- Son modificadores de acceso: ", line[i])
				contador++
			} else if validateInstDeAlto(line[i]) {
				fmt.Println(contador, "- Instruccones de alto: ", line[i])
				contador++
			} else if validateTipDeDtsGral(line[i]) {
				fmt.Println(contador, "- Tipo de dato general: ", line[i])
				contador++
			} else if validateExecpciones(line[i]) {
				fmt.Println(contador, "- Instrucciones para el control de excepciones: ", line[i])
				contador++
			} else if validateAccesibilidad(line[i]) {
				fmt.Println(contador, "- Instrucciones de accesibilidad: ", line[i])
				contador++
			} else if validateSpacName(line[i]) {
				fmt.Println(contador, "- Palabras clave del espacio de nombres: ", line[i])
				contador++
			} else if validateConver(line[i]) {
				fmt.Println(contador, "- Palabras claves para conversiones: ", line[i])
				contador++
			} else if validateParM(line[i]) {
				fmt.Println(contador, "- Parámetros de métodos: ", line[i])
				contador++
			} else if validateClavAc(line[i]) {
				fmt.Println(contador, "- Palabras clave de acceso: ", line[i])
				contador++
			} else if validateClavOP(line[i]) {
				fmt.Println(contador, "- Palabras clave de operador: ", line[i])
				contador++
			} else if validateTipRef(line[i]) {
				fmt.Println(contador, "- Tipos de referencia: ", line[i])
				contador++
			} else if validateClvIns(line[i]) {
				fmt.Println(contador, "- Palabras clave de instrucciones: ", line[i])
				contador++
			} else if validateTipRef2(line[i]) {
				fmt.Println(contador, "- Tipos de referencia especificos: ", line[i])
				contador++
			} else if validateClvLite(line[i]) {
				fmt.Println(contador, "- Palabras clave de literales: ", line[i])
				contador++
			} else if validateClvConx(line[i]) {
				fmt.Println(contador, "- Palabras contextuales: ", line[i])
				contador++
			} else if validateEmail(line[i]) {
				fmt.Println("Es un Email: ", line[i])
			} else if validateURL(line[i]) {
				fmt.Println("Es un URL: ", line[i])
			} else if validateFuncOBD(line[i]) {
				fmt.Println(contador, "- Es un funcion ODBC: ", line[i])
				contador++
			} else if validateDigits(line[i]) {
				fmt.Println(contador, "- Es un Numero: ", line[i])
				contador++
			} else if validateOpLog(line[i]) {
				fmt.Println(contador, "- Son operadores logicos: ", line[i])
				contador++
			} else if validateMath(line[i]) {
				fmt.Println(contador, "- Son operadores matematicos (constantes) : ", line[i])
				contador++
			} else if validateOprel(line[i]) {
				fmt.Println(contador, "- Son operadores relacionales: ", line[i])
				contador++
			} else if validateCMDctrl(line[i]) {
				fmt.Println(contador, "- Son comandos de manejo: ", line[i])
				contador++
			} else if validateSecStat(line[i]) {
				fmt.Println(contador, "- Son declaraciones de seguridad: ", line[i])
				contador++
				//	contador++
				//} else if validateFuncOBD( line[i] ) {
			} else if validateID(line[i]) {
				fmt.Println(contador, "- Es una variable: ", line[i])
				contador++
			} else if validateSPACE(line[i]) {
				fmt.Println("Es un espacio: ", line[i])
				contador++
			} else {
				c.Println("Error invalid entry: (", line[i], ") on  line : ", contador)

			}
			color.Unset()
		}

	} else if menu == "man" {
		/*si es posible abrir el doc de word directo*/
		b, _ := ioutil.ReadFile("C://Download//Pokemon//menu.txt") // just pass the file name
		g.Println(string(b))

	} else if menu == "quit" {
		g.Print("\n                         THANKS FOR USING THE PROGRAM.\n")
		os.Exit(0)
	} else {
		c.Println("\n                         INVALID ENTRY,PLEASE RE-TRY:", menu)
	}
	fmt.Println("\n")
	contador = 1
	goto L
}
