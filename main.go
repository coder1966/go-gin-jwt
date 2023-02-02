package main

import (
	"go-gin-jwt/model"
	"go-gin-jwt/router"
)

func main() {
	// 引用数据库
	model.InitDB()

	router.InitRouter()
}
