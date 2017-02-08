/*
	LENGUAJES Y AUTOMATAS I
	HERNANDEZ CONTRERAS UZZIEL
	13211489
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func evaluarLinea(linea string) string {
	linea1 := strings.Replace(linea, " ", "", -1)

	res := "!--Expresión no válida--!"

	re := regexp.MustCompile(principal())
	if re.MatchString(linea1) {
		res = "programaPrincipal"
	}

	re = regexp.MustCompile(asignacionSimple())
	if re.MatchString(linea1) {
		res = "id asigna numeros"
	}

	re = regexp.MustCompile(asignacionOp())

	if re.MatchString(linea1) {
		res = "id asigna numeros opMat numeros"
	}

	re = regexp.MustCompile(asignacionIdID())

	if re.MatchString(linea1) {
		res = "id asigna id"
	}

	re = regexp.MustCompile(asignacionIdNum())

	if re.MatchString(linea1) {
		res = "id asigna id opMat numeros"
	}

	re = regexp.MustCompile(asignacionIdIdId())

	if re.MatchString(linea1) {
		res = "id asigna id opMat id"
	}

	re = regexp.MustCompile(asignacionNumId())
	if re.MatchString(linea1) {
		res = "id asigna numeros opMat id"
	}

	re = regexp.MustCompile(claveS())
	if re.MatchString(linea1) {
		res = "clave parDer evaluacion parIzq agDer"
	}

	re = regexp.MustCompile(agIzq())
	if re.MatchString(linea1) {
		res = "agIzq"
	}

	re = regexp.MustCompile(claveC())
	if re.MatchString(linea1) {
		res = "clave parDer evaluacion parIzq agDer bloqueCodigo agIzq"
	}

	re = regexp.MustCompile(saltoLinea())
	if re.MatchString(linea1) {
		res = ""
	}

	return res
}

func main() {
	f, _ := os.Open("codigoFuente.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea + "\t" + evaluarLinea(linea))
	}
}
