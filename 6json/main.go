package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main(){

	// fmt.Println("Hey")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){

	myCourses := []course{
		{"Reactjs Bootcamp", 199, "MuawiyahLearn.in" , "pass123", []string{"web-dev", "html | css | js"}},
		{"Web3 CoHort", 249, "MuawiyahLearn.in" , "ssaapp321", []string{ "web3" , "blockchain"}},
		{"Rust marathon", 299, "MuawiyahLearn.in" , "word454", nil},
	}

	// package this data as json data

	finalJson , err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson(){

	jsonDataFromWeb := []byte(`
	{
		"coursename": "Web3 CoHort",
		"Price": 249,
		"website": "MuawiyahLearn.in",
		"tags": [
				"web3","blockchain"]
		}
	`)

	var myCourse course 

	checkValid  := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Printf("JSON valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Printf("JSON not valid !!!\n")
	}


	// some cases where we want to add data to key value

	//using interface because we don't know what tyoe of data will come in
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k,v := range myOnlineData {
		fmt.Printf("Key is : %v and Value is : %v\n", k , v)
	}
}