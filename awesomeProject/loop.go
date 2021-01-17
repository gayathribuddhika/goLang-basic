package main

import "fmt"

func main () {
	x :=1
	for x<=5 {
		fmt.Println(x)
		x++
	}

	for i :=1; i<=5; i++ {
		fmt.Println("Gayathri ", i)
	}

	for x := 0; x<100; x++ {
		if x!= 0 && x%3 == 0 && x%7 == 0 {
			fmt.Println(x)
			continue  // print all the numbers that condition true
		}
	}
	fmt.Println(" ")

	for x := 0; x<100; x++ {
		if x!= 0 && x%3 == 0 && x%7 == 0 {
			fmt.Println(x)
			break // print only the first number that condition true
		}
	}
}
