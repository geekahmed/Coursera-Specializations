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
	mySlc := make([]int,3)
	sliceCntr := len(mySlc) - 1
	for true{
		fmt.Printf("Enter an Integer (X to Quit): ")
		scanner.Scan()
		line := scanner.Text()
		if line == "X" {
			break
		}
		iii, _ := strconv.Atoi(line)

		if sliceCntr < 0{
			mySlc = append(mySlc, iii)
		} else {
			mySlc[sliceCntr] = iii
		}
		sort.Ints (mySlc)
		fmt.Printf("\nYour Sorted Slice: ")
		printSlice(mySlc)
		sliceCntr--
	}

	sort.Ints (mySlc)
	fmt.Printf("\nYour Sorted Slice: ")
	printSlice(mySlc)

}

func printSlice(intSlice	 []int) {
	fmt.Printf("%v\n", intSlice)
}
