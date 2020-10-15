package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//goroutineDemo()
	//goroutineDemo2()
	//goroutineDemo3()
	//goroutineDemo4()

	//channelDemo()
	//channelDemo1()
	//channelDemo2()
	//channelDemo3()
	//channelDemo4()
	//channelDemo5()
	//channelDemo6()
	//workpoolDemo()
	//selectDemo()
	//lockDemo1()
	//lockDemo2()
	//lockDemo3()
	//lockDemo4()
	//onceDemo()
	//mapDemo()
	//mapDemo2()
	automicdemo()
}

//顺序执行
func goroutineDemo(){
	hello()
	fmt.Println("main goroutine done!")
}

//并发执行
func goroutineDemo2(){
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}

//使用了sync.WaitGroup来实现goroutine的同步
func goroutineDemo3(){
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello2(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束s
}

//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。 将逻辑核心数设为2，此时两个任务并行执行
func goroutineDemo4(){
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}



func hello() {
	fmt.Println("Hello Goroutine!")
}

//使用了sync.WaitGroup来实现goroutine的同步
var wg sync.WaitGroup

func hello2(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
