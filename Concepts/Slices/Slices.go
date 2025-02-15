package main

import (
	"fmt"
)

func main() {
	// Slice = Son apuntadores a un Array, no posee datos.
	things := [7]string{"Carros", "Carros2", "Carros3", "Computadora", "Reloj", "Raton", "Monitor"}

	//Slice
	cars := things[0:2]

	fmt.Println("Things :", things)

	// No agrega el ultimo elemento "Carros3", se tendria que aumentar el indice a 3, (0:3)
	fmt.Println("Cars : ", cars)

	// Otro Slices
	fmt.Println("Ultimos 3", cars[4:7])

	// Accesando un elemento del Slice.
	fmt.Println("Accesando a un elemento del Slice Cars poscion 0 = ", cars[0])

	// Modificar un Slice.
	cars[1] = "Carro33"

	fmt.Println("Modificando el valor del Slice Cars , modificando Carro3 a ", cars)

	// cuando se cambia un valor en el Slide se cambia en el Arraglo, por esta razon son Punteros para el Array.

	// Tamano y Capacidad.
	// len() = Numero de elementos en el Slice
	// cap() = Numero de elementos del array Origen a partir del indice donde se crea el slice

	animal := [5]string{"Gorila", "Gato", "Perro", "Pajaro", "Camello"}
	pets := animal[1:3]
	fmt.Println("animals : ", animal)
	fmt.Println("Pets : ", pets)

	fmt.Println("Tamano de Pets ", len(pets))

	// La cpacidad dede donde inicia el Slice, en este caso es 1 al 4
	fmt.Println("Capacidad de Pets ", cap(pets))

	// Agregando un elemento al slice, tener encuenta que afecta al Array
	pets = append(pets, "Jirafa")
	fmt.Println("Imprimiendo un Slices con un elemento agregado ", pets)

	// Creando un Slices sin tomar la referencia un Array.
	pets2 := []string{"Perro", "Gato"}

	fmt.Println("Slices pets2 ", pets2)
	fmt.Println("Tamano de Pets2 ", len(pets2))
	fmt.Println("Capacidad de Pets2 ", cap(pets2))

	//Otra forma de Definir un "Slice"
	pets3 := make([]string, 0, 3)

	pets3 = append(pets3, "Jirafa", "Lagartija", "Hipototmo")

	fmt.Println("Slices pets3 ", pets3)
	fmt.Println("Tamano de Pets3 ", len(pets3))
	fmt.Println("Capacidad de Pets3 ", cap(pets3))

	// En el caso de que se rebase la capacidad, duplicara la capacidad
	pets3 = append(pets3, "Jirafa2")
	fmt.Println("Slices pets3 ", pets3)
	fmt.Println("Tamano de Pets3 ", len(pets3))
	fmt.Println("Capacidad de Pets3 ", cap(pets3))

	// Slices cuando es CERO
	// Declarando el Slice
	var pets4 []string
	fmt.Println("Slices pets4 ", pets4)
	fmt.Println("Tamano de Pets4 ", len(pets4))
	fmt.Println("Capacidad de Pets4 ", cap(pets4))

	// Asignando un tipo de datos
	// Cuando se declara de esta manera se inicializa el Slices.
	pets5 := []string{}
	fmt.Println("Slices pets5 ", pets5)
	fmt.Println("Tamano de Pets5 ", len(pets5))
	fmt.Println("Valor Cero ", pets5 == nil)

}
