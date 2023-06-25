package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: shred <file_path>")
	}

	filepath := os.Args[1]
	Shred(filepath)
}
