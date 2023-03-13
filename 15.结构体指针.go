package main

import "fmt"

//结构体指针和普通的变量指针相同

func test21() {
	type Person struct {
		id   int
		name string
		age  int
	}

	tom := Person{
		id:   101,
		name: "tom",
		age:  20,
	}

	var p_person *Person
	p_person = &tom
	fmt.Printf("tom:%v\n", tom)
	fmt.Printf("p_person:%p\n", p_person)
	fmt.Printf("*p_person:%v\n", *p_person)

	//可以使用new关键字创建结构体指针，得到的是具体结构体的地址
	var p_person1 = new(Person)
	(*p_person1).id = 1
	p_person1.name = "mike"
	println(p_person1) //结构体的指针可以直接使用点引用访问内部变量，尽管变量为地址,Go语言是可以省略的

}

func main() {
	test21()
}
