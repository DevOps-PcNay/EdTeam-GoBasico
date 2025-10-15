package main

import (
	"flag"
	"fmt"
)

func main() {
	// Definiendo el "flag" para filtar por extension
	flagPattern := flag.String("p", "", "filter by pattern")

	// Definiendo el "flag" para filtrar para que muestre todos los archivos.
	flagAll := flag.Bool("a", false, "all files including hide files ")
	// Para mapear cada uno de los flags(parametros) que se envien.

	// Para limitar los registros a mostrar en pantalla.
	flagNumberRecords := flag.Int("n", 0, "Number of record to show ")

	// flag para organizar la salida
	// flag order
	hasOrderByTime := flag.Bool("t", false, "sort by time oldest to first")
	hasOrderBySize := flag.Bool("s", false, "sort by file size small to first ")
	hasOrderReverse := flag.Bool("r", false, "Reverse order while display")

	flag.Parse()
	// Para que muestre el contenido de la variable donde se asigno el "flag"
	fmt.Println(*flagPattern)
	fmt.Println(*flagAll)
	fmt.Println(*flagNumberRecords)
	fmt.Println(*hasOrderByTime)
	fmt.Println(*hasOrderBySize)
	fmt.Println(*hasOrderReverse)
}
