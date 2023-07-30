package main

import "fmt"

// khai báo struct cho tam giác
type Triangle struct {
	bottomEdge, height float64
}

// tính diện tích tam giác
func (triangle Triangle) Area() float64 {
	return triangle.bottomEdge * triangle.height / 2
}

// nhập vào cạnh đáy và chiều cao của tam giác
func (triangle *Triangle) Input() {
	fmt.Print("Nhập cạnh đáy: ")
	fmt.Scan(&triangle.bottomEdge)
	fmt.Print("Nhập chiều cao: ")
	fmt.Scan(&triangle.height)
}
