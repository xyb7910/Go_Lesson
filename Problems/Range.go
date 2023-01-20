package main

import "fmt"

func main() {
	var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range pow {
		fmt.Printf("第%d个数是：%d\n", i, v)
	}
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}
