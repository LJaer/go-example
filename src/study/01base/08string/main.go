package main

import (
	"fmt"
	"strings"
)


//字符串
//go语言中字符串是用双引号包裹的
//go语言中单引号包裹的是字符
func main() {
	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1)
	fmt.Println(s2)

	//字符串转义符
	//\r	回车符（返回行首）
	//\n	换行符（直接跳到下一行的同列位置）
	//\t	制表符
	//\'	单引号
	//\"	双引号
	//\\	反斜杠
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	//多行字符串
	//反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	s3 := `第一行
第二行
第三行`
	println()
	fmt.Println(s3)
	//len(str)	求长度
	fmt.Println(len(s3))
	//+或fmt.Sprintf	拼接字符串
	fmt.Println(s1+"123")
	var ss1 = fmt.Sprintf("Sprintf: %s%s","1","2")
	fmt.Println(ss1)
	fmt.Printf("Printf: %s%s \n","1","2")
	//strings.Split	分割
	var s4 = "1,2,3,4,5,1"
	fmt.Println(strings.Split(s4,","))
	//strings.contains	判断是否包含
	fmt.Println(strings.Contains(s4,"6"))
	//strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	fmt.Println("hasPrefix",strings.HasPrefix(s4,"1"))
	fmt.Println("hasSuffix:", strings.HasSuffix(s4,"5"))
	//strings.Index(),strings.LastIndex()	子串出现的位置
	fmt.Println("index:",strings.Index(s4,"1"));
	fmt.Println("lastIndex:",strings.LastIndex(s4,"1"));
	//strings.Join(a[]string, sep string)	join操作
	var arr = []string{"张三","李四"}
	fmt.Println(strings.Join(arr,"|"))

	//字符串
	s5 := "Hello 沙河!"
	fmt.Println(s5)

	//单独的字符、汉字、符号表示一个字符
	c1 := 'h'
	c2 := '1'
	c3 := '沙'
	fmt.Println(c1,c2,c3)

	//字节：1字节=8Bit(8个二进制位)
	//1个字符 'A' = 1个字节
	//1个utf8编码的汉字 '沙' = 一般占3个字节



}
