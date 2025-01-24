package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ScanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
func main() {
	ScanDirectory("D:")

}
