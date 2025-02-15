package Emoji

import (
	"fmt"
	"github.com/enescakir/emoji"
)

// Funcion publica
func Fig_emoji(opcion int) string {
	switch opcion {
	case 1: // Archivo Regular
		Folder := fmt.Sprintf("%v", emoji.FoldedHands)
		return Folder
	case 2: // Directorio
		Folder := fmt.Sprintf("%v", emoji.PalmTree)
		return Folder
	case 3: // EXE
		Folder := fmt.Sprintf("%v", emoji.Rocket)
		return Folder
	case 4: // Comprido
		Folder := fmt.Sprintf("%v", emoji.Books)
		return Folder
	case 5: // Imagen
		Folder := fmt.Sprintf("%v", emoji.FramedPicture)
		return Folder
	case 6: // Enlace
		Folder := fmt.Sprintf("%v", emoji.Link)
		return Folder
	}
	return ""
}
