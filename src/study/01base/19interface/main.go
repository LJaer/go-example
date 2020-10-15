package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Cat struct{}

type Dog struct{}

func main() {
	//interfaceDemo2()
	//interfaceDemo3()
	//interfaceDemo4()
	//interfaceDemo5()
	//interfaceDemo6()
	//interfaceDemo7()
	//interfaceDemo8()
	//interfaceDemo10()
	interfaceDemo11()
}

func interfaceDemo2() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}

func interfaceDemo3() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}

func interfaceDemo4() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func interfaceDemo5() {
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}

func interfaceDemo6() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}

//空接口
func interfaceDemo7() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v p:%p\n", x, x, &x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v p:%p\n", x, x, &x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v p:%p\n", x, x, &x)
}

func interfaceDemo8() {
	var x interface{}
	show(x)
}

func interfaceDemo9() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

func interfaceDemo10(){
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
	println(w)
}

//断言
func interfaceDemo11(){
	var x interface{}
	x = "Hello "
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

type cat struct{}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

// Sayer 接口
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

type car struct {
	brand string
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
