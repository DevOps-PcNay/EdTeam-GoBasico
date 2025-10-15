package Arrays

import (
	"fmt"
)

func Arreglos() {

	// Los Array son de tamano Fijo
	// Declarando el Array
	var flags [3]string

	// Asignando valores al Array
	flags[0] = "Bandera Mexico"
	flags[1] = "Bandera Estados Unidos"
	flags[2] = "Bandera Irlandesa"

	// Otra forma de declarar los arreglos
	flags_colors := [3]string{"Bandera Roja", "Bandera Negra", "Bandera Blanca"}

	// Desplegando el contenido del Array
	fmt.Println(flags)
	fmt.Println(flags_colors)

	// Asignacion semidinamica.
	estados := [...]string{"Baja California", "Chihuahua", "Sonora", "Sinaloa"}

	fmt.Println(estados)

}
