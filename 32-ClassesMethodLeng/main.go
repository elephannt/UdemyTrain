package main

import "fmt"

type Empleado struct {

	nombre string
	direccion string
}


//METODO O FUNCION QUE NOS PERMITE OBETENER EL NOMBRE

func (h Empleado) ObtenerNombre() {
	fmt.Printf("\nHola, mi nombre es: %s\n", h.nombre)
}

//METODO O FUNCION QUE NOS PERMITE ESTABLECER UN NOMBRE

func EstablecerNombre(x string) Empleado {
	nuevo := Empleado{x," ,Sin direccion"}
        return nuevo
}

func main() {

        pedro := Empleado{"Pedro Martinez Seaz","C/ Trujillo N-20"}
	sonia := Empleado{"Sonia Garcia Soto","Plaza Espana N-11"}

	var p,s Empleado

	p = pedro
	p.ObtenerNombre()
	s = sonia
	s.ObtenerNombre()

	var j string

	fmt.Println("\nEstablecer el nombre: ")
        fmt.Scan(&j)
	l := EstablecerNombre(j)
	fmt.Println(l)
}