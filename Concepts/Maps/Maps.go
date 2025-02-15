package main

import (
	"fmt"
)

func main() {
	music := make(map[string]string)
	music["guitar"] = "Guitarra"
	music["violin"] = "VIOLIN"

	fmt.Println(music)

	// Otra forma de definirla
	tech := map[string]string{
		"computer": "Desktop",
		"laptop":   "Laptop - Tablet",
	}

	fmt.Println(tech)

	// Para eliminar un elemento.
	delete(tech, "computer")
	fmt.Println(tech)

	// Obteniendo un valor
	fmt.Println(music["violin"])

	// Los mapas retornan dos valores : Contenido y Valor Bool
	// Se utiliza para verificar si Existe.
	content, ok := music["guitar"]

	fmt.Println(content, ok)

}
