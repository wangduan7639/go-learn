package main

import "fmt"

func main()  {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值是：%d\n", ret)

	i, j := swap("Google", "Runoob")
	fmt.Println(i, j)
}

/**
函数
 */
func max(num1, num2 int) int  {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**
函数返回多个值
 */
func swap(x, y string) (string, string)  {
	return y, x
}