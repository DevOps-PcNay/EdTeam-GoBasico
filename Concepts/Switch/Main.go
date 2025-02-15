package main

import (
	"fmt"
	"github.com/DevOps-PcNay/EdTeam-GoBasico/edls/Emoji"
)

func main() {

	character := Emoji.Fig_emoji(3)
	canSearch := false

	/*
		switch character {
		case Emoji.Fig_emoji(1), Emoji.Fig_emoji(3):
			fmt.Println("Es una Hoja ", character)
		case Emoji.Fig_emoji(2):
			fmt.Println("Es una Palma ", character)
		default:
			fmt.Println("No aplica")
		}
	*/

	// Con esta forma se puede validar diferentes condiciones
	// Cumpliendo la primer condicion sale del "switch"

	switch {
	case !canSearch:
		fmt.Println("No se permite la busqueda")
	case character == Emoji.Fig_emoji(1) || character == Emoji.Fig_emoji(3):
		fmt.Println("Es una Hoja ", character)
	case character == Emoji.Fig_emoji(2):
		fmt.Println("Es una Palma ", character)
	default:
		fmt.Println("No aplica")
	}

}
