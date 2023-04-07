package main

import (
	"fmt"
	"the_way_to_go/slice"
)

func main() {
	//slice.CopyAppend()
	s1 := []int{1, 2, 3}
	s1 = slice.MagnifySlice(s1, 3)
	fmt.Println("The length of s1 after enlarging is:", len(s1))
}
