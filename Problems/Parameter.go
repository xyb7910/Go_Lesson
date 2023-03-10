package main

import "fmt"

// 声明全局变量g，全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
var g int = 10

func main() {
	/*
		函数中的参数：
				·函数内定义的变量称为局部变量，作用域只在函数体内，参数和返回值变量也是局部变量。
				·函数外定义的变量称为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
				·函数定义中的变量称为形式参数，形式参数会作为函数的局部变量来使用。
	*/
	//声明局部变量
	var a, b, c int
	//初始化局部变量
	a, b, c = 1, 2, 3
	//声明一个与全局变量相同的局部变量g，此时局部变量会被优先考虑
	var g int = 20
	fmt.Println(a, b, c)
	fmt.Println(g)
}
