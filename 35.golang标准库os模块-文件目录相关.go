package main

import (
	"fmt"
	"os"
)

//创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("f.Name():%v\n", f.Name())
	}
}

//创建目录
func createDir() {
	//创建单个目录
	//err := os.Mkdir("test", os.ModePerm) //os.ModePerm表示最高权限
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//}
	err := os.Mkdir("test/a/b", os.ModePerm) //os.ModePerm表示最高权限
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

//删除目录
func removeDir() {
	//err := os.RemoveAll("test")
	//if err != nil {
	//	fmt.Printf("err:%v\n", err)
	//}

	err := os.RemoveAll("test/a/b")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}

//获得工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("dir:%v\n", dir)
	}
}

func main() {
	createFile()
}
