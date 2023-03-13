package main

import "fmt"

func main() {

	//方法1 for循环+if判断
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Printf("data:%v\n", data)
		} else {
			break
		}
	}

	//方法2 for range
	//c2 := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c2 <- i
	//	}
	//	close(c2)
	//}()
	//
	//for v := range c2 {
	//	fmt.Printf("v:%v\n", v)
	//}

	//注意：如果通道关闭，读多写少，没有了就是默认值，例如，int就是0，如果没有关闭就会死锁
}
