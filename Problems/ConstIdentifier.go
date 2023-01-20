package main

import "fmt"

func main() {
	/*
			常量的使用：
		只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
			格式：
		const identifier [type] = value  ，我们可以省略变量的类型说明符，让编译器自己判断
		- 显式类型定义： const b string = 'abc'
		- 隐式类型定义： const b = 'abc'

		多个常量的定义：
			格式：
		const c_name1, c_name2 = value1, value2
	*/
	const hour int = 24 //显式
	const minutes = 60  //隐式
	const length, height, width = 1, 2, 3
	var v int
	v = length * height * width
	fmt.Println(v)
	/*
		常量还可以使用枚举：
				const {
					a = 1
					b = 2
					c = 3
				}
		iota:是一个可以被编译器修改的常量。
		iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
				const (
					a = iota  //iota = 0
					b = iota  //iota = 1
					c = iota  //iota = 2
				)
	*/
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		q = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", q)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
