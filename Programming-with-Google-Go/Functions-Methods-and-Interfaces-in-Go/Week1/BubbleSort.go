package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter Integeres upto 10: ")
	nums := make([]int, 0)
	for x:= 0; x < 10;x++{
		var i int
		fmt.Scanf("%d", &i)
		nums = append(nums, i)
	}
	fmt.Println("Sorted Slice:")
	BubbleSort(nums)
	fmt.Printf("%v", nums)

}

func Swap(slc []int , i int)  {
	temp := slc[i]
	slc[i] = slc[i+1]
	slc[i+1] = temp
}

func BubbleSort(slc []int) {
	n := len(slc)
	for i := 0; i < n-1; i++{
		for j := 0; j < n-i-1; j++{
			if slc[j] > slc[j+1]{
				Swap(slc, j)
			}
		}

	}

}
