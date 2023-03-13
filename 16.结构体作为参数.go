package main

//Go结构体可以像普通变量一样，作为函数的参数，传递给函数，这里分两种情况
//1.直接传递结构体，这是一个副本（拷贝），在函数内部不会改变外面结构体内容
//2.传递结构体指针，这时在函数内部，能够改变外部结构体内容

type Student struct {
	name string
	sex  string
}

func showStudent(student2 Student) {
	student2.name = "1"
	student2.sex = "男"
}
