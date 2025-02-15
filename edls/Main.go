package main

// Paquete para la linea de comandos, para pasar parametros (banderas)
import (
	"flag"
	"fmt"
	//"github.com/DevOps-Pcnay/EdTeam-GoBasico/Emoji"
)

func main() {

	// Filter patter

	// Devuelve una direccion de memoria.
	// Se envia banderas en la linea de comandos, puede ser "-" o "--"
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "All File Including Hide Files")
	flagNumberRecords := flag.Int("n", 0, "Number Of Records")

	// Orders Flags, Ordenar la salida.
	hasOrderbyTime := flag.Bool("t", false, "Sort By time, oldest Firts")
	hasOrderbySize := flag.Bool("s", false, "Sort By size Smallest Firts")
	hasOrderReverse := flag.Bool("r", false, "Reverse order while sorting")

	// Mapea cada uno de los "flag" para almacenarlo en la variable para ser utilizado, esta linea siempre se realiza
	flag.Parse()

	// Muestra la direccion de memoria
	fmt.Println(flagPattern)

	// Para mostrar el contenido, se utiliza el simbolo de redesferenciacion "*""
	fmt.Println("pattern = ", *flagPattern)
	fmt.Println("All = ", *flagAll) // Obtiene todos los archivos del directorio incluso los ocultos.
	fmt.Println("view = ", *flagNumberRecords)
	fmt.Println("Order By Date ", *hasOrderbyTime)
	fmt.Println("Order By Size ", *hasOrderbySize)
	fmt.Println("Order Reverse ", *hasOrderReverse)

	//Folder := fmt.Sprintf("%v", emoji.OpenFileFolder)
	//fmt.Println(Folder)

	// fmt.Println(emoji.FoldedHands) //Archivo Regular -> 1
	// fmt.Println(emoji.PalmTree)         //Directorio   -> 2
	// fmt.Println(emoji.Rocket + "Texto") // EXE		-> 3
	// fmt.Println(emoji.Books)            // Comprimido		-> 4
	// fmt.Println(emoji.FramedPicture)    //Imagen		-> 5
	// mt.Println(emoji.Link)             //Enlace		-> 6

	// fmt.Println(emoji.FileFolder)  // Carpeta

}
