package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo biến Map sử dụng make
	var (
		myMap = make(map[string]int)
	)

	// 02. In màn hình
	if myMap == nil {
		fmt.Println("MyMap is Nil")
	} else {
		fmt.Println("MyMap not Nil")
	}

	// 03. Gán Key-Value cho Map
	myMap["one"] = 1
	myMap["two"] = 2

	fmt.Println(myMap)

}
