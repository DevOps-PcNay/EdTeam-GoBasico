package Funciones

import ()

func Sums(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
