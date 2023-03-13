package main

import "fmt"

//golang没有构造函数的概念，可以使用函数来模拟构造函数的功能

type Person2 struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person2, error) {
	if name == "" {
		return nil, fmt.Errorf("name不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age 不能小于0")
	}
	return &Person2{name: name, age: age}, nil
}

func main() {
	person, err := NewPerson("tom", 20)
	if err == nil {
		fmt.Printf("person:%v\n", *person)
		//fmt.Printf("*Person2:%v\n", *Person2)
	}
}
