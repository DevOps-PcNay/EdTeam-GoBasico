package main

import (
	"fmt"
)

// Puntero : Variable que almacenan la direccion en memoria de un valor
func main() {
	var color string = "red"
	// Definiendo una variable de tipo puntero.
	var pointerColor *string

	// Asignando al puntero el valor de la direccion de memoria
	// Obtener la direccion de memmoria para asignarlo, es decir esta apuntando.
	pointerColor = &color

	// Asignando un valor a la direccion de memoria
	*pointerColor = "Azul"

	fmt.Printf("Tipo : %T, Valor : %s  Direccion %v \n ", color, color, &color)
	// *pointerColor = Es la Desreferenciacion
	fmt.Printf("Tipo :%T, Valor %v, Desreferenciacion: %s \n ", pointerColor, pointerColor, *pointerColor)

}
