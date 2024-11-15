package main

import (
	"fmt"
	"time"
)


func main() {
	//基本的switch语句
	i := 2
	fmt.Println("Write ",i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//使用多个表达式的case语句，包含默认分支
    switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
    }

    //不带表达式的 switch 用于替代if/else

    t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

    //类型 switch 检查接口类型
    whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
whatAmI(true)
whatAmI(1)
whatAmI("hey")

}