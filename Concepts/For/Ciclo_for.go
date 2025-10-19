package For

import (
	"fmt"
)

func For_clasico() {
	for i := 0; i < 10; i++ {
		fmt.Println("Valor de - i - FOR clasico", i)
	}
}

func For_while() {
	i := 0
	for i < 10 {
		fmt.Println("Valor de - i - FOR Simula While ", i)
		i++
	}
}

// Ejecuntando "for" para siempre.
// Este es util cuando se quiere que se ejecuten procesos por siempre como:
// Web socket para que escuche indefinidamente
// Manejo de concurrencia, para que este escuchando siempre peticiones
/*
func For_infinito() {
	i := 0
	for {
		fmt.Println("Valor de -i- FOR infinito", i)
		i++
	}
}
*/

func For_infinito_break() {
	i := 0
	for {

		if i == 6 {
			break
		}

		fmt.Println("Valor de -i- FOR infinito con Break", i)
		i++
	}
}

func For_range() {
	slice := [3]string{"Valor1", "Valor2", "Valor3"}
	// Tambiense puede agregar el arreglo de slice
	for indice, valor := range slice {
		fmt.Println("Indice = ", indice, " valor = ", valor)
	}

}

func For_modif_slice() {
	numbers := []uint8{2, 4, 6, 8}
	// Para simplificarlo se suprime la columna "valor" incluyendo "_"
	for i := range numbers {
		// Para que se modifique el contenido del slice "numbers", se tiene que accesar al indice del arreglo
		numbers[i] *= 2
	}
	fmt.Println("Valor modificado Slice ", numbers)
}

// Usandolo con "map"
func For_map() {
	food := map[string]string{
		"P": "Pizza",
		"H": "Hamburguesa",
		"A": "Manzana",
		"M": "Mango",
	}

	for key, value := range food {
		fmt.Println("Comida ", key, " valor = ", value)
	}

}

func For_string() {
	// value = Cuando se maneja rangos con cadenas el "value" es de tipo "rune"
	// string(value) = Para que imprima el caracter.
	for index, value := range "Hola Mundo" {
		fmt.Println("Indice ", index, "Valor ", string(value))
	}
}
