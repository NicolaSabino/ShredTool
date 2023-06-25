package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

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

func Shred(p string) {

	log.Println("Open file at", p)
	f, err := os.OpenFile(p, os.O_RDWR, 0666)

	if err != nil {
		log.Fatalln("Unable to open file", err)
	}

	info, err := f.Stat()
	if err != nil {
		f.Close()
		log.Fatalln("Unable to read file stat", err)
	}
	fSize := info.Size()
	log.Println("File size", fSize)

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

	log.Println("Close file")
	err = f.Close()
	if err != nil {
		log.Fatalln("Unable to write file", err)
	}

	log.Println("Delete file")
	err = os.Remove(p)

	if err != nil {
		log.Fatalln("Unable to delete file", err)
	}
}
