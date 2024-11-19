package main

import "fmt"

//函数返回多个值
func vals()(int,int)  {
	return 3,7
}

func main()  {
	//使用多重赋值接收返回值
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//忽略不需要的返回值
	_,c := vals()
	fmt.Println(c)
}