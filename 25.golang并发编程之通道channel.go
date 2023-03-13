package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Go提供了一种称为通道的机制，用于在goroutine之间共享数据。当goroutine执行并发活动时，需要在goroutine之间共享
	资源或数据，通道冲到goroutine之间的管道并提供一种机制来保证同步交换

	需要在声明通道时指定数据类型。我们可以共享内置，命名，结构和引用类型的值和指针。数据在通道上传递：在任何给定时间只有
	一个goroutine可以访问数据项：因此按照设计不会发生数据竞争

	根据数据交换的行为，有两中类型的通道：无缓冲通道和缓冲通道。无缓冲通道用于执行goroutine之间的同步通信，而缓冲通道
	用于执行异步通信。无缓冲通道保证在发送和接收发生的瞬间执行goroutine之间的交换。缓冲通道没有这样的保证
*/

//通道由make函数创建，该函数指定chan关键字和通道的元素类型
//有无缓冲通道的创建
//Unbuffered := make(chan int) //整型无缓冲通道
//buffered := make(chan int,10) //整型有缓冲通道

//使用内置函数make创建有无缓冲通道。make的第一个参数需要关键字chan，然后是通道允许交换的数据类型

//这是将值发送到通道的代码块需要使用的<-运算符
//语法
//goroutine1 := make(chan string,5) //字符串缓冲通道
//goroutine1 <- "Australia" //通过通道发送字符串

//接收通道值的代码块
//data := <- goroutine1 //从通道接收字符串
//<- 运算符附加到通道变量(goroutine)的左侧，可以接收来自通道的值

//无缓冲通道
/*
	在接收到任何值之前都没有能力保存它。在这种通道中，发送和接收goroutine在任何发送或接收操作
	完成之前的同一时刻都准备就绪。如果两个goroutine没有在同一时刻准备好，则通道会让执行其各自
	发送或者接收操作的goroutine首先等待。同步是通道上发送和接收之间交互的基础。没有另一个就不
	可能发生
*/

//缓冲通道
/*
	在缓冲通道中，有能力在接收到一个或多个值之前保存它们。在这种类型的通道中，不要强制goroutine在
	同一时刻准备好执行发送和接收。当发送或接收阻塞时也有不同的条件。只有当通道中没有要接收的值时，接收
	才会阻塞。仅当没有可用缓冲区来放置正在发送的值时，发送才会阻塞
*/

//通道的发送和接收特性
/*
	1.对于一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
	2.发送操作和接收操作中对元素值的处理都是不可分割的
	3.发送操作在完全完成前会被阻塞。接收操作也是如此
*/

//创建int类型通道，只能传入int类型值
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}

func main() {
	//从通道中接收
	defer close(values)
	go send()
	println("wait...")
	value := <-values
	fmt.Printf("receive:%v\n", value)
	println("end...")
}
