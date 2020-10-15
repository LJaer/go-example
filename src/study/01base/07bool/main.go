package main

import "fmt"

//布尔类型变量的默认值为false。
//Go 语言中不允许将整型强制转换为布尔型.
//布尔型无法参与数值运算，也无法与其他类型进行转换。
func main()  {
	b := true
	var isOK bool //默认是false
	fmt.Printf("%T\n",isOK)
	fmt.Printf("%T value:%v\n",b,b)
}
