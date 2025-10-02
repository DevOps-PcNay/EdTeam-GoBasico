package Operadores

import (
	"fmt"
)

func Operadores_arit() {
	var a = 2 + 3
	var b, c, d, e uint

	b = (3 + 3) * 10
	c = 4*5 + (10 + 23)
	d = c
	d += 4
	e = 0

	// Declaracion Postecremento
	e++

	fmt.Println("Valor de a ", a)
	fmt.Println("Valor de b ", b)
	fmt.Println("Valor de c ", c)
	fmt.Println("Valor de d ", d)
	fmt.Println("Valor de e ", e)
}
