package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	err := os.Mkdir("dist", 0755)
	if err != nil {
		fmt.Println(err)
	}

	files, err := os.Open("src")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(files)
}