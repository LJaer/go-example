package main

//常量
//常量定义之后不能修改
//在程序运行期间不会改变
const pi = 3.1415926

//批量声明
const (
	httpOk = 200
	notFound = 404
)

//多个iota定义在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

//批量声明常量时，如果某一行声明后没有赋值
//默认和上一行一样
const (
	a1 = 100
	a2
	a3
)

//iota
//在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
//使用iota能简化定义，在定义枚举时很有用。
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

//使用_跳过某些值
const (
	b1 = iota //0
	b2        //1
	_
	b3        //3
)

//iota声明中间插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const c5 = iota //0

//定义数量级
//（这里的<<表示左移操作，
//1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
//同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main()  {
	println("a1:",a1);
	println("a2:",a2);
	println("a3:",a3);

	println("n1:",n1);
	println("n2:",n2);
	println("n3:",n3);
	println("n4:",n4);

	println("b1:",b1);
	println("b2:",b2);
	println("b3:",b3);

	println("c1:",c1);
	println("c2:",c2);
	println("c3:",c3);
	println("c4:",c4);
	println("c5:",c5);

	println("a:",a);
	println("b:",b);
	println("c:",c);
	println("d:",d);
	println("e:",e);
	println("f:",f);
}
