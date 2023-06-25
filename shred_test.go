package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const (
	dummyTextFile = "tmp/dummy.txt"
	dummyBinFile  = "tmp/dummy.bin"
	rootFile      = "sample/root_file_sample.bin"
	writeOnlyFile = "samples/write_only_file_sample.txt"
	readOnlyFile  = "samples/read_only_file_sample.txt"
)

func createDummyTextFile() {
	f, err := os.Create(dummyTextFile)

	if err != nil {
		log.Fatalln("Unable to create dummy file", err)
	}

	defer f.Close()

	_, err = f.WriteString("Dummy content\n")

	if err != nil {
		log.Fatalln("Unable to write content in dummy file", err)
	}
}

func createDummyBinFile(m fs.FileMode) {
	dummyContent := []byte{0x12, 0x34, 0x56, 0x78}
	err := ioutil.WriteFile(dummyBinFile, dummyContent, m)
	if err != nil {
		log.Fatalln("Unable to create dummy file", err)
	}
}

func TestShredText(t *testing.T) {
	createDummyTextFile()
	Shred(dummyTextFile)

	_, err := os.Stat(dummyTextFile)
	if os.IsNotExist(err) {
		t.Log("File successfully deleted")
	} else if err != nil {
		t.Fatalf("Error occurred while checking file existence: %s", err)
	} else {
		t.Errorf("File should not exist, but it does: %s", dummyTextFile)
	}
}

func TestShredBin(t *testing.T) {
	createDummyBinFile(0666)
	Shred(dummyBinFile)

	_, err := os.Stat(dummyBinFile)
	if os.IsNotExist(err) {
		t.Log("File successfully deleted")
	} else if err != nil {
		t.Fatalf("Error occurred while checking file existence: %s", err)
	} else {
		t.Errorf("File should not exist, but it does: %s", dummyTextFile)
	}
}

func TestMissingFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No expected panic `missing file`")
		}
	}()

	Shred("noFile.txt")
}

func TestNoAccessToMetadata(t *testing.T) {
}