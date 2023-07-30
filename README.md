# Golang boiler-plate
- Mục đích của source code này là để giúp các bạn có thể nhanh chóng tìm hiểu về golang và có thể viết được một ứng dụng backend đơn giản.
- Source code này gồm 3 phần:
    - geometry: bài tập về hình học để hiểu cơ bản về golang
    - backend-basic: viết server backend đơn giản trong một file gồm các API và connect database
    - backend-refactor: tái cấu trúc code backend-basic
## need to install before run
- [golang](https://golang.org/doc/install)
- [docker](https://docs.docker.com/get-docker/)
- postgresql
- dbeaver or pgadmin
- postman
- goland or vscode

## source code structure
```
├── README.md
├── geometry # bài tập về hình học
├── backend-basic # viết API và connect database
├── backend-refactor # tái cấu trúc code backend-basic
    ├── api # api mà gọi đến service khác
    ├── route # api mà frontend sẽ gọi
    ├── handler # xử lý các request và trả về response
    ├── service # logic chính của ứng dụng
    ├── repository # tương tác với database
    ├── config # các file config
        ├── app # config app
        ├── db # config database
        ├── env # chứa các biến môi trường
    ├── middleware # các hàm trung gian, xử lý yêu cầu trước khi nó tới handler
    ├── model # lưu các struct
    ├── util # chứa các file utils
        ├── constant.go # chứa các constant
        ├── error.go # chứa các error
        ├── response.go # chứa các response
        |── util.go # chứa các hàm tiện ích
        ├── pointer.go # chứa các hàm kiểm tra pointer
    ├── main.go # file chạy chính
    ├── go.mod # file quản lý các package
```

## package used
- github.com/gin-gonic/gin
- gorm.io/gorm
- gorm.io/driver/postgres
- github.com/caarlos0/env
- errors
- context
- fmt

## how to run
- clone project
- run command `go run main.go` to run project
- open postman and import file `postman_collection.json` to test API