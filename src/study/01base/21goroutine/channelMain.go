package main

import (
	"fmt"
)

//通道是引用类型，通道类型的空值是nil。
func channelDemo() {
	ch := make(chan int)
	fmt.Println(ch) // <nil>
}

func channelDemo1(){
	ch := make(chan int)
	ch <- 10 // 把10发送到ch中
	x := <- ch // 从ch中接收值并赋值给变量x
	///<-ch       // 从ch中接收值，忽略结果
	println(x)
}

func channelDemo2(){
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

//使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
func channelDemo3() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

//有缓冲区
func channelDemo4(){
	ch := make(chan int,1)
	ch <- 10
	fmt.Printf("len:%d cap:%d \n",len(ch),cap(ch))
	x := <- ch
	fmt.Printf("%v \n",x)
}

func channelDemo5(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~10的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

func channelDemo6(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}