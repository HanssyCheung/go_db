package main

import "fmt"

//跳到指定位置开始执行，break是跳出当前层
func goto1() {
	a := 0
	if a == 1 {
		goto LABEL1
	} else {
		fmt.Printf("other...")
	}
LABEL1:
	fmt.Printf("next...")
}

func main() {
	goto1()
}
