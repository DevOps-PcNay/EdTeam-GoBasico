package Variables

import (
	"fmt"
)

func Variable() {
	// Formas de declarar variables
	var apple string

	// Creando otras variables:
	var banana string
	var orange string

	apple = "Apple"
	banana = "Banana"
	orange = "Orange"

	// Otra forma de declarar variables :
	var apple2 string = "Apple2"
	var banana2 string = "Banana2"
	var orange2 string = "Orange2"

	// Otra forma de declarar variables :
	var (
		apple3  string = "Apple3"
		banana3 string = "Banana3"
		orange3 string = "Orange3"
	)

	// Otra forma de declarar variables :
	// Mas compacta cuando son varias del mismo tipo

	var apple4, banana4, orange4 string = "Apple4", "Banana4", "Orange4"

	// Declarando inferiendo el tipo:
	apple5, banana5, orange5, number := "Apple5", "Banana5", "Orange5", 20

	fmt.Println("Dsplegando contenido de Variables ", apple)
	fmt.Println("Dsplegando contenido de Variables ", banana)
	fmt.Println("Dsplegando contenido de Variables ", orange)

	fmt.Println("Dsplegando contenido de Variables ", apple2)
	fmt.Println("Dsplegando contenido de Variables ", banana2)
	fmt.Println("Dsplegando contenido de Variables ", orange2)

	fmt.Println("Dsplegando contenido de Variables ", apple3)
	fmt.Println("Dsplegando contenido de Variables ", banana3)
	fmt.Println("Dsplegando contenido de Variables ", orange3)

	fmt.Println("Dsplegando contenido de Variables ", apple4)
	fmt.Println("Dsplegando contenido de Variables ", banana4)
	fmt.Println("Dsplegando contenido de Variables ", orange4)

	fmt.Println("Dsplegando contenido de Variables ", apple5)
	fmt.Println("Dsplegando contenido de Variables ", banana5)
	fmt.Println("Dsplegando contenido de Variables ", orange5)
	fmt.Println("Dsplegando contenido de Variables ", number)

}

/*
package main

import (
	"fmt"
)

func main() {
	//var apple string
	//var pera string = "Pera"

	//var apple string = "Apple"
	//var pera string = "Pera"

	// var apple, pera string = "Apple", "Pera"
	apple, pera := "Apple", "Pera"
	fmt.Println("Imprimiendo contenido variable ", apple, pera)

}
*/
