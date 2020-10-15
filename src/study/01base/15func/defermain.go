package main

import "fmt"

func main() {
	deferDemo()
	//deferDemo1()
	//deferDemo2()
}

//defer语句
//defer语句能非常方便的处理资源释放问题
func deferDemo(){
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

//这块是懵逼的
func deferDemo1(){
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func calc4(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
func deferDemo2(){
	x := 1
	y := 2
	defer calc4("AA", x, calc4("A", x, y))
	x = 10
	defer calc4("BB", x, calc4("B", x, y))
	y = 20
	defer calc4("CC", x, calc4("C", x, y))
}
