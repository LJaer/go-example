package main

import "fmt"

//go语言推荐使用驼峰命名
//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "张三"
	age = 18
	isOk = true
	// Go 语言中非全局变量声明后必须使用，不使用就编译不过去
	var heihei string
	//%s：占位符 使用 name 这个变量的值去替换占位符
	fmt.Printf("name:%s", name)
	fmt.Println(heihei)
	//类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明，只能在函数里用
	s3 := "哈哈哈"
	fmt.Println(s3)
	//s3 := "哈哈哈" //同一个作用域({}) 中不能重复声明同名的变量

	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

func foo() (int, string) {
	return 10, "Q1mi"
}
