package main

import (
	"fmt"
	"time"
)

//Timer只执行一次，TicKer可以周期的执行
func main() {
	//ticker := time.NewTicker(time.Second)
	//counter := 1
	//for _ = range ticker.C {
	//	fmt.Println("ticker 1")
	//	counter++
	//	if counter >= 5 {
	//		break
	//	}
	//}
	//ticker.Stop()

	chanInt1 := make(chan int) //创建一个通道，只要数据写入就会造成通道阻塞
	fmt.Printf("chanInt1:%p\n", &chanInt1)

	ticker2 := time.NewTicker(time.Second)
	println("111", &ticker2, &ticker2.C)
	go func() {
		for _ = range ticker2.C {
			select { //没有条件，会随机选择一个进行写入
			case chanInt1 <- 1:
			case chanInt1 <- 2:
			case chanInt1 <- 3:

			}
		}
	}()

	sum := 0
	fmt.Printf("chanInt1:%p\n", &chanInt1)
	for v := range chanInt1 { //哪些情况会解引用
		fmt.Printf("接收:%v\n", v)
		sum += v
		if sum >= 10 {
			fmt.Printf("sum:%v\n", sum)
		}
		break
	}
}
