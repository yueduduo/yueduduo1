package router

// 这个包是写客户端的请求路径与请求方法, 是main程序直接调用执行
import (
	"mygo/handler"

	"github.com/gin-gonic/gin"

	
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var RoutesSet = func() {

	eng := gin.Default()

	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// eng.POST("/login", handler.Login)

	// 登记创建players和teams信息并返回到客户端反馈
	eng.POST("/teams", handler.ChinTeam)
	eng.POST("/players", handler.ChinPlayer)
	// 创建比赛
	eng.POST("/games", handler.CreateGame)
	// 创建用户与预约
	eng.POST("/users", handler.CreateUser)
	eng.POST("/users/appoint", handler.CreateAppoint)

	// 查询players和teams信息localhost:8080/players?playerid=1
	eng.GET("/players", handler.GetPlayerById)
	eng.GET("/teams", handler.GetTeamById)
	// 查询球赛信息
	eng.GET("/games", handler.GetAGames)
	// 查询用户预约
	eng.GET("/users/appoint", handler.GetAppoint)

	// 更改player的newteamid
	eng.PUT("/players", handler.UpdatePTeam)
	// 更新修改团队的信息
	eng.PUT("/teams/name", handler.UpdateTeamName)
	eng.PUT("/teams/logo", handler.UpdateTeamLogo)
	eng.PUT("/teams/info", handler.UpdateTeamInfo)
	// 更改用户权限
	eng.PUT("/user", handler.UpdateUserRight)

	eng.Run(":8080")
}
