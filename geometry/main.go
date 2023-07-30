package main

import (
	"fmt"
)

func main() {
	var rectangle Geometry = &Rectangle{} // Tính kế thừa: struct Rectangle kế thừa interface Geometry
	rectangle.Input()
	fmt.Printf("Diện tích hình chữ nhật là: %.2f\n", rectangle.Area())

	var triangle Geometry = &Triangle{}
	triangle.Input()
	fmt.Printf("Diện tích tam giác là: %.2f\n", triangle.Area())

	var circle Geometry = &Circle{}
	circle.Input()
	fmt.Printf("Diện tích hình tròn là: %.2f\n", circle.Area())
	//fmt.Printf("Chu vi hình tròn là: %.2f\n", circle.ChuVi()) // phương thức ChuVi() không được khai báo ở interface Geometry

}
