package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello world 中国"
	b := []byte(s)
	b[5] = ','
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)

	r := []rune(s)
	r[6] = '中'
	r[7] = '国'
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", string(r))

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
