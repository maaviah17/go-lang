package main

import "fmt"

func main(){

	msg := 31
	fmt.Println(msg)

	var p *int32 = new(int32) 
	var i int32 = 17

	// *p = 10
	p = &i
	*p=1
	fmt.Println( p,i)

}