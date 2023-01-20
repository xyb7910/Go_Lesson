package main

import "fmt"

type Books struct {
	bookName string
	bookId   int
	data     string
	author   string
}

func main() {
	/*
		结构体定义格式：
			type struct_variable_type struct {
						   member definition
						   member definition
						   ...
						   member definition
						}
		使用结构体声明变量：
				variable_name := structure_variable_type {value1, value2...valuen}
				或
				variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
	*/
	// 创建一个新的结构体
	fmt.Println(Books{"C语言", 01, "2023-1-20", "yxc"})
	// 也可以使用 key => value 格式
	fmt.Println(Books{bookName: "C语言", bookId: 01, data: "2023-1-20", author: "yxc"})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{bookName: "Go语言", author: "yxc"})

	/*
		访问结构体成员格式：结构体.成员名
	*/
	var Book1 Books
	var Book2 Books
	/* book 1 描述 */
	Book1.bookName = "Go 语言"
	Book1.author = "yxc"
	Book1.bookId = 02
	Book1.data = "2023-1-1"

	/* book 2 描述 */
	Book2.bookName = "Python 教程"
	Book2.author = "ypb"
	Book2.bookId = 03
	Book2.data = "2023-1-23"

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 BookName : %s\n", Book1.bookName)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 BookId : %d\n", Book1.bookId)
	fmt.Printf("Book 1 data : %s\n", Book1.data)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 BookName : %s\n", Book2.bookName)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 bookId : %d\n", Book2.bookId)
	fmt.Printf("Book 2 data : %s\n", Book2.data)

	/*
		结构体指针的作用：var struct_pointer *Books，定义的指针变量可以存储结构体变量的地址。
		查看结构体变量地址，可以将 & 符号放置于结构体变量前：struct_pointer = &Book1
	*/
}
