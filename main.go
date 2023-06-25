package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Usage: shred <file_path>")
	}

	filepath := os.Args[1]
	Shred(filepath)
}
