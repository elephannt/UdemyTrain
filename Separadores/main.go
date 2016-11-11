package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	//"strings"
	//"unicode"
	"strings"
	//"unicode"
	//"unicode"
	"unicode"
)

func validateciclo(id string) bool {
	Re := regexp.MustCompile(`^(goif)+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateDigits(digits string) bool {
	Re := regexp.MustCompile(`^[0-9]+$`)
	return Re.MatchString(digits)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateID(id string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z]+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateSPACE(id string) bool {
	Re := regexp.MustCompile(`^[ ]*$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func validateagru(id string) bool {
	Re := regexp.MustCompile(`^[+()-=*><]+$`)
	return Re.MatchString(id)
	//Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
}
func main() {
	f, _ := ioutil.ReadFile("C://Download//Pokemon//ejemplo.txt")
	s := string(f[:])
	// MARCO TODOS LOS SIGNOS SIN LETRAS  CON O SIN ESPACIO

	// valores[i] := strings.FieldsFunc(s, func (r rune) bool {
	//	return !unicode.IsSymbol(r) && !unicode.IsPunct(r)
	//})

	//MARCA TODAS LAS LETRAS SIN SIGNO CON O SIN ESPACIO.

	valores := strings.FieldsFunc(s, func (r rune) bool {
		return !unicode.IsLetter(r)
	})
//valores := strings.Replace(s,""," ",-1)
	//AGRUPADORES
	//func IsOneOf(ranges []*RangeTable, r rune) bool
	//valores := strings.FieldsFunc(s, func(r rune) bool {
	//	switch r {
	//	case ' ':
	//		return true
	//	}
	//	return false
	//})

	for i := range  valores {
		if validateciclo( valores[i]) {
			fmt.Println("Es un ciclo: ",  valores[i])
		} else if validateagru( valores[i] ) {
			fmt.Println("Es un agrupador: ",  valores[i] )
		} else if validateDigits( valores[i] ){
			fmt.Println("Es un digito",  valores[i] )
		} else if validateID( valores[i] ){
			fmt.Println("Es un id: ",  valores[i] )
		} else if validateSPACE( valores[i] ) {
			fmt.Println("es un espacio: ",  valores[i] )
		} else {
			fmt.Println("error",  valores[i] )
		}

	}
}
