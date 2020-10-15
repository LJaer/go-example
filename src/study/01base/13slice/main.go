package main

import (
	"fmt"
	"sort"
)

//切片slice
//因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。

//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
func main() {
	//sliceDemo1()
	//sliceDemo2()
	//sliceDemo3()
	//sliceDemo4()
	//sliceDemo5()
	//sliceDemo6()
	//sliceDemo7()
	//sliceDemo8()
	//sliceDemo9()
	//sliceDemo10()
	//sliceDemo11()
	//sliceDemo12()
	//sliceDemo13()
	//sliceDemo14()
	//sliceDemo15()
	sliceDemo16()
}

//这个求和函数只能接受[3]int类型，其他的都不支持。
func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//数组a中已经有三个元素了，我们不能再继续往数组a中添加新元素了。
func arrayDemo1() {
	a := [3]int{1, 2, 3}
	fmt.Print(a)
}

func sliceDemo1() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

func sliceDemo2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	b := a[2:] // 等同于 a[2:len(a)]
	c := a[:3] // 等同于 a[0:3]
	d := a[:]  // 等同于 a[0:len(a)]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", b, len(b), cap(b))
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", c, len(c), cap(c))
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", d, len(d), cap(d))
}

func sliceDemo3() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

func sliceDemo4() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

func sliceDemo5() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}

func sliceDemo6() {
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", s1, len(s1), cap(s1))
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", s2, len(s2), cap(s2))
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", s3, len(s3), cap(s3))
}

//切片的赋值拷贝
func sliceDemo7() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}

//切片遍历
func sliceDemo8(){
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

//append()方法为切片添加元素
//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。
func sliceDemo9(){
	var s []int
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...)    // [1 2 3 4 5 6 7]
}

//append()添加元素和切片扩容
//append()函数将元素追加到切片的最后并返回该切片。
//切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
func sliceDemo10(){
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

//append()函数还支持一次性追加多个元素。
func sliceDemo11(){
	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a := []string{"成都", "重庆"}
	citySlice = append(citySlice, a...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
}

//由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。
func sliceDemo12(){
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}

//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
func sliceDemo13(){
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}

//从切片中删除元素
func sliceDemo14(){
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}

//练习题
func sliceDemo15(){
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

//排序
func sliceDemo16(){
	a := []int{3, 7, 8, 9, 1}
	//正序
	sort.Ints(a)
	fmt.Println(a)
	//逆序
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}
