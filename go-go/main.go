// // package main
// //
// // import "fmt"
// //
// //	func main() {
// //		//var str string = "this is an example of a string"
// //		//fmt.Printf("t/f does the string \"%s\", have contains %s?", str, "th")
// //		//fmt.Printf("%t\n", strings.ContainsAny(str, "th"))
// //		//fmt.Printf("%s\n", strings.Replace(str, "t", "T", -1))
// //		//fmt.Printf("%s\n", strings.ReplaceAll(str, "s", "S"))
// //		//var oldStr = "hi there! "
// //		//var newStr string
// //		//newStr = strings.Repeat(oldStr, 5)
// //		//fmt.Printf("%s\n", newStr)
// //		//strings.HasPrefix();strings.HasSuffix()字符串前缀和后缀
// //		//strings.ContainsAny()字符串包含
// //		//strings.Index();strings.LastIndex();strings.IndexRune()字符串第一个/最后一个/非ascii索引位置
// //		//strings.Replace()字符串替换
// //		//strings.Contains()统计字符串出现次数
// //		//strings.Repeat()重复字符串
// //		//strings.ToLower();strings.ToUpper字符串转小写转大写
// //		//strings.TrimSpace();stings.Trim(s, "cut")剔除字符串前后空格或指定字符串
// //		//strings.Fields(s)利用空白符分割字符串，返回slice
// //		//strings.Split(s, sep)利用指定字符分割字符串，返回slice
// //		//strings.Join(sl []string, sep string)用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：
// //		//strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数。strconv.Itoa(12)
// //		//strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串
// //		//strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
// //		//strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。
// //		//val, err = strconv.Atoi(s)
// //		//k := 6
// //		//switch k {
// //		//case 4:
// //		//	fmt.Println("was <= 4")
// //		//	fallthrough
// //		//case 5:
// //		//	fmt.Println("was <= 5")
// //		//	fallthrough
// //		//case 6:
// //		//	fmt.Println("was <= 6")
// //		//	fallthrough
// //		//case 7:
// //		//	fmt.Println("was <= 7")
// //		//	fallthrough
// //		//case 8:
// //		//	fmt.Println("was <= 8")
// //		//	fallthrough
// //		//default:
// //		//	fmt.Println("default case")
// //		//}
// //		//for i := 0; i < 5; i++ {
// //		//	var v int
// //		//	fmt.Printf("%d ", v)
// //		//	v = 5
// //		//}
// //		//for i := 0; i < 3; {
// //		//	fmt.Println("Value of i:", i)
// //		//}
// //		//for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
// //		//	s = i+1, j+1, s+"a" {
// //		//	fmt.Println("Value of i, j, s:", i, j, s)
// //		//}
// //		//LABEL1:
// //		//	for i := 0; i <= 5; i++ {
// //		//		for j := 0; j <= 5; j++ {
// //		//			if j == 4 {
// //		//				continue LABEL1
// //		//			}
// //		//			fmt.Printf("i is: %d, and j is: %d\n", i, j)
// //		//		}
// //		//	}
// //		//i := 0
// //		//for { //since there are no checks, this is an infinite loop
// //		//	if i >= 3 {
// //		//		break
// //		//	}
// //		//	//break out of this for loop when this condition is met
// //		//	fmt.Println("Value of i is:", i)
// //		//	i++
// //		//}
// //		//fmt.Println("A statement just after for loop.")
// //		//for i := 0; i < 7; i++ {
// //		//	if i%2 == 0 {
// //		//		continue
// //		//	}
// //		//	fmt.Println("Odd:", i)
// //		//}
// //		n := 0
// //		reply := &n
// //		Multiply(10, 5, reply)
// //		fmt.Println("Multiply:", *reply) // Multiply: 50
// //	}
// //
// // // this function changes reply:
// //
// //	func Multiply(a, b int, reply *int) {
// //		*reply = a * b
// //	}
// package main
//
// import "fmt"
//
// // 使用 defer 语句实现代码追踪，一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息，即可以提炼为下面两个函数：
// // func trace(s string)   { fmt.Println("entering:", s) }
// // func untrace(s string) { fmt.Println("leaving:", s) }
//
//	func trace(s string) string {
//		fmt.Println("entering:", s)
//		return s
//	}
//
//	func un(s string) {
//		fmt.Println("leaving:", s)
//	}
//
//	func a() {
//		defer un(trace("a"))
//		fmt.Println("in a")
//	}
//
//	func b() {
//		defer un(trace("b"))
//		fmt.Println("in b")
//		a()
//	}
//
//	func main() {
//		b()
//	}
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	//内置函数
//	//	close(); len() cap(); new() make(); copy() append(); panic() recover();	print() println(); complex() real() imag();
//	//for i := 0; i <= 10; i++ {
//	//	j, res := fibonacci(i)
//	//	fmt.Printf("fibonacci(%d) is : %d\n", j, res)
//	//}
//	//f := fib()
//	//for i := 0; i < 9; i++ {
//	//	println(i+2, f())
//	//}
//	//addBmp := MakeAddSuffix(".bmp")
//	//addJpeg := MakeAddSuffix(".jpeg")
//	//fmt.Println(addBmp("file"))
//	//fmt.Printf(addJpeg("file"))
//	//s1 := []byte{'p', 'o', 'e', 'm'}
//	//s2 := s1[2:]
//	//fmt.Printf("%s\n", s1) //poem
//	//fmt.Printf("%s\n", s2) //em
//	//s2[1] = 't'
//	//fmt.Printf("%s\n", s1) //poet
//	//fmt.Printf("%s\n", s2) //et
//
//	//bytes 包
//	//var buffer bytes.Buffer <==> var r *bytes.Buffer = new(bytes.Buffer)
//	/*	var buffer bytes.Buffer
//		for {
//			if s, ok := getNextString(); ok { //method getNextString() not shown here
//				buffer.WriteString(s)
//			} else {
//				break
//			}
//		}
//		fmt.Print(buffer.String(), "\n")
//	*/
//	TestFor()
//}
//func TestFor() {
//	//var slice1 []int = make([]int, 4)
//	//slice1[0] = 1
//	//slice1[1] = 2
//	//slice1[2] = 3
//	//slice1[3] = 4
//	//for i, value := range slice1 {
//	//	fmt.Printf("slice at %d is: %d\n", i, value)
//	//}
//	items := [...]int{10, 20, 30, 40, 50}
//	for i, _ := range items {
//		items[i] = items[i] * 2
//	}
//	for i, item := range items {
//		fmt.Printf("%d : %d\n", i, item)
//	}
//}
//
//func fibonacci(n int) (i, res int) {
//	i = n
//	if n <= 1 {
//		res = 1
//	} else {
//		_, res1 := fibonacci(n - 1)
//		_, res2 := fibonacci(n - 2)
//		res = res1 + res2
//	}
//	return i, res
//}
//
//// 闭包的斐波那契
//func fib() func() int {
//	a, b := 1, 1
//	return func() int {
//		a, b = b, a+b
//		return b
//	}
//}
//
//// 一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用：
//// 书写一个工厂函数而不是针对每种情况都书写一个函数。下面的函数演示了如何动态返回追加后缀的函数：
//func MakeAddSuffix(suffix string) func(string) string {
//	return func(name string) string {
//		if !strings.HasSuffix(name, suffix) {
//			return name + suffix
//		}
//
//		return name
//	}
//}
//
////addBmp := MakeAddSuffix(".bmp")
////addJpeg := MakeAddSuffix(".jpeg")
////fmt.Println(addBmp("file"))
////fmt.Printf(addJpeg("file"))

package main

func main() {

}
