package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	x      int64
	wg2    sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg2.Done()
}

//最终结果不是1w
func lockDemo1() {
	wg2.Add(2)
	go add()
	go add()
	wg2.Wait()
	fmt.Println(x)
}

func add2() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg2.Done()
}

//最终结果是1w
func lockDemo2() {
	wg2.Add(2)
	go add2()
	go add2()
	wg2.Wait()
	fmt.Println(x)
}

func lockDemo3(){
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func lockDemo4(){
	wg2.Add(1)
	go hello3() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg2.Wait()
}

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg2.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg2.Done()
}

func hello3() {
	defer wg2.Done()
	fmt.Println("Hello Goroutine!")
}

func sayHello(){
	fmt.Println("sayHello")
}

var sayHelloOnce sync.Once

func onceDemo(){
	sayHelloOnce.Do(sayHello)
}

//单例模式
type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func mapDemo(){
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var m2 = sync.Map{}

func mapDemo2(){
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}