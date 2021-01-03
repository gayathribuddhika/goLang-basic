package main

import "fmt"

func add(x int, y int) int {
	out := x + y
	return out
}

func calc(x int, y int) (int, int) {
	out1 := x + y
	out2 := x - y
	return out1, out2
}

//func calc(x int, y int) (out1, out2 int) {  // this also can be used
//	out1 = x + y
//	out2 = x - y
//	return
//}

func main () {
	//num1 := 1
	//num2 := 2
	//result := add(num1, num2)
	result1 := add(5, 4)
	fmt.Println(result1)

	result2, result3 := calc(500, 200)
	//fmt.Println(result2, result3) // output: 700 300
	fmt.Println(result2)
	fmt.Println(result3)
}
