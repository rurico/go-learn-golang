package main

import "fmt"

func main() {
	// var a string = "Runoob"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)
	fun2()
}

func fun1() {
	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
}

func fun2() {
	var a *int
	var b []int
	var c map[string]int
	var d chan int
	// var e func(string) int
	var f error // error 是接口
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// fmt.Println(e)
	fmt.Println(f)
}
