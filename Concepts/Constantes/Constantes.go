package Constantes

import (
	"fmt"
)

// otra forma de declarar constantes, a nivel de paquete
const (
	os     = "Linux"
	fecha  = "Fecha de Hoy"
	numero = 10
)

func Constante() {

	// Inferir el valor del tipo de datow de la Constante.
	const pi = 3.1415

	// Esta a nivel local de funcion.
	// Uitlizando el operador "iota" para incrementar en 1 el valor.
	// Se agrega uno porque "iota" inicia en 0
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

	fmt.Println("Valor constante pi ", pi)
	fmt.Println(os, fecha, numero)
	fmt.Println("Desplegando el valor de mes ", Jan, Feb, Mar, Apr)

}
