package main

import "fmt"

func main() {
	//变量的学习
	//定义单个变量: var identifier type
	var a string = "ypb"
	fmt.Println(a)
	//定义多个变量: var identifier1, identifier2 type
	var b, c int = 1, 2
	fmt.Println(b, c)

	//变量的声明
	//1.指定变量类型，如果没有初始化，则变量默认为零值。
	// var v_name v_type
	// v_name = v_value
	var d int
	fmt.Println(d)
	d = 8
	//2.根据值自行判定变量类型
	var e = true
	fmt.Println(e)
	//3.如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
	//var val int
	//val := 8
	//这样会报错
	//可以直接使用:=
	val := 8
	//等价于
	// var val int
	// val = 8
	fmt.Println(val)

	//多个变量的声明
	/*
		//类型相同多个变量, 非全局变量
		var vname1, vname2, vname3 type
		vname1, vname2, vname3 = v1, v2, v3

		var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

		vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


		// 这种因式分解关键字的写法一般用于声明全局变量
		var (
		    vname1 v_type1
		    vname2 v_type2
		)
	*/
	var val1, val2, val3 int
	val1, val2, val3 = 1, 2, 3
	fmt.Println(val1, val2, val3)
	var v1, v2, v3 = 1, 2, 3
	fmt.Println(v1, v2, v3)
	v4, v5, v6 := 1, 2, 3
	fmt.Println(v4, v5, v6)

	var (
		a1 int
		a2 int
	)
	a1, a2 = 1, 2
	fmt.Println(a1, a2)
	a1, a2 = a2, a1
	fmt.Println(a1, a2)
}
