package main

import (
	"fmt"
	"math"
)

//浮点型
func main() {
	//math.MaxFloat32 //float32
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	f1 := 1.23456
	fmt.Printf("%T\n",f1) //默认go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n",f2) //显示声明float32类型
	//f1 = f2 // float32类型的值不能直接赋值给float64类型的变量
}
