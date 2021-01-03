package main

import "fmt"

var a = 9 	// this called as a package level variable.This can access in other functions in this file
			// var a = 10 => can't define two variables in same name here

func demo () {
	var a = 10
	fmt.Println("in demo ", a) // can't access this variable in other functions in this file
}

func main () {
	demo()
	fmt.Println("in main ",a)
}
