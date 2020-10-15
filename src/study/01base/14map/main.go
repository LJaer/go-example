package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	//mapDemo1()
	//mapDemo2()
	//mapDemo3()
	//mapDemo4()
	//mapDemo5()
	//mapDemo6()
	//mapDemo7()
	//mapDemo8()
	//mapDemo9()
	mapDemo10()
}

func mapDemo1() {
	scoreMap := make(map[string]int, 10)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of a:%T\n", scoreMap)
}

//在声明的时候填充元素
func mapDemo2() {
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo)
}

//判断某个键是否存在
func mapDemo3() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

//遍历
//注意： 遍历map时的元素顺序与添加键值对的顺序无关。
func mapDemo4() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//只输出key
	for k := range scoreMap {
		fmt.Println(k)
	}
	//只输出value
	for _, v := range scoreMap {
		fmt.Println(v)
	}
}

//删除键值对
func mapDemo5() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap, "小明") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

//按照指定顺序遍历map
func mapDemo6() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string,0,200)
	for k := range scoreMap{
		keys = append(keys, k)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

//元素为map类型的切片
func mapDemo7()  {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	fmt.Println(mapSlice[0])
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

//值为切片类型的map
func mapDemo8()  {
	sliceMap := make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string,0,2)
	}
	value = append(value,"北京","上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

//统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func mapDemo9(){
	str := "how do you do"
	words := strings.Split(str," ")
	fmt.Printf("v:%v \n",words)
	wordMap := make(map[string]int,10)
	for _,v := range words{
		_,ok := wordMap[v]
		if !ok {
			wordMap[v] = 1
		}else{
			wordMap[v] += 1
		}
	}
	fmt.Print(wordMap["do"])
}

//观察输出结果
func mapDemo10(){
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //[1 2 3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s) //[1 3]
	fmt.Printf("%+v\n", m["q1mi"]) //[1 3 3]
}
