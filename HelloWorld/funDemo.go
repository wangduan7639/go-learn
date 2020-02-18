package main

import (
	"fmt"
	"math"
	)

func main()  {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值是：%d\n", ret)

	i, j := swap("Google", "Runoob")
	fmt.Println(i, j)

	//函数作为实参
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 Circle
	c1.radius = 10.0
	fmt.Println("圆的面积 = ", c1.getArea())
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

/**
闭包
 */
func getSequence() func() int  {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
语言函数方法
 */
//定义结构体
type Circle struct {
	radius float64
}

//该 method 属于Circle类型对象中的方法
func (c Circle) getArea() float64  {
	return 3.14 * c.radius * c.radius
}