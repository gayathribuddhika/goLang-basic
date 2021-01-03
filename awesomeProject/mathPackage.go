package main

import (
	"fmt"
	"math"
)

func main () {
	fmt.Println("Gayathri")
	var num float64 = 24
	var result = math.Sqrt(num)
	var intResult1 = math.Ceil(result)
	var intResult2 = math.Floor(result)
	//fmt.Println("Squireroot of", num , "is:",result)
	fmt.Printf("The output is %g. Thank You", result)
	fmt.Println()
	fmt.Printf("The output is %.2f. Thank You", result)
	fmt.Println()
	fmt.Printf("The output is %.2f.", intResult1)
	fmt.Println()
	fmt.Printf("The output is %.2f.", intResult2)
}

