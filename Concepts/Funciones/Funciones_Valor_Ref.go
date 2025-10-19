package Funciones

import (
	"fmt"
	"strings"
)

func Funcion_sin_Parametro() {
	fmt.Println("Hola Mundo")
}

func Funcion_con_parametros(cadena1, cadena2 string) {
	fmt.Println("Uniendo los parametros Cadena1 = ", cadena1+" Cadena2 = ", cadena2)
}

// por defecto las funciones siempre trabajan por defecto "por  valor"
// Por lo que se tiene que modificar en la funcion para que pueda trabajar por referentencia
// *string = Es el puntero de la cadena a que se esta pasando.
// *text = Es la desrefenciacion del puntero para poder accesar al valor de la direccion de memoria
func Convertir_Mayus(texto *string) {
	*texto = strings.ToUpper(*texto)
	fmt.Println("Cadena convertida en Mayusculas", *texto)
}
