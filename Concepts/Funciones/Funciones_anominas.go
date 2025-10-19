package Funciones

import (
	"fmt"
)

func Funcion_anonima() {

	// Pasando como parametro y autoejecutandose
	/*
		func(name string) {
			fmt.Println("Hola ", name)
		}("Gente")
	*/

	saludos := func() {
		fmt.Println("Hola, soy una funcion anonima")
	}
	saludos()

}
