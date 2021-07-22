package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo 1 Map
	var (
		myMap map[string]int
	)

	// 02. In màn hình
	fmt.Println("MyMapp = ", myMap)
	if myMap == nil {
		fmt.Println("MyMap is Nil")
	}

}
