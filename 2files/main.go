package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	fmt.Println("Welcome to files in GoLang")
	content := "This content needs to go in a file."

	file , err := os.Create("./myContent.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length of the file is :" , length)
	defer file.Close()

	readFile("./myContent.txt")

}

func readFile(filename string){

	databyte,err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
} 