package Estructuras

import (
	"fmt"
)

// Estas estructuras las declaran a nivel de paquetes. Para que esten disponibles en otros paquetes o funciones.

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func TipoDatoEstructura() {

	// Es mejor definirlo de esta manera, ya que cuando son demasiados se pierde.
	// Ademas se puede definir los campos que se desean.
	Juan := Person{
		Name:        "Juan",
		Age:         45,
		HasChildren: false,
	}

	//%+v = Para que imprima la estructura.
	fmt.Printf("%+v \n", Juan)

	// Definiendo otra instancia :
	Beto := Person{"Ortiz", 20, false}

	fmt.Printf("%+v \n", Beto)

	// Definiendo un solo campo. Se recomienda definir el nombre del campo, para asegurar que se esta asignando.

	Paco := Person{Age: 30}
	fmt.Printf("%+v \n", Paco)

	// Accesando a los campos de la estructura.
	fmt.Println("Accesando al campo Nombre : ", Beto.Name)

	// Punteros en Estructuras
	Alvaro := &Juan

	fmt.Printf("%+v \n", Juan)
	fmt.Printf("%+v \n", Alvaro)

	// Modificando el contenido utilizando un puntero.
	//Usando el Desreferenciacion, para cambiar un campo de la estructura de Juan, ya que "Alvaro" esta apuntando a "Juan"
	(*Alvaro).Age = 10
	// Tambien se puede realizar :
	Alvaro.Name = "Juan Perez"

	fmt.Println(Juan.Age)
	fmt.Println(Juan.Name)

}
