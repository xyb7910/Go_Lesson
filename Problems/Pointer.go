package main

import "fmt"

func main() {
	var a int = 10
	/*
		指针变量的定义格式：var var_name *var-type
		指针变量的使用：
				- 定义指针变量。
				- 为指针变量赋值。
				- 访问指针变量中指向地址的值。
	*/
	var ad *int
	ad = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ad)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ad)

	var pr *int
	fmt.Printf("ptr 的值为：\n", pr)

	//指针数组的使用
	arr := []int{1, 2, 3, 4}
	var ptr [4]*int

	var i int
	for i = 0; i < len(arr); i++ {
		ptr[i] = &arr[i] // 将整数地址赋给指针数组
	}

	for i = 0; i < len(arr); i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
