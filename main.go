package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getFiles(folderPath string) []string {
	var files []string
	err := filepath.Walk(folderPath, func (path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error")
	}

	return files
}

func main() {
	fmt.Println("Hello world!")
	err := os.Mkdir("dist", 0755)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(getFiles("./src"))
}