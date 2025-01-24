package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ScanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := ScanDirectory(filePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}
func main() {
	err := ScanDirectory("D:")
	if err != nil {
		log.Fatal(err)
	}

}
