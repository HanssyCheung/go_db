package main

import "fmt"

//数组定义语法
//var variable_name [size] variable_type
//variable_name:数组名称 size数组长度，必须是常量 variable_type：数组保存元素的类型

func main() {
	var a [3]int
	var s [2]string
	fmt.Printf("a:%T\n", a)
	fmt.Printf("s:%T\n", s)
}
