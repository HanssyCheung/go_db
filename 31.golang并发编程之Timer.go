package main

import (
	"fmt"
	"time"
)

//Timer顾名思义，就是定时器的意思，内部也是通过channel来实现的

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)

	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	//如果只是想单纯的等待，可以使用time.Sleep来实现
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("2s后")

	time.Sleep(time.Second * 2)
	fmt.Println("再次两秒后")

	<-time.After(time.Second * 2) //time.After函数的返回值是chan Time
	fmt.Println("又一次2s后")

	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("timer3 expired")
	}()

	stop := timer3.Stop() //停止计时器
	//阻止timer事件发生，当该函数执行后，timer计时器停止，相应的事件不再执行
	if stop {
		fmt.Println("Timer 3 stopped")
	}
	println("before")
	timer4 := time.NewTimer(time.Second * 5) //原来设置5s
	timer4.Reset(time.Second * 2)            //重新设置时间，即修改NewTimer的时间
	<-timer4.C
}
