package main

import "fmt"

func main() {
	//ifDemo1()
	//ifDemo2()

	//forDemo()
	//forDemo2()
	//forDemo3()

	//rangeDemo1()
	//rangeDemo2()
	//rangeDemo3()
	//rangeDemo4()

	//switchDemo1()
	//testSwitch3()
	//switchDemo4()
	//switchDemo5()

	//gotoDemo1()
	//gotoDemo2()

	//breakDemo1()

	//continueDemo()

	jiujiuDemo()
}


func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func rangeDemo1() {
	s := []string{"1", "2", "3", "4", "5", "6"}
	for i, v := range s {
		fmt.Printf("i:%d v:%s \n", i, v)
	}
}

func rangeDemo2() {
	s := []string{"张", "三", "李", "四", "王", "赵"}
	for i, v := range s {
		fmt.Printf("i:%d v:%s \n", i, v)
	}
}

//一个汉字至少占3个字节
func rangeDemo3() {
	s := "张三李四王五"
	for i, v := range s {
		fmt.Printf("i:%d v:%c \n", i, v)
	}
}

//map
func rangeDemo4() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	for i, v := range countryCapitalMap {
		fmt.Printf("i:%s v:%s \n", i, v)
	}
}

func switchDemo1() {
	finger := 2
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo4() {
	age := 60
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

//break语句可以结束for、switch和select的代码块。
//
//break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，
//标签要求必须定义在对应的for、switch和 select的代码块上。
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

//99乘法表
/*
1*1
2*1 2*2
3*1 3*2 3*3
4*1 4*2 4*3 4*4
5*1 5*2 5*3 5*4 5*5
6*1 6*2 6*3 6*4 6*5 6*6
7*1 7*2 7*3 7*4 7*5 7*6 7*7
8*1 8*2 8*3 8*4 8*5 8*6 8*7 8*8
9*1 9*2 9*3 9*4 9*5 9*6 9*7 9*8 9*9
*/
func jiujiuDemo(){
	for i:=1;i<10;i++{
		for j:=1;j<10;j++{
			fmt.Printf("%v*%v ",i,j)
			if(i==j){
				fmt.Print("\n")
				break
			}
		}
	}
}