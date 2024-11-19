package main

import "fmt"

//定义一个函数，接收两个 int 类型的参数，返回它们的和
func plus(a int, b int) int{
	return a + b
}
// 如果参数类型相同，可以省略前两个参数的类型，只在最后一个参数上写类型
func plusPlus(a,b,c int) int  {
	return a + b + c
}

func main()  {
	// 调用函数 plus 并打印结果
	res := plus(1,2)
	fmt.Println("1+2=", res)

	// 调用函数 plusPlus 并打印结果
	res2 := plusPlus(1,2,3)
	fmt.Println("1+2+3=", res2)

}
