package main

import "fmt"

func fact(n int) int {
	// 基准条件：当 n == 0 时，阶乘为 1
	if n == 0 {
		return 1
	}
	return  n * fact(n-1)
	// 递归步骤：n * (n-1)!
}

func main()  {
	fmt.Println(fact(7))

	// 提前声明变量 fib，其类型为 func(int) int
	var fib func(n int) int
	// 给 fib 赋值为一个匿名函数
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// 在函数体内通过 fib 调用自身
		return fib(n-1) + fib(n-2)

	}
	fmt.Println(fib(7))
}