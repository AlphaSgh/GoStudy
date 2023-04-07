package slice

import "fmt"

func CopyAppend() {
	//func append(s[]T, x ...T) []T 其中 append() 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；
	//追加的元素必须和原切片的元素是同类型。
	//如果 s 的容量不足以存储新增元素，append() 会分配新的切片来保证已有切片元素和新增元素的存储。
	//因此，返回的切片可能已经指向一个不同的相关数组了。append() 方法总是返回成功，除非系统内存耗尽了。
	//如果你想将切片 y 追加到切片 x 后面，只要将第二个参数扩展成一个列表即可：x = append(x, y...)

	//func copy(dst, src []T) int 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，
	//并且返回拷贝的元素个数。源地址和目标地址可能会有重叠。拷贝个数是 src 和 dst 的长度最小值。
	//如果 src 是字符串那么元素类型就是 byte。如果你还想继续使用 src，在拷贝结束后执行 src = dst
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

func MagnifySlice(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	fmt.Println("The length of ns  is:", len(ns))
	copy(ns, s)
	fmt.Println(ns)
	s = ns
	fmt.Println(s)
	fmt.Println("The length of s after enlarging is:", len(s))
	return s
}
