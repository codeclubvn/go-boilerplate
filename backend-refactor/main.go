package main

import (
	"boilerplate/backend-refactor/config"
	"boilerplate/backend-refactor/router"
)

func main() {
	config.SetEnv()
	router.NewService()
}
