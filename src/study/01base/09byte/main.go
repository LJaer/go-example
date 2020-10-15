package main

import (
	"fmt"
	"math"
)

//Go 语言的字符有以下两种：
//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
//rune类型，代表一个 UTF-8字符。
//当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
func main()  {
	a1 := '中'
	b1 := 'x'
	println(a1)
	println(b1)

	traversalString()

	changeString()

	//类型转换
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

	//双引号是
	d1 := "红"
	d2 := '红' // rune(int32)
	fmt.Printf("d1: %T d2: %T\n",d1,d2)

	d3 := "H" //string
	d4 := 'h' //int32
	fmt.Printf("d3: %T d4: %T\n",d3,d4)
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune 从字符串中拿出具体的字节
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

//修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2) //把字符串强制转换成一个 rune 切片
	runeS2[0] = '红'
	fmt.Println(string(runeS2)) //把 rune 切片强制转换为字符串
}
