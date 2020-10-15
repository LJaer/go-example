package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

//结构体中字段大写开头表示可公开访问，小写表示私有(仅在定义当前结构体的包中可访问)
//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	//structDemo()
	//structDemo1()
	//structDemo2()
	//structDemo3()
	//structDemo4()
	//structDemo5()
	//structDemo6()
	//structDemo7()
	//structDemo8()
	//structDemo9()
	//structDemo10()
	//structDemo11()
	//structDemo12()
	//structDemo13()
	//structDemo14()
	//structDemo15()
	//structDemo16()
	//structDemo17()
	//structDemo18()
	//structDemo19()
	//structDemo20()
	//structDemo21()
	//structDemo22()
	//structDemo23()
	//structDemo24()
	//structDemo25()
	//structDemo26()
	structDemo27()
}

func structDemo() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}

type person struct {
	name string
	city string
	age  int8
}

type person1 struct {
	name, city string
	age        int8
}

func structDemo1() {
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
}

//匿名结构体
func structDemo2() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

func structDemo3() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person 指针地址
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
}

func structDemo4() {
	var p2 = new(person)
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}
}

//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
//p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。
func structDemo5() {
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
}

//没有初始化的结构体，其成员变量都是对应其类型的零值。
func structDemo6() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
}

//使用键值对初始化
func structDemo7() {
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
}

//对结构体指针进行键值对初始化
func structDemo8() {
	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
}

//没有指定初始值的字段的值就是该字段类型的零值
func structDemo9() {
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
}

//使用值的列表初始化
//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
//
//使用这种格式初始化时，需要注意：
//必须初始化结构体的所有字段。
//初始值的填充顺序必须与字段在结构体中的声明顺序一致。
//该方式不能和键值初始化方式混用。
func structDemo10() {
	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
}

//结构体占用一块连续的内存
func structDemo11() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	//fmt.Println(unsafe.Sizeof(n))  // 0
}

//空结构体是不占用空间的
func structDemo12() {
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0
}

type student struct {
	name string
	age  int
}

//一道有意思的面试题
//0xc00009ad78
//0xc00009ad60
//0xc00009ad48
//0xc0000a8020
//0xc0000a8020
//0xc0000a8020
//小王子 => 大王八
//娜扎 => 大王八
//大王八 => 大王八
func structDemo13() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	a := stus[0]
	b := stus[1]
	c := stus[2]
	println(&a)
	println(&b)
	println(&c)
	for _, stu := range stus {
		m[stu.name] = &stu
		println(&stu)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func structDemo14() {
	p9 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}
}

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型
func structDemo15() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

func structDemo16() {
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

//当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份
func structDemo17() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Println(p1.age) // 25
	p1.SetAge2(30)      // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 25
}

//MyInt 将int定义为自定义MyInt类型
type MyInt2 int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt2) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func structDemo18() {
	var m1 MyInt2
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}

//Person 结构体Person类型
//结构体的匿名字段
type Person2 struct {
	string
	int
}

func structDemo19() {
	p1 := Person2{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	fmt.Println(p1.string, p1.int) //北京 18
}

func structDemo20() {
	//Address 地址结构体
	type Address struct {
		Province string
		City     string
	}

	//User 用户结构体
	type User struct {
		Name    string
		Gender  string
		Address Address
	}

	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("user1=%#v\n", user1) //user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
}

func structDemo21() {
	//Address 地址结构体
	type Address struct {
		Province string
		City     string
	}

	//User 用户结构体
	type User struct {
		Name    string
		Gender  string
		Address //匿名字段
	}
	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
}

//嵌套结构体的字段名冲突
func structDemo22() {
	//Address 地址结构体
	type Address struct {
		Province   string
		City       string
		CreateTime string
	}

	//Email 邮箱结构体
	type Email struct {
		Account    string
		CreateTime string
	}

	//User 用户结构体
	type User struct {
		Name   string
		Gender string
		Address
		Email
	}

	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
}

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

//结构体的“继承”
func structDemo23() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

func structDemo24() {
	//Student 学生
	type Student struct {
		ID     int
		Gender string
		Name   string
	}

	//Class 班级
	type Class struct {
		Title    string
		Students []*Student
	}

	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	println()
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

//json序列化
func structDemo25() {
	//Student 学生
	type Student struct {
		ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
		Gender string //json序列化是默认使用字段名作为key
		name   string //私有不能被json包访问
	}

	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}

type Person3 struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person3) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func structDemo26() {
	p1 := Person3{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?
}

type Student struct {
	id    int
	name  string
	age   int
	score int
}

func newStudent(id, age, score int, name string) Student {
	return Student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

func structDemo27() {
	var class []Student //学生列表
	for i := 0; i < 10; i++ {
		cl := newStudent(i, i, i, "张")
		class = append(class, cl)
	}
	fmt.Printf("%v \n",class)
	//添加学生
	class = append(class,newStudent(11, 11, 11, "李"))
	fmt.Printf("%v \n",class)
	//编辑学生
	class[10].name = "赵"
	fmt.Printf("%v \n",class)
	//删除学生
	class = append(class[:10], class[11:]...)
	fmt.Printf("%v \n",class)
}
