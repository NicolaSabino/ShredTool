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
)

// crateDummyTextFile create a text file
// in `./tmp` folder
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

// createDummyBinFile create binary file
// in `./tmp` folder based on given file mode
func createDummyBinFile(m fs.FileMode) {
	dummyContent := []byte{0x12, 0x34, 0x56, 0x78}
	err := ioutil.WriteFile(dummyBinFile, dummyContent, m)
	if err != nil {
		log.Fatalln("Unable to create dummy file", err)
	}
}

// TestShredText shred textual file
//
// The file exists and everything should be
// performed without issues
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

// TestShredBin shred binary file
//
// The file exists and everything should be
// performed without issues
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

// TestMissingFile shred missing file
//
// We expect a `panic` to be raised by `Shred`
// function while calling a non existing file
func TestMissingFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No expected panic `missing file`")
		}
	}()

	Shred("noFile.txt")
}

// TestNoAccess shred file without permissions
//
// Try to shred file owned by root user
// We expect a `panic` to be raised
func TesNoAccess(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("No expected panic `permission denied`")
		}
	}()
	Shred(rootFile)
}
