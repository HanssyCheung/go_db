package main

import "fmt"

//Go中数组的指针是指向数组中的每一个元素，而c的是指向数组的首地址

const MAX = 3

func main() {

	a := []int{1, 3, 5}
	var i int
	var ptr [MAX]*int
	fmt.Printf("ptr:%v\n", ptr) //[<nil><nil><nil>]
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] //整数地址赋值给指针数组
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i]) //*ptr[i]就是打印出相关指针的值了
	}
}
