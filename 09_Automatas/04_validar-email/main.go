package main

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
	return Re.MatchString(email)
}
func main() {
	email := "abc@abc12"
	if !validateEmail(email) {
		fmt.Println("Email address is invalid")
	} else {
		fmt.Println("Email adress is VALID")
	}
	email = "Ab/c@abc123.com"
	if !validateEmail(email) {
		fmt.Println("Email adress is invalid")
	} else {
		fmt.Println("Email is VALID")
	}
}
//^ Inicio de la cadena.

//A-Za-z0-9 nos dice que puede incluir cualquier letra y cualquier numero

//._% puede tener esos caracteres.

/*+ causes the resulting RE to match 1 or more repetitions of the preceding RE.
 ab+ will match ‘a’ followed by any non-zero number of ‘b’s; it will not match just ‘a’.*/

// \- diciendo que puede aceptar guiones.

// el {2,4} solo limita el numero de caracteres despues del ., si es mas de 4 INVALID .com .comm
