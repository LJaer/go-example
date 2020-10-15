package main

import (
	"errors"
	"fmt"
	"strings"
)

//使用type关键字来定义一个函数类型
type calculation func(int, int) int

func main()  {
	//sum := intSum(1,2)
	//println(sum)

	//sayHello()

	/*ret1 := intSum3()
	ret2 := intSum3(10)
	ret3 := intSum3(10, 20)
	ret4 := intSum3(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60*/

	/*ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160*/

	/*a,b := calc(5,1)
	fmt.Println(a,b)*/

	/*a,b := calc2(5,1)
	fmt.Println(a,b)*/

	/*a := someFunc("")
	fmt.Println(a)*/

	//funcDemo1()
	//funcDemo2()
	//funcDemo3()

	/*a,err := do("+")
	sum := a(1,2)
	print(sum,err)*/

	//funcDemo4()
	//funcDemo5()
	//funcDemo6()
}

func intSum(x int, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello 沙河")
}

//类型简写
//函数的参数中如果相邻变量的类型相同，则可以省略类型
func intSum2(x, y int) int {
	return x + y
}

//可变参数
func intSum3(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：
func intSum4(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

//多个返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return nil
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func funcDemo1(){
	var c calculation
	c = add
	a := c(1,2)
	println(a)
}

//声明函数类型的变量并且为该变量赋值
func funcDemo2(){
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}

//高阶函数
//函数作为参数
func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func funcDemo3() {
	ret2 := calc1(10, 20, add)
	fmt.Println(ret2) //30
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//匿名函数
func funcDemo4(){
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

//闭包
func funcDemo5(){
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//闭包
func funcDemo6(){
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

func calc3(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

//闭包
func funcDemo7(){
	f1, f2 := calc3(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}






