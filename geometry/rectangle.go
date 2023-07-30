package main

import "fmt"

// Khai báo struct cho hình chữ nhật
type Rectangle struct {
	length, width float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.length * rectangle.width
}

// nhập vào chiều dài và chiều rộng của hình chữ nhật
func (rectangle *Rectangle) Input() {
	fmt.Print("Nhập chiều dài: ")
	fmt.Scan(&rectangle.length)
	fmt.Print("Nhập chiều rộng: ")
	fmt.Scan(&rectangle.width)
}
