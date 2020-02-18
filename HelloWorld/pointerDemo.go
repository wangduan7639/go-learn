package main

import "fmt"

func main()  {
	var a int = 20
	var ip *int

	ip = &a  //指针变量的存储地址
	fmt.Printf("a 变量的地址是：%x\n", &a )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量存储的指针地址：%x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值：%d\n", *ip )

	////空指针
	//var ptr *int
	//fmt.Printf("ptr 的值为 ：%x\n", ptr )

	//指向指针的指针
	var ptr *int
	var pptr **int

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
