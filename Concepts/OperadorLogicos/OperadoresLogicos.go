package main

import (
	"fmt"
)

func main() {

	// Compara dos valores y retorna un boolean
	// <, >, ==, != >= <=
	fmt.Println(4 > 6)
	fmt.Println((4 + 8) < 6)

	// Operadores logicos: && , ||
	var age uint = 33
	fmt.Println("Es adulto? ", age >= 18 && age <= 60)
	fmt.Println("Es nino o Adulto ? ", age >= 18 || age <= 60)

	// Operadores logicos unario : !
	fmt.Println(!(4 == 4))

}
