package main

import (
	"fmt"
	"sync"
	"time"
)

//除了使用channel实现同步之外，还可以使用Mutex互斥锁的方式实现同步

var m int = 100
var lock sync.Mutex
var wt sync.WaitGroup

func addLock() {
	defer wt.Done()
	lock.Lock()
	m += 1
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("m++:%v\n", m)
	lock.Unlock()

}

func subLock() {
	defer wt.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	m -= 1
	fmt.Printf("m++:%v\n", m)
	lock.Unlock()

}

func main() {

	for i := 0; i < 1000; i++ {
		wt.Add(1)
		go addLock()
		wt.Add(1)
		go subLock()
	}
	wt.Wait()
	fmt.Printf("end m: %v\n", m)
}
