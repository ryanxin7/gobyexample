package main

import "fmt"


// 定义一个可变参数函数，接收任意数量的 `int` 类型参数
func sum(nums ...int)  {
	fmt.Println(nums,"")
	total := 0
	// 可变参数在函数内部被表示为切片，可以使用 range 遍历
	for _,num := range nums{
		total += num // 累加所有参数的值
	}
	fmt.Println(total)

}

func main()  {

	sum(1,2)
	sum(1,2,3)

	// 将切片作为参数传入可变参数函数
	nums := []int{1,2,3,4}
	// 使用 `slice...` 将切片解包为独立参数传递
	sum(nums...)
}