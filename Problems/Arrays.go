package main

import "fmt"

func main() {
	/*
		数组的声明格式：
				var variable_name [SIZE] variable_type
		举例：
				var balance [10] float32
	*/

	var arr1 = [6]int{1, 2, 3, 4, 5, 6}
	arr2 := [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 6; i++ {
		fmt.Println(arr1[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Println(arr2[i])
	}
	arr3 := [...]string{"Alice", "Bob", "Tom", "Tim"}
	fmt.Println(len(arr3))
	/*
		访问数组元素:
			访问格式为：通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。
			var salary float32 = balance[9]
	*/
	for i := 0; i < 6; i++ {
		arr1[i] += 1
		fmt.Println(arr1[i])
	}

	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < len(balance2); j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	/*
		多维数组的使用：
			声明格式：
					var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
			举例：
					var threedim [5][10][4]int
	*/
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	/*
		初始化二维数组：
			通过大括号来初始值：
					a := [3][4]int{
						 {0, 1, 2, 3} ,    第一行索引为 0
						{4, 5, 6, 7} ,     第二行索引为 1
						{8, 9, 10, 11},    第三行索引为 2
					}
	*/
	/*
		访问二维数组：通过指定坐标来访问
			举例：val := a[2][3]
						或
				var value int = a[2][3]
	*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	/* 输出数组元素 */
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	rw1 := []string{"fish", "shark", "eel"}
	rw2 := []string{"bird"}
	rw3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, rw1)
	animals = append(animals, rw2)
	animals = append(animals, rw3)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}
