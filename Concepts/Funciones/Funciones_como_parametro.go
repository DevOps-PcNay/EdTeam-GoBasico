package Funciones

import ()

func CallBack_mayores(filter_for int, num int) bool {
	return num > filter_for
}

func CallBack_menores(filter_for int, num int) bool {
	return num < filter_for
}

func Filter(nums []int, filter_for int, CallBack func(filter_for int, num int) bool) []int {
	// Se define el Slice de 0 elementos y la capacidad es de la longitud del arreglo de numeros que manden como parametro
	// La Funcion "callback" se define mas adelante. Se le pasa como parametro para haga la logica de filtrado.
	result := make([]int, 0, len(nums))
	for _, num := range nums {
		if CallBack(filter_for, num) {
			result = append(result, num)
		}
	}
	return result // Los numeros filtrados atraves de la funcion callback

}
