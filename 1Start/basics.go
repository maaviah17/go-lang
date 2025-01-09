package main

import "fmt"

func concat(s1, s2 string) string{
	return (s1+s2)
}

// was func main but had to change name because created another new main func
func basics(){
	
	// fmt.Println("Hey Muawiyah")
	// num := 2
	// num++
	// num--
	// num++
	// name := `Hello
	// my name is Muawiyah`

	// name :="hello"
	msg := concat("my name's " , "not khalid")
	fmt.Println(msg)

	naaaam:= "muawiyah"
	if len(naaaam) < 7 {
		fmt.Println("you're gay")
	}	else{
		fmt.Println("not gay")
	}

	
	
}
