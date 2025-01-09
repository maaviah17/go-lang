package main

import (
	"fmt"
	"time"
)


func main() {
	// fmt.Println("Welcome to Routines and Channels lecture")
	// fmt.Printf("\n")

	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func(){
		// sending data
		myChannel <- "sending data to channel" 
	}()

	go func(){
		anotherChannel <- "cow" 

	}()

	// just a random delay
	time.Sleep(time.Microsecond*20)

	select {
	case msgFromMyChannel := <- myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <- anotherChannel:
		fmt.Println(msgFromAnotherChannel)	
	}
	practice()
}



func practice(){

	thisChannel := make(chan string)

	go func(){
		// sending data
		thisChannel <- " "
	}()

	// recieving data
	mainF := <- thisChannel
	fmt.Println(mainF)
}