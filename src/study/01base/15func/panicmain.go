package main

import (
	"fmt"
	"strings"
)

func main() {
	//panicDemo()
	//panicDem2()
	demo()
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcB2() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func panicDemo(){
	funcA()
	funcB()
	funcC()
}

func panicDem2(){
	funcA()
	funcB2()
	funcC()
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (x int){
	for _,v := range users {
		if strings.Contains(v,"e") || strings.Contains(v,"E") {
			coins -= 1
		}else if strings.Contains(v,"i") || strings.Contains(v,"I") {
			coins -= 2
		}else if strings.Contains(v,"o") || strings.Contains(v,"O") {
			coins -= 3
		}else if strings.Contains(v,"u") || strings.Contains(v,"U") {
			coins -= 4
		}
	}
	return coins
}

func demo(){
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
