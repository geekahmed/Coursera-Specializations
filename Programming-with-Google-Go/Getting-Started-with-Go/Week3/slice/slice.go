package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sliceCntr := 0
	mySlc := make([]int,3)
	for true{
		fmt.Printf("Enter an Integer (X to Quit): ")
		scanner.Scan()
		line := scanner.Text()
		if line == "X" {
			break 
		}
		iii, _ := strconv.Atoi(line)

		if sliceCntr < 3{
			mySlc[sliceCntr] = iii
			sliceCntr++
		} else {
			mySlc = append(mySlc, iii)
		}

	}
	sort.Ints (mySlc)
	fmt.Printf("\nYour Sorted Slice: ")
	printSlice(mySlc)

}

func printSlice(intSlice	 []int) {
	fmt.Printf("len=%d cap=%d\n%v\n", len(intSlice), cap(intSlice), intSlice)
}
