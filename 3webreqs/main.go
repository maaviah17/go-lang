package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"


func main(){

	fmt.Println("Hey, working with HTTP")

	response,err := http.Get(url)
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response : %T\n", response)
	
	// callers responsibility to close the the connection
	defer response.Body.Close()

	databytes,err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}