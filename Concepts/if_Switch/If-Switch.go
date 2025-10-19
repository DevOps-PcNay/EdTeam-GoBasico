package if_Switch

import (
	"fmt"
	"github.com/DevOps-PcNay/EdTeam-GoBasico/edls/Emoji"
	//"github.com/DevOps-Pcnay/EdTeam-GoBasico/Emoji"
)

// El alcance de la variable es a nivel paquete
// "character"
func If_Sentence() {
	// El alcance de la variable es a nivel Funcion
	// "character"
	character := Emoji.Fig_emoji(101)

	if character == Emoji.Fig_emoji(1) {
		// El alcance de la variable es a nivel IF
		// "character"

		fmt.Println("Es una hoja ", character)
	} else if character == Emoji.Fig_emoji(2) {
		fmt.Println("Es una palma ", character)
	} else {
		fmt.Println("No existe")
	}

} // func If_Sentence() {

func Switch_Sentence() {
	// Sentencia mas limpia para varias condiciones.
	character := Emoji.Fig_emoji(202)

	switch character {
	case Emoji.Fig_emoji(1):
		fmt.Println("Es una hoja", Emoji.Fig_emoji(1))
	case Emoji.Fig_emoji(2):
		fmt.Println("Es una palma", Emoji.Fig_emoji(2))
	default:
		fmt.Println("No Existe")
	} // switch character

} // func Switch_Sentence() {

func Switch2_Sentence() {
	// Usando multiples condiciones en un case
	var cadena1, cadena2 string
	var condicional_search bool

	cadena1 = "" // "Hulk"
	cadena2 = "Huason"
	condicional_search = true

	// Conestra formade Switch nos deja evaluar cualquier cosa que se requiera
	switch {
	// Usando una condicional antes de evaluar
	case !condicional_search:
		fmt.Println("Busqueda Activa")
	case cadena1 == "Hulk" || cadena1 == "4-Fantasicos":
		fmt.Println("Es un Heroe")
	case cadena2 == "Huason" || cadena2 == "Joker":
		fmt.Println("Es un Villano")
	default:
		fmt.Println("Indefinido")
	}

}
