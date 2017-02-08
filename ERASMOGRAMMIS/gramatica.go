/*
	LENGUAJES Y AUTOMATAS I
	HERNANDEZ CONTRERAS UZZIEL
	13211489
*/

package main

func evaluacion() string {
	return "(" + id() + "|" + numeros() + "){1}(" + opRel() + ")?(" + id() + "|" + numeros() + "){1}"
}

func claveS() string {
	return pregunta() + parDer() + evaluacion() + parIzq() + agDer()
}

func claveC() string {
	return pregunta() + parDer() + evaluacion() + parIzq() + agDer() + bloque() + agIzq()
}

func asignacionSimple() string {
	return id() + asigna() + "(" + numeros() + "|" + id() + "){1}$"
}

func claseN() string {
	return clase() + " " + id()
}

func asignacionOp() string {
	return id() + asigna() + numeros() + opMat() + numeros()
}

func asignacionIdID() string {
	return id() + asigna() + id() + "$"
}

func asignacionIdNum() string {
	return id() + asigna() + id() + opMat() + numeros()
}

func asignacionNumId() string {
	return id() + asigna() + numeros() + opMat() + id() + "$"
}

func asignacionIdIdId() string {
	return id() + asigna() + id() + opMat() + id() + "$"
}

func principal() string {
	return "^" + elMain() + "$"
}
