package main

import (
	"fmt"
	//"slices"
	//Go 1.17 不支持 slices 包， slices 是 Go 1.21 中新增的标准库功能。
)


func main(){
    //声明了一个切片 s，但未初始化。
	var s []string
	fmt.Println("uninit:", s,s ==nil, len(s) == 0)

	//创建一个长度为 3 的字符串切片。
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:",len(s),"cap:",cap(s))


	// 使用 s[索引] 方式设置和获取切片的元素值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:",s)
	fmt.Println("get:", s[2])

   //使用内置函数 len 获取切片的长度。
   fmt.Println("len:", len(s))

	//append 函数用于向切片添加元素，返回一个新的切片。

	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("apd:",s)

	//复制切片
	//copy(dst, src) 将 src 切片的内容复制到 dst 切片。
	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:", c)

	//通过索引范围进行截取
	//获取从索引 low 到 high-1 的子切片。
	l := s[2:5]
	fmt.Println("sl1:",l)

	//获取从切片开头到 high-1 的子切片
	l = s[:5]
	fmt.Println("sl2:",l)

	//获取从索引 low 到切片末尾的子切片。
	l = s[2:]
    fmt.Println("sl3:",l)

   //使用字面量初始化切片
   t := []string{"g","h","i"}
   fmt.Println("dcl:",t)

   //比较切片,判断两个切片是否相等
   //t2 := []string{"g","h","i"}
   //	if slices.Equal(t, t2) {
   //	fmt.Println("t == t2")
   //}

   //多维切片
   twoD := make([][]int,3)
   for i :=0; i < 3; i++ {
   	innerLen := i +1
   	twoD[i] = make([]int, innerLen)
	   for j := 0; j < innerLen; j++ {
	   	twoD[i][j] = i + j
	   }
	}
	fmt.Println("2d:",twoD)

}
