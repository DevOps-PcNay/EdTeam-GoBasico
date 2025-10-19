package Funciones

import "strings"

func Sum(a, b int) int {
	return a + b

}

// Ahora la funcion retornara mas de un valor.
func Convert(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return lower, upper
}
