package main

import (
	"fmt"
)

func main() {

	// Inferir el valor del tipo de datow de la Constante.
	//const os, fecha, numero = "Linux", "Fecha de Hoy", 10

	// otra forma de declarar constantes:
	const (
		os     = "Linux"
		fecha  = "Fecha de Hoy"
		numero = 10
	)

	// Uitlizando el operador "iota" para incrementar en 1 el valor.
	const (
		//	Jan = 1
		//	Feb = 2
		//	Mar = 3
		//	Apr = 4
		Jan = iota + 1
		Feb
		Mar
		Apr
	)

	fmt.Println(os, fecha, numero)
	fmt.Println("Desplegando el valor de mes ", Jan, Feb, Mar, Apr)

}
