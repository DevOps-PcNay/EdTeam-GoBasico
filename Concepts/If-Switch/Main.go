package main

import (
	"fmt"

	"github.com/DevOps-PcNay/EdTeam-GoBasico/edls/Emoji"
	//"github.com/DevOps-Pcnay/EdTeam-GoBasico/Emoji"
)

// El alcance de la variable es a nivel paquete
// "character"
func main() {
	// El alcance de la variable es a nivel Funcion
	// "character"
	character := Emoji.Fig_emoji(2)

	if character == Emoji.Fig_emoji(1) {
		// El alcance de la variable es a nivel IF
		// "character"

		fmt.Println("Es una hoja ", character)
	} else if character == Emoji.Fig_emoji(2) {
		fmt.Println("Es una palma ", character)
	} else {
		fmt.Println("No existe")
	}

}
