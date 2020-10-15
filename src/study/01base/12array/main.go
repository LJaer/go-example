package main

import "fmt"

//数组
//var 数组变量名 [元素数量]T
func main() {
	//arrayDemo1()
	//arrayDemo2()
	//arrayDemo3()
	//arrayDemo4()
	//arrayDemo5()
	//arrayDemo6()
	//arrayDemo7()
	//arrayDemo8()
	//arrayDemo9()
	//arrayDemo10()
	arrayDemo11()
}

func arrayDemo1() {
	// 定义一个长度为3元素类型为int的数组a
	var a [3]int
	for k, v := range a {
		fmt.Printf("k: %d v: %d", k, v)
	}
}

func arrayDemo2() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 3}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}

//让编译器根据初始值的个数自行推断数组的长度
func arrayDemo3() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

//使用指定索引值的方式来初始化数组
func arrayDemo4(){
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}

//遍历数组
func arrayDemo5(){
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

//多维数组
func arrayDemo6(){
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}

//二维数组的遍历
func arrayDemo7(){
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

//注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
func arrayDemo8(){
	//支持的写法
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//不支持多维数组的内层使用...
	/*
	b := [3][...]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	*/
	fmt.Print(a)
}


func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func arrayDemo9() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

//求和
func arrayDemo10(){
	nums := [8]int{1, 3, 5, 7, 8}
	var sum = 0;
	for _,v := range nums {
		sum += v
	}
	print(sum)
}

//找出数组中和为指定值的两个元素的下标
//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func arrayDemo11(){
	sum := 8;
	nums := [5]int{1, 3, 5, 7, 8}
	res := [8]int{}
	count := 0
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums);j++{
			if i!=j && nums[i]+nums[j] == sum {
				if !arrayCheckRepeat(res,nums[i]) {
					res[count] = nums[i]
					res[count+1] = nums[j]
					count += 2
					println(i,j)
				}
			}
		}
	}
}

//判断数组里面是否已包含
func arrayCheckRepeat(nums [8]int,a int)(bool){
	for _,v := range nums{
		if v == a {
			return true
		}
	}
	return false
}