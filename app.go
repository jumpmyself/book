package main

import (
	"book/model"
	"book/router"
)

// Start 这是一个启动器方法
func Start() {
	model.NewMysql()
	model.NewRdb()
	defer func() {
		model.Close()

	}()

	router.New()

}
