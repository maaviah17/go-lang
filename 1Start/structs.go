package main 

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

type Person struct {
	Name string
	Age int 
	Country string
}



// this is Interface :
type Speaker interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "woof"
}

func (c Cat) Speak() string {
	return "meow"
}


func structs(){

	// this is struct : it groups together variables under a single name, representing an entity with multiple attributes
	var myEngine gasEngine = gasEngine{mpg:25, gallons: 15}
	fmt.Println("hello, here are some info")
	fmt.Println(myEngine.mpg, myEngine.gallons)

	p1 := Person{
		Name : "Muawiyah",
		Age: 20,
		Country: "India",
	}

	p2 := Person{"Jane", 25, "Canada"}

	fmt.Println(p1)
	fmt.Println(p2)

	// use of Interface : 
	var s Speaker

	s = Dog{}
	fmt.Println(s.Speak())

	s = Cat{}
	fmt.Println(s.Speak())

	// this is called anonymous function
	sum := func (a,b int) int {
		return a+b
	}

	fmt.Println("Sum : ", sum(2,3))
}