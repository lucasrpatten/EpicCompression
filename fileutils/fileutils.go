package fileutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(path string) (file *os.File) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	return file

}

func ReadFile(file string) (fileBytes []byte) {
	fileBytes, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return fileBytes
}
