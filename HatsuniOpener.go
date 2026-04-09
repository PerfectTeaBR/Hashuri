package main

import (
	"log"
	"os"
)

func HatsuniOpener() {

	file, err := os.Open("logs.txt")


	if err != nil {
		log.Fatal("HatsuniOpener não conseguiu abrir o arquivo", file)
	}

}
