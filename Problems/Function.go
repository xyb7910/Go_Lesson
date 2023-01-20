package main

import "fmt"

func main() {
	/*
		Go语言函数声明格式：
			func function_name( [parameter list] ) [return_types] {
		   				函数体
				}
		举例：以max函数为例，该函数传入两个整型参数 num1 和 num2，并返回这两个参数的最大值：
			func max(num1, num2 int) int {
				var result int
				if (num1 > num2) {
				result = num1
			} else {
				result = num2
			}
				return result
			}
	*/
	//值传递：在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	//默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	var a int = 100
	var b int = 200

	fmt.Printf("交换前a的值为：%d\n", a)
	fmt.Printf("交换前b的值为：%d\n", b)

	swap1(&a, &b)

	fmt.Printf("交换后a的值为：%d\n", a)
	fmt.Printf("交换后b的值为：%d\n", b)

	//引用传递:在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
}

func swap(num1, num2 int) (int, int) {
	return num2, num1
}

func add(num1, num2 int) int {
	var result int
	result = num1 + num2
	return result
}

func swap1(num1 *int, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}
