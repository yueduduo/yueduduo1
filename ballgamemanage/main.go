package main
// 程序启动包
import (
	"mygo/model"
	"mygo/router"
)

func main() {
	model.Init()
	router.Login()
}
