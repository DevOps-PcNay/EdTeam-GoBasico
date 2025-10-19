package Errores

import (
	"errors"
	"fmt"
	"strconv"
)

// Creando un error.
var errNotFound = errors.New("Not Found Error	")

var food = map[int]string{
	1: "Apple",
	2: "Banana",
	3: "Orange",
}

/*
func search(key string) (cadena string, err error) {
	fmt.Println(key)
	return "", errNotFound
}


func Buscar_Cadena() {
	found, err := search("a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Valor encontrado : ", found)

}
*/

/*
func Errores_go() {
	num, err := strconv.Atoi("ddfd")
	if err != nil { // Hubo un error
		fmt.Println(err)
		return // Ya no se continue ejecutando
	}
	fmt.Println(num)
}
*/

func Buscar2_Cadena(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strvconv.Atoi(): %w ", err)
	}
	encontro, err := findFood(num)

	//if err == errNotFound {
	if errors.Is(err, errNotFound) {
		fmt.Println("Pudimos controlar el error ")
		return "", nil
	}

	return encontro, err
}

func findFood(id int) (string, error) {
	value, ok := food[id]
	if !ok { // Si no existe en el "map"
		return "", fmt.Errorf("Error en -findFood- %w  ", errNotFound) // Retorna el error personalizado
	}
	return value, nil
}
