package router

// 这个包是写客户端的请求路径与请求方法, 是main程序直接调用执行
import (
	"mygo/handler"

	"mygo/model"//仅在登录时用一下

	"github.com/gin-gonic/gin"
)

var Login = func() {
	// 登录// 0.登录, 若成功则判断用户role，再选反过去重选路由
	eng := gin.Default()
	eng.POST("/login", func(c *gin.Context) {
		UserName := c.Query("username")
		response, Role := model.Login(UserName)
		if Role != -1 {
			c.JSON(200, response)
			switch Role {
			case 1:
				Role1Route()
			case 2:
				Role2Route()
			case 3:
				Role3Route()
			}
		}else{
			c.JSON(400, response)
		}	
	})
	eng.Run(":8088")
}

// 分用户级别设置路由
var Role3Route = func() {
	eng := gin.Default()
	// 登记创建players和teams信息并返回到客户端反馈
	eng.POST("/teams", handler.ChinTeam)
	eng.POST("/players", handler.ChinPlayer)
	// 创建比赛
	eng.POST("/games", handler.CreateGame)
	// 创建用户与预约
	eng.POST("/user", handler.CreateUser)
	eng.POST("/user/appoint", handler.CreateAppoint)

	// 查询players和teams信息localhost:8080/players?playerid=1
	eng.GET("/players", handler.GetPlayerById)
	eng.GET("/teams", handler.GetTeamById)

	// 查询球赛信息
	eng.GET("/games", handler.GetAGames)
	// 查询用户预约
	eng.GET("/user/appoint", handler.GetAppoint)

	// 更改player的newteamid
	eng.PUT("/players/update", handler.UpdatePTeam)
	// 更新修改团队的信息
	eng.PUT("/teams/update/name", handler.UpdateTeamName)
	eng.PUT("/teams/update/logo", handler.UpdateTeamLogo)
	eng.PUT("/teams/update/info", handler.UpdateTeamInfo)
	// 更改用户权限
	eng.PUT("/user", handler.UpdateUserRight)

	eng.Run(":8080")
}

var Role2Route = func() {
	eng := gin.Default()
	// 登记创建players和teams信息并返回到客户端反馈
	eng.POST("/teams", handler.ChinTeam)
	eng.POST("/players", handler.ChinPlayer)
	// 创建比赛
	eng.POST("/games", handler.CreateGame)
	// 创建比赛预约
	eng.POST("/user/appoint", handler.CreateAppoint)

	// 查询players和teams信息localhost:8080/players?playerid=1
	eng.GET("/players", handler.GetPlayerById)
	eng.GET("/teams", handler.GetTeamById)

	// 查询球赛信息
	eng.GET("/games", handler.GetAGames)
	// 查询用户预约
	eng.GET("/user/appoint", handler.GetAppoint)

	// 更改player的newteamid
	eng.PUT("/players/update", handler.UpdatePTeam)
	// 更新修改团队的信息
	eng.PUT("/teams/update/name", handler.UpdateTeamName)
	eng.PUT("/teams/update/logo", handler.UpdateTeamLogo)
	eng.PUT("/teams/update/info", handler.UpdateTeamInfo)

	eng.Run(":8080")
}

var Role1Route = func() {
	eng := gin.Default()
	// 创建比赛预约
	eng.POST("/user/appoint", handler.CreateAppoint)

	// 查询players和teams信息localhost:8080/players?playerid=1
	eng.GET("/players", handler.GetPlayerById)
	eng.GET("/teams", handler.GetTeamById)

	// 查询球赛信息
	eng.GET("/games", handler.GetAGames)
	// 查询用户预约
	eng.GET("/user/appoint", handler.GetAppoint)

	eng.Run(":8080")
}


// 各个用户预约各个比赛，最后每个比赛的总预约数，不是靠mysql自己生成（不是要你设计好建的表），
// 用update语句，一个个增加

// 用三级用户更改用户的Role值来达到控制权限的目的
