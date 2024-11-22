package main

import (
	"fmt"
)

func main(){
	//使用range 计算切片中所有数字的和。数组也可以这样使用
	nums := []int{2,3,4}
	sum := 0 //是初始化 sum 变量
	for _, nums := range nums{
		sum += nums
	}
	println("sum:", sum)

	//使用 range 获取数组或切片中的每个元素的索引和值

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}


	//使用 range 遍历map，返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range kvs  {
		fmt.Printf("%s -> %s\n", k,v)
	}

    //使用 range 遍历map ，仅返回建
    for k := range kvs {
		fmt.Println("key:", k)
	}

    //使用 range 迭代字符串，遍历 Unicode 码点
    for i, c := range "go" {
    	fmt.Println(i,c)
	}

}
