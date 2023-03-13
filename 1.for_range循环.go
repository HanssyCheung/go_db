package main

import "fmt"

//Go语言可以使用for range遍历数组、切片、字符串、map以及通道（channel)。通过for range遍历的返回有以下规律
//1.数组、切片、字符串返回索引和值
//map返回键和值
//3.通道（channel)只返回通道内的值
func f() {
	var a = [...]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i:%d,v:%v\n", i, v)
	}
}

func main() {
	f()
}
