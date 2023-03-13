package main

import (
	"fmt"
	"sync/atomic"
)

//atomic提供的原子操作能够确保任意时刻都只有一个goroutine对变量进行操作，善用atomic能够避免程序中出现大量的锁操作

/* atomic常见操作有：
-增减
-载入 read
-比较并交换cas
-交换
-储存 write
*/

func main() {
	//增减操作
	//func AddInt32(addr *int32,delta int32)(new int32)
	//这里有个细节，像 AddUint32 针对无符号整型的操作，它实现减法的操作略微复杂一些，可以利用计算机补码的规则，把减法变成加法。
	//
	//以 AddUint32 为例：
	//atomic.AddUint32(&i, ^uint32(i-1))
	//载入操作
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i:%v\n", i)
	//储存
	atomic.LoadInt32(&i)
	atomic.StoreInt32(&i, 200)

	//cas 是上面操作的根基
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("i:%v\n", i)
}
