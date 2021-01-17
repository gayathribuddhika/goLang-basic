package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"  // used to convert string to other formats
)

func  main()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name: ")
	scanner.Scan() 							// default the value that we entered get as a string here
	input1 := scanner.Text()
	
	fmt.Println("Enter your age: ")
	scanner.Scan() 							// default the value that we entered get as a string here
	input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	
	fmt.Printf ("You entered your name as %q ", input1)
	fmt.Println()
	fmt.Printf ("Your age is %d ", input2)
}
