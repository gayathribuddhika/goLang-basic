package main

import "fmt"

func main() {
	//fmt.Println("Hello Gayathri")
	//fmt.Println("Welcome to golang")

	//var x int = 10
	//var y  int = 20
	//var result = x + y

	//var x, y int = 10, 20        // can use like this also
	//var result = x + y

	result := 50  // this is same as var result = 50
	fmt.Println("The value of x + y is ",result)

	var num = 2
	num = 10 // can assign different values for num when declare using var

	const num2 = 5
	// num2 = 10 cannot assign value for num2 when declare using const

	fmt.Print("The value of x + y is ",num)
	fmt.Print(num2)
}
