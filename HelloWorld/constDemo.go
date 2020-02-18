package main

import (
	"fmt"
	"unsafe"
)

//const可以作为枚举
const (
	l = "abc"
	m = len(l)
	n = unsafe.Sizeof(l)
)

//iota
/*iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/
const (
	aa = iota
	bb
	cc
	dd = "ha"
	ee
	ff = 100
	gg
	hh = iota
	ii
)

/*iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。

简单表述:

i=1：左移 0 位,不变仍为 1;
j=3：左移 1 位,变为二进制 110, 即 6;
k=3：左移 2 位,变为二进制 1100, 即 12;
l=3：左移 3 位,变为二进制 11000,即 24。
 */
const (
	i = 1 << iota
	j = 3 << iota
	k
	kk
)

func main()  {
	const LENGTH int  = 100
	const WIDTH int  = 5
	var area int
	const a, b, c  = 1, false, "str" //多重赋值

	area = LENGTH + WIDTH
	fmt.Println("面积为 : ", area)
	println("面积为 ：", area)
	println()
	println(a, b, c)

	println(l, m, n)

	println(aa, bb, cc, dd, ee, ff, gg, hh, ii)

	println(i, j, k, kk)
}