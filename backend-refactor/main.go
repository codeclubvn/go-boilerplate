package main

import (
	"boilerplate/backend-refactor/config"
	"boilerplate/backend-refactor/route"
)

func main() {
	config.SetEnv()
	app := route.NewService()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
