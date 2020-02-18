package main

import "fmt"

func main()  {

	//算术运算符
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n",c)

	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n",c)

	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n",c)

	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n",c)

	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n",c)

	c++
	fmt.Printf("第六行 - c 的值为 %d\n",c)

	c--
	fmt.Printf("第七行 - c 的值为 %d\n",c)

	//关系运算符
	if a == b {
		fmt.Printf("第八行 - a 等于 b\n")
	} else {
		fmt.Printf("第八行 - a 不等于 b\n")
	}

	if a < b {
		fmt.Printf("第九行 - a 小于 b\n")
	} else {
		fmt.Printf("第九行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第十行 - a 大于 b\n")
	} else {
		fmt.Printf("第十行 - a 不大于 b\n")
	}

	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第十一行 - a 小于等于 b\n")
	}

	if a >= b {
		fmt.Printf("第十一行 - a 大于等于 b\n")
	}
	
	//逻辑运算符
	var i bool = true
	var j bool = false
	if i && j {
		fmt.Printf("i && j - 条件为 true\n" )
	}

	if i || j {
		fmt.Printf("i || j - 条件为 true\n" )
	}

	i = false
	j = true
	if i && j {
		fmt.Printf("i && j - 条件为 true\n" )
	} else {
		fmt.Printf("i && j - 条件为 false\n" )
	}

	if !(i && j) {
		fmt.Printf("!(i && j) - 条件为 true\n")
	}

	otherOperation()
}

/**
其他运算符
 */
func otherOperation()  {
	//位运算符
	var a uint = 60  /* 60 = 0011 1100 */
	var b uint = 13  /* 13 = 0000 1101 */
	var c uint

	fmt.Printf("a & b 的值为 %d\n", a & b)
	fmt.Printf("a | b 的值为 %d\n", a | b)
	fmt.Printf("a ^ b 的值为 %d\n", a ^ b)
	fmt.Printf("a << 2  的值为 %d\n", a << 2 )
	fmt.Printf("a >> 2 的值为 %d\n", a >> 2)

	//赋值运算符
	fmt.Printf("赋值运算符\n")
	c =  a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

	c +=  a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

	c -=  a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

	c *=  a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

	c /=  a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

	c  = 200;

	c <<=  2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

	c >>=  2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

	c &=  2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

	c ^=  2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

	c |=  2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
}