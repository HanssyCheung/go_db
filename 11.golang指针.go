package main

/*
	Go语言中的函数传参都值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。
	传递数据使用指针，而无须拷贝数据。
	类型指针不能进行偏移和运算
	Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）
*/

/*
	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量
	进行取地址操作。Go语言中的值类型（int、float、string、array、struct）都有对应的指针类型，如*int
	、*int64、*string等
*/

//语法
//var var_name *var_type
//var_name:指针变量名 var_type: 指针类型 *用于指定一个变量是作为指针

func main() {
	var ip *int //nil
	i := 100
	ip = &i //0xc000047f68
	println(ip, *ip)
}
