package main

import (
	"fmt"
)

func loops(){

	// for i:=0;i<5;i++ {
	// 	fmt.Println(i)
	// }

	// for i:=0; i<5; i++ {
	// 	for j:=0; j<5; j++{
	// 		fmt.Println("*")
	// 	}
	// }

	// var intArr [4]int32

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])
	// fmt.Println(&intArr[3])

	// arr := [3]int{1,2,3}
	// fmt.Println(arr)

	// sliceArr := []int{2,42,6,7,4}
	// sliceArr = append(sliceArr,7)
	// fmt.Println(sliceArr)

	myMap := make(map[string]int)
	myMap["age"]=20
	myMap["gpa"]=7
	myMap["heigh"]=174
	myMap["body count"]=0

	// fmt.Println(myMap["age"])

	// value, exists := myMap["body count"]
	// if exists {
	// 	fmt.Println("body count exists : " , value)
	// }else{
	// 	fmt.Println("doesn't exist")
	// }

	// delete(myMap,"body count")
	// fmt.Println("After deletion : ", myMap)

	// for name:=range myMap{
	// 	fmt.Printf("Name : %v\n", name)
	// }

	// myMap2 := make(map[string]int)

	// myMap2["muawiyah"]=22
	// myMap2["serena"]=35
	// myMap2["lewis"]=44

	// for name, age:= range myMap2{
	// 	fmt.Printf("Name : %v, Age: %v\n", name,age)
	// }

	// for i:=10; i>=0 ; i-- {
	// 	fmt.Println(i)
	// }
	n := 5
	for i:=1 ; i<=n; i++ {
		for j:=1; j<=n-i; j++{
			fmt.Print(" ")
		}
		
		for k:=1; k <= i ; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	var emoji rune = '😊'
	fmt.Println("Emoji : ", string(emoji))

}