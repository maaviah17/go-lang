package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main(){

	fmt.Println("hey")
	// PerformGetReq()
	// PerformJsonReq()
	PerformPostFormReq()
}

func PerformGetReq(){
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content lenght : ", response.ContentLength)

	content , _ := io.ReadAll(response.Body)

	// fmt.Println(string(content))

	var responseString strings.Builder
	bytesCount, _ := responseString.Write(content) 
	fmt.Println("Bytes count : ", bytesCount)
	fmt.Println(responseString.String())
	


}

func PerformJsonReq(){
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
	 	{
			"coursename":"LesssGO with Golang",
			"price" : 199,
			"platform" : "muawiyahlearn.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	
	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormReq(){
	const myUrl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "muawiyah")
	data.Add("lastname", "khalid")
	data.Add("email", "mmk@mail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content , _ := io.ReadAll(response.Body)
	fmt.Println(string(content))


}