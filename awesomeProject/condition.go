package main

import (
	"fmt"
	"time"
)

func main()  {
	num := 5
	if num < 2 {
		fmt.Println("Hii")
	} else {
		fmt.Println("Byee")
	}

	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("3")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None")
	}
	var today = time.Now()
	fmt.Println(today)
}
