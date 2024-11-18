package main

import (
	"fmt"
)

func main()  {
	//创建空 Map
	m := make(map[string]int)

	//设置键值对
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:",m)


	//获取键对应的值
	v1 := m["k1"]
	fmt.Println("v1:",v1)

	//获取不存在的键
	v3 := m["k3"]
	fmt.Println("v3:",v3)

	//获取 Map 的长度
	fmt.Println("len:",len(m))

	//删除键值对
	delete(m,"k2")
	fmt.Println("map:",m)

	//清空 Map
	//方法1
	m = make(map[string]int)

	fmt.Println("After clear:", m)



	m["k4"] = 100
	fmt.Println("map:", m)

	//方法2
	fmt.Println("Before clear:", m)


	for k := range m {
		delete(m,k)
	}

	fmt.Println("After clear:", m)

	//检查键是否存在
	_, prs := m["k2"]
	fmt.Println("prs:",prs)

    //初始化并声明 Map

    n := map[string]int{"foo": 1, "bar":2}
    fmt.Println("map:",n)

	//比较两个 Map 是否相等
	//n2 := map[string]int{"foo":1,"bar":2}
	//if map.Equal(n,n2) {
	//	fmt.Println("n == n2")
	//}

}
