package Funciones

import ()

// Modificacion para que pueda recibir "n" numeros de valores.
// Go Land lo uso como un slice internamente.
func Sum_n(nums ...int) (total int) {
	var totales int = 0
	for _, num := range nums {
		totales += num
	}
	return total
}
