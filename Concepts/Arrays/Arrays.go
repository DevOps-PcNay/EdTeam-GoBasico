package main

import (
	"fmt"
)

func main() {

	// Los Array son de tamano Fijo
	// Declarando el Array
	// var flags [3]string

	// Asignando valores al Array
	//flags[0] = "Bandera Roja"
	//flags[1] = "Bandera Negra"
	//flags[2] = "Bandera Blanca"

	// Otra forma de declarar los arreglos
	flags := [3]string{"Bandera Roja", "Bandera Negra", "Bandera Blanca"}

	// Desplegando el contenido del Array
	fmt.Println(flags)

}
