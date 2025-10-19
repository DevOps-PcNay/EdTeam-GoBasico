package main

import (
	"fmt"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Constantes"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Variables"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Comentarios"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Operadores"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/OperadorLogicos"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Punteros"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Arrays"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Slices"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Maps"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Estructuras"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/if_Switch"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/For"
	//"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Funciones"
	"github.com/DevOps-PcNay/EdTeam-GoBasico/Concepts/Errores"
)

func main() {
	//Constantes.Constante()
	//Variables.Variable()
	//Constantes.Constante()
	//Comentarios.Comentario()
	// Operadores.Operadores_arit()
	//OperadorLogicos.OperadoresLogicos()
	//Punteros.Puntero()
	//Arrays.Arreglos()
	//Slices.Slice()
	//Maps.Estructura_Maps()
	//Estructuras.TipoDatoEstructura()
	//if_Switch.If_Sentence()
	//if_Switch.Switch_Sentence()
	//if_Switch.Switch2_Sentence()
	//For.For_clasico()
	//For.For_while()
	//For.For_infinito()
	//For.For_range()
	//For.For_modif_slice()
	//For.For_map()
	//For.For_string()
	//Funciones.Funcion_sin_Parametro()
	//Funciones.Funcion_con_parametros("Valor1", "Valor2")
	//var cadena = "convertir para mayusculas"

	// &cadena = El "&" obtiene la direccion de la memoria de "cadena"
	//Funciones.Convertir_Mayus(&cadena)

	/* Seccion **Funciones **
	result := Funciones.Sum(5, 10)
	fmt.Println("El valor de lo retorna la Funcion -Sum- ", result)

	minusculas, mayusculas := Funciones.Convert("MaYuCuLaS y MiNuScUlAs")
	fmt.Println("Minusculas : ", minusculas, "Mayusculas : ", mayusculas)

	var filter_for int = 7
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	number_filters := Funciones.Filter(nums, filter_for, Funciones.CallBack_mayores)
	fmt.Println("Numeros filtrados : ", number_filters)

	number_filters_menor := Funciones.Filter(nums, filter_for, Funciones.CallBack_menores)
	fmt.Println("Numeros filtrados : ", number_filters_menor)

	results := Funciones.Sums(2)(3)
	fmt.Println("El resultado de Funcion que retorna una funcion = ", results)

	// Ejecuntado por parte las funciones.
	parte1_sum := Funciones.Sums(10)
	result_sum := parte1_sum(10)
	result_sum2 := parte1_sum(30)

	fmt.Println("El resultado de la suma es : ", result_sum)
	fmt.Println("El resultado de la suma es : ", result_sum2)
	*/

	/*
		totales := Funciones.Sum_n(4)
		fmt.Println("El total de la suma variatica es : ", totales)
		totales2 := Funciones.Sum_n(34, 56, 45, 39)
		fmt.Println("El total de la suma variatica es : ", totales2)

		Funciones.Funcion_anonima()
	*/
	// Errores.Errores_go()
	encontro, err := Errores.Buscar2_Cadena("23")

	if err != nil {
		fmt.Printf("Error en Errores.Buscar2_Cadena() %v \n ", err)
		return
	}

	if len(encontro) != 0 {
		fmt.Println("Valor Encontrado : ", encontro)
	}

}
