package main

import "fmt"

func main()  {
	//创建数组，默认数组的值为零
	var a [5]int
	fmt.Println("emp:",a) //空的

   //设置和获取数组元素
   a[4] = 100
   fmt.Println("set:",a)
   fmt.Println("set",a[4])

  //获取数组长度
  //使用内置行数len获取数组长度
  fmt.Println("len:", len(a))


   //声明并初始化数组

   b := [5]int{1, 2, 3, 4,5}
   fmt.Println("dcl:", b) //declaration

   //自动推断数组长度
   b2 := [...]int{1, 2, 3, 4, 5}
   //b2 := [5]int{1, 2, 3, 4, 5} 等于手动推断5
   fmt.Println("dcl:", b2)


   //指定索引初始化
   b = [...]int{100, 3: 400,500}
   fmt.Println("idx:",b)
   //使用索引语法初始化数组：3: 400 表示第 4 个元素（索引为 3）的值为 400。

   //二维数组
   var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++  {
			twoD[i][j] = i + j
    //[2][3]int 是 一个二维数组 的声明，它表示一个 2 行 3 列 的数组。
		}
	}
	fmt.Println("2d:", twoD)


	//一次性初始化二维数组
	twoD2 := [2][3]int{
		{1,2,3},
		{1,2,3},
	}
	fmt.Println("2d2: ", twoD2)

}