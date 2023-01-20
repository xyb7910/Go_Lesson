package main

import "fmt"

func main() {
	/*
			功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
		定义切片：var identifier []type
		创建切片：
			使用make()函数来创建切片：
					var slice1 []type = make([]type, len)
					也可以简写为：
					slice1 := make([]type, len)
			也可以指定容量，其中 capacity 为可选参数。
					make([]T, length, capacity)
	*/
	var a = []int{1, 2, 3, 4, 5, 6}
	S0 := a[1:3] //startIndex:endIndex
	S1 := a[:5]
	S2 := a[3:]
	for i := 0; i < len(S0); i++ {
		fmt.Print(S0[i])
	}
	for i := 0; i < len(S1); i++ {
		fmt.Print(S1[i])
	}
	for i := 0; i < len(S2); i++ {
		fmt.Print(S2[i])
	}

	/*
		切片中的len()与cap（）
			 len() 方法获取长度
			 cap() 可以测量切片最长可以达到多少
	*/
	var numbers = make([]int, 5)
	printSlice(numbers)
	var num []int
	if num == nil {
		fmt.Printf("切片是空的")
	}

	/*
		append()和copy()函数：拷贝切片的 copy 方法和向切片追加新元素的 append 方法
	*/
	var numbers2 []int
	printSlice(numbers2)

	/* 允许追加空切片 */
	numbers2 = append(numbers2, 0)
	printSlice(numbers2)

	/* 向切片添加一个元素 */
	numbers2 = append(numbers2, 1)
	printSlice(numbers2)

	/* 同时添加多个元素 */
	numbers2 = append(numbers2, 2, 3, 4)
	printSlice(numbers2)

	numbers3 := make([]int, len(numbers2), (cap(numbers2) * 2))

	copy(numbers3, numbers2)
	printSlice(numbers3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
