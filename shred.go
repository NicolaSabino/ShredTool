package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

// getRandData return a slice of random bytes
// according to given input size
func getRandData(size int) []byte {
	buf := make([]byte, size)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	_, err := r.Read(buf)
	if err != nil {
		// if rand fails buf contains zeroes
		log.Panicln("Error while generating random data:", err)
	}
	return buf
}

// Shred apply the shred logic to given
// input file path `p`
//
// If by any chance an error occurs during
// the shred execution a `panic` is issued
// and has to be handled by the caller
func Shred(p string) {

	// Open the target file stored at `p`
	log.Println("Open file at", p)
	f, err := os.OpenFile(p, os.O_RDWR, 0666)
	if err != nil {
		log.Panicln("Unable to open file", err)
	}

	// Obtain file dimension via stat command
	info, err := f.Stat()
	if err != nil {
		f.Close()
		log.Panicln("Unable to read file stat", err)
	}
	fSize := info.Size()
	log.Println("File size", fSize)

	// Overwrite three times the file content
	// with random data
	log.Println("Write file with random data")
	for i := 0; i < 3; i++ {
		random := getRandData(int(fSize))
		_, err := f.WriteAt(random, 0)
		if err != nil {
			f.Close()
			log.Panicln("Unable to write file", err)
		}
		err = f.Sync()
		if err != nil {
			f.Close()
			log.Panicln("Unable to sync file", err)
		}
	}

	// Close the target file
	log.Println("Close file")
	err = f.Close()
	if err != nil {
		log.Panicln("Unable to close file", err)
	}

	// Delete the target file
	log.Println("Delete file")
	err = os.Remove(p)
	if err != nil {
		log.Fatalln("Unable to delete file", err)
	}
}
