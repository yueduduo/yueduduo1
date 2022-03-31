package main

// 程序启动包
import (
	"mygo/model"
	"mygo/router"
	
	_"github.com/swaggo/gin-swagger"
	_"github.com/swaggo/gin-swagger/swaggerFiles"
	_"mygo/docs"
)

// @title 球赛管理系统
// @version 最后版未完成
// @description swagger学习文档，这条非必须

// @host 127.0.0.1:8080
// @BasePath 


func main() {
	model.Init()
	router.RoutesSet()
}
