package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup


func main() {
	fmt.Printf("Enter the size of the array: ")
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("Enter %d integers: ", n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	slice1 := arr[:n / 4]
	slice2 := arr[n / 4:n / 2]
	slice3 := arr[n /2: 3 * n / 4]
	slice4 := arr[3 * n / 4:]

	fmt.Println("Partitioned Slices", slice1, slice2, slice3, slice4)
	wg.Add(4)
	sortArrs(slice1)
	sortArrs(slice2)
	sortArrs(slice3)
	sortArrs(slice4)
	wg.Wait()

	newSlice := mergeAndSort(slice1, slice2, slice3, slice4)

	fmt.Println("Final Sorted Slice ", newSlice)

}

func sortArrs (unsortedArr []int) []int{
	sort.Ints(unsortedArr)
	fmt.Println(unsortedArr)
	wg.Done()
	return unsortedArr
}

func mergeAndSort(slice1 []int, slice2 []int, slice3 []int, slice4 []int) []int {
	newSlice := make([]int, 0)
	newSlice = append(slice1, slice2...)
	newSlice = append(newSlice, slice3...)
	newSlice = append(newSlice, slice4...)
	sort.Ints(newSlice)
	return newSlice
}
