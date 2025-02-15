package main

import (
	"fmt"
)

func main() {
	// Tipos de datos : Bool, String, numeric
	var a bool = true
	var cadena string = "EDTeam"
	var edad uint8 = 30
	var valor_ascII rune = 'b'

	// Tipo de Datos Unicot
	fmt.Printf("Tipo : %T, valor : %v \n ", a, a)
	fmt.Printf("Tipo : %T, valor : %v \n ", cadena, cadena)
	fmt.Printf("Tipo : %T, valor : %v \n ", edad, edad)
	fmt.Printf("Tipo : %T, valor : %v \n ", valor_ascII, valor_ascII)
}
