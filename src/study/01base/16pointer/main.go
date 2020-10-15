package main

import "fmt"

func main() {
	//pointerDemo1()
	//pointerDemo2()
	//pointerDemo3()
	//pointerDemo4()
	//pointerDemo5()
	//pointerDemo6()
	pointerDemo7()
}

func pointerDemo1() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}

func pointerDemo2(){
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func pointerDemo3(){
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}

func pointerDemo4(){
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}

func pointerDemo5(){
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

func pointerDemo6(){
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}

func pointerDemo7(){
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}