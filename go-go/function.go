/*
Go 里面有三种类型的函数：

普通的带有名字的函数
匿名函数或者lambda函数（参考 第 6.8 节）
方法（Methods，参考 第 10.6 节）
除了 main()、init() 函数外，其它所有类型的函数都可以有参数与返回值。函数参数、返回值以及它们的类型被统称为函数签名。
在函数调用时，像切片 (slice)、字典 (map)、接口 (interface)、通道 (channel) 这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。
如果一个函数需要返回四到五个值，我们可以传递一个切片给函数（如果返回值具有相同类型）或者是传递一个结构体（如果返回值具有不同的类型）
*/
package main

//使用结构接受多个类型的变长参数
/*
type Options struct {
	par1 type1,
	par2 type2,
	...
}
F1(a, b, Options {par1:val1, par2:val2})
*/
//如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数
//一般而言我们会使用一个 for-range 循环以及 switch 结构对每个参数的类型进行判断：
/*
func typecheck(..,..,values … interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: …
			case float: …
			case string: …
			case bool: …
			default: …
		}
	}
}
*/
