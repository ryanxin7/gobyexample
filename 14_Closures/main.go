package main

import "fmt"

//定义一个函数 `intSeq`，返回值是一个匿名函数
// 返回的匿名函数会形成闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i ++ // 匿名函数中使用并更新 i 的值
		return i
	}
}

func main()  {
	nextInt := intSeq()
	// 调用 intSeq，将返回的匿名函数赋值给 nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 创建一个新的闭包函数，变量 i 将重新初始化
	newInts := intSeq()
	fmt.Println(newInts())
}

