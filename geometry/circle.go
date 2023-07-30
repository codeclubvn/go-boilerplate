package main

import "fmt"

// khai báo struct cho hình tròn
type Circle struct {
	radius float64
}

// tính diện tích hình tròn
func (circle Circle) Area() float64 {
	return circle.radius * circle.radius * 3.14
}

// nhập vào bán kính của hình tròn
func (circle *Circle) Input() {
	fmt.Print("Nhập bán kính: ")
	fmt.Scan(&circle.radius)
}
