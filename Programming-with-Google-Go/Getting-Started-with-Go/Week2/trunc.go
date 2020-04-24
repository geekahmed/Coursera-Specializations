package main

import "fmt"

func main() {
	fmt.Printf("Enter a Floating Point Number: ")
	var num float32
	fmt.Scan(&num)
	numTrun := int32(num)

	fmt.Printf("Truncated Number: %d", numTrun)

}
