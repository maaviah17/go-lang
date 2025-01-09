package main

import (
	"fmt"
	"net/url"
)


const myUrl string = "https://learngolang:3020/learn?coursename=reactjs&paymentid=hur376fs&time=1700"

func main(){
	fmt.Println("Welcome to handling URL in Go")
	fmt.Println(myUrl)

	// parsing the url
	result, err := url.Parse(myUrl)

	if err!=nil{
		panic(err) 
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	// query parameters
	qparams := result.Query()
	fmt.Printf("types of query params are : %T\n", qparams)

	fmt.Println(qparams["paymentid"])

	for _,val := range qparams{
		fmt.Println("param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "http",
		Host: "muawiyah17",
		Path: "/baller",
		RawPath: "user=junior",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}