package main

import (
	"flag"
	"juego/models"
)

func main() {
	silent := flag.Bool("silent", false, "no reproducir sonido")
	flag.Parse()
	models.StartGame(*silent)
}
