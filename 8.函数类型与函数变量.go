package main

import "fmt"

//可以使用type关键字来定义一个函数类型，语法格式如下:
//type fun func(int,int) int
//上面语句定义了一个 fun 函数类型，它是一种函数类型，这种函数接收两个int类型的参数并返回一个类型的返回值

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//定义一个fun函数，把sum和max赋值给它 sum和max是逻辑运算，要和类型匹配得上才能赋值给一个新韩淑
type fun func(int, int) int

func main() {
	var f1 fun
	f1 = sum
	f2 := f1(3, 4)

	fmt.Printf("f2:%v\n", f2)
}
