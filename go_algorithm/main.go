package main

import (
	"fmt"
	"go_algorithm/sort"
)

func main() {
	//quicksort
	//arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//fmt.Println(sort.QuickSort(arr))
	//bubblesort
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(sort.BubbleSort(arr))
	//fmt.Println(sort.GetMax(arr))
	fmt.Println(sort.GetMax(arr))
}
