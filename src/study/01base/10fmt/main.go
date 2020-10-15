package main

import "fmt"

func main()  {
	//fmt占位符
	n := 100
	//查看类型
	fmt.Printf("%T\n",n)
	fmt.Printf("%v\n",n)
	fmt.Printf("%b\n",n)
	fmt.Printf("%d\n",n)
	fmt.Printf("%o\n",n)
	fmt.Printf("%x\n",n)
	s := "张三"
	fmt.Printf("字符串：%s\n",s)
	fmt.Printf("字符串：%v\n",s)
	fmt.Printf("字符串：%#v\n",s)
}
