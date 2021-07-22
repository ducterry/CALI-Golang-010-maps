package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	//01. Khai báo Map
	var (
		myMap = map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
		}
	)
	// 02. In màn hình
	fmt.Println("MyMap Data is: \n")
	fmt.Println(myMap)
}
