package main

import "fmt"

//整型
func main()  {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010   占位符%b表示二进制
	fmt.Printf("%o \n", a)  // 12 	 占位符%o表示八进制
	fmt.Printf("%x \n", a)  // a      占位符%x表示十六进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF

	//查看变量类型
	fmt.Printf("%T\n",c)

	//声明int8类型的变量
	i4 := int8(9) //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n",i4)
}
