package main

import (
	"fmt"
	"github.com/enescakir/emoji"
	"time"
)

const (
	// File Type
	fileRegular int = iota
	fileDirectory
	fileExecutble
	fielCompress
	fileImage
	fileLink
)

// Files Types
const ()

// File extension
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ",png"
	jpg = ".jpg"
	gif = ".gif"
)

type file struct {
	// Contendra la informacion de los archivos que son necesarios para renderizar
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string // Es el tipo de acceso que tiene en Owner, Grupo, Raiz
}

type styleFileType struct {
	icon   string
	color  string
	symbol string
}

func fig_emoji(opcion int) string {
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

var mapStyleByFileType = map[int]styleFileType{
	//fileRegular:   {icon: Emoji(1)},
	fileRegular:   {icon: fig_emoji(1)},
	fileDirectory: {icon: fig_emoji(2), color: "BLUE", symbol: "/"},
	fileExecutble: {icon: fig_emoji(3), color: "GREEN", symbol: "*"},
	fielCompress:  {icon: fig_emoji(4), color: "RED"},
	fileImage:     {icon: fig_emoji(5), color: "MAGENTA"},
	fileLink:      {icon: fig_emoji(6), color: "CYAN"},
}
