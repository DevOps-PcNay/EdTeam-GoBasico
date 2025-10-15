package Slices

import (
	"fmt"
)

func Slice() {

	// Slice = Son apuntadores a un Array, no posee datos.
	things := [7]string{"Carros", "Carros2", "Carros3", "Computadora", "Reloj", "Raton", "Monitor"}

	//Slice, indice 0 al 2
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

	// Se altera el arreglo principal (things), aun cuando se haya modificado en un slices, por esta razon los slices hacen referenia a la direccion de memoria.
	ultimo_3 := things[4:]
	fmt.Println("Imprimiendo los ultimos 3 elementos del arrelgo ", ultimo_3)

	primeros_3 := things[:3]
	fmt.Println("Imprimiendo los primeross 3 elementos del arrelgo ", primeros_3)

	todo := things[:]
	fmt.Println("Impirmiendo todo el arreglo\n ", todo)

	// Tamano y Capacidad.
	// len() = Numero de elementos en el Slice
	// cap() = Numero de elementos del array Origen a partir del indice donde se crea el slice

	animal := [5]string{"Gorila", "Gato", "Perro", "Pajaro", "Camello"}
	pets := animal[1:3]
	fmt.Println("animals : ", animal)
	fmt.Println("Pets : ", pets)

	// Tamano = Numero de elementos que contiene el Slice.
	fmt.Println("Tamano de Pets ", len(pets))

	// La capacidad = Es el numero de elementos del array origen apartir del indice que nosotros especificamos para crear nuestro slice
	// En este caso es 4 elementos (iniciando de "Gato" terminando en "Camello")
	fmt.Println("Capacidad de Pets ", cap(pets))

	// Con los "Slices" se puede agregar elmentos al arreglo de forma dinamica, esto se hace atraves de una funcion preconstruida llamada "append()"

	pets = append(pets, "Jirafa")
	fmt.Println("Imprimiendo un Slices con un elemento agregado ", pets)

	// Pero se modifica el array principal "animal" se debe tener encuenta
	fmt.Println("Imprimiendo el Slice principal -animal- ", animal)

	// Ahora se rebasara la capacidad del slice "pets" que sucedera ?
	pets = append(pets, "Cocodrilo", "Gasela", "Tigre")
	fmt.Println("Imprimiendo el Slice pets", pets)

	// Go land de forma automatica duplica la caapcidad el forma doble de la capaciad actual.
	fmt.Println("Capacidad de Pets ", cap(pets))
	fmt.Println("Imprimiendo el Slice principal -animal- ", animal)

	// Lo que hace Go Land es que como ya no se puede agregar mas elementos al arreglo "pets"[Slices], crea uno nuevo "pets" y le agrega los elementos nuevos, ademas ya no modifica el arreglo original "animal" y por tanto el slice "pets" ya no hace referencia al arreglo "animal"

	// Creando un Slices sin tomar la referencia un Array.
	// No se especifica el tamano.
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

	// En el caso de que se rebase la capacidad, duplicara la capacidad, para este caso de 3 a 6.
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
