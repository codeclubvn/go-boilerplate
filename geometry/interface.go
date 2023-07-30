package main

// khai bao interface hinh hoc
// Tính đa hình: các phương thức của interface dù có tên giống nhau nhưng có cách sử dụng khác nhau
// Tính trừu tượng: các phương thức của interface không có nội dung, tùy vào từng struct sẽ có nội dung khác nhau
type Geometry interface {
	Area() float64
	Input()
}
