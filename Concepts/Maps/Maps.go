package Maps

import (
	"fmt"
)

func Estructura_Maps() {
	// Es una estructua de datos de tipo "Clave - Valor"

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

	// Para eliminar un elemento. Se indica el nombre del mapa y la clave a eliminar.
	delete(tech, "computer")
	fmt.Println(tech)

	// Obteniendo un valor
	fmt.Println(music["violin"])

	// Los mapas retornan dos valores : Contenido y Valor Bool
	// Se utiliza para verificar si Existe.
	content, ok := music["guitar"]

	fmt.Println(content, ok)
}
