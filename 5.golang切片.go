package main

import "fmt"

//数组是固定长度，可以容纳相同数据类型元素的集合。当长度固定时就会不可避免的带来浪费内存和容量不够的情况
//js中的数组是可以自动扩容的，与go不同
//切片就是可变长度的数组，底层是使用数组实现的，增加了自动扩容功能。切片(Slice)是一个拥有相同类型元素的可变长度的序列

//切片语法
//声明一个切片和声明一个数组类似，只要不添加长度就可
// var identifier []type

//切片是引用类型，可以使用make函数来创建切片：
//var slice []type = make([]type,len)
//简写
//slice1 := make([]type,len)
//也可以指定容量，其中capacity为可选参数
//make([]T,length,capacity) //length指的是初始长度，capacity指的是初始容量

func main() {
	//初始化
	s := []int{1, 2, 3} //直接使用数组类似的方式不加size

	arr := [...]int{1, 2, 3}
	s1 := arr[:] //使用:符号将里数组变成切片

	fmt.Print(s, arr, s1)

	//也可以使数组部分切片话 左包含，右不包含
	s2 := arr[2:]
	fmt.Print(s2)

	//切片的添加和删除 append copy拷贝
	s3 := append(s1, 9, 10)
	s4 := append(s3[0:], s3[1:]...)
	fmt.Print(s3[0:])
	s5 := append(s3[:0], s3[1:]...)
	s6 := append(s3[0:2], s3[3:]...)
	fmt.Printf("s3:%v\n,s4:%v\n,s5:%v\n,s6:%v\n", s3, s4, s5, s6)
}
