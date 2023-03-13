package main

import (
	"fmt"
	"os"
)

//打开关闭文件
func openCloseFile() {
	//只能读
	f, _ := os.Open("a.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())

	//根据第二个参数，可以读写或者创建
	f2, _ := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2.Name(): %v\n", f2.Name())

	err := f.Close()
	fmt.Printf("err: %v\n", err)
	err2 := f.Close()
	fmt.Printf("err2: %v\n", err2)
}

//创建文件
func createFile1() {
	//等价于，OpenFile(name,O_RDWR|O_TRUNC,0666)
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
}
