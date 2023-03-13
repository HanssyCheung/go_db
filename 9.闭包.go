package main

/*
	闭包可以理解成定义在一个函数内部的函数。在本质上，闭包是将函数内部和函数外部连起来的桥梁。或者说是函数和其
	引用环境的组合体

	闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
*/

//例子：返回一个函数
func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func add1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	//作为变量保存
	var f = add()
	println(f, f(10))
	println(f, f(20)) //变量逃逸，变量从栈里跑到堆里面了

	//作为参数保存
	f1 := add1(10)
	println(f1(20))
}
