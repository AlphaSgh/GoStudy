package main

import "fmt"

func pointer() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	//不能获取字面量或常量的地址
	//const i = 5
	//ptr := &i //error: cannot take the address of i
	//ptr2 := &10 //error: cannot take the address of 10

	//对一个空指针的反向引用是不合法的，并且会使程序崩溃：
	//var p *int = nil
	//*p = 0
}
