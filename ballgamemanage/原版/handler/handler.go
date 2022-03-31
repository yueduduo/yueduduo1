package handler

// 这个包是用gin写接送与处理客户端信息、调取发送服务端信息
import (
	"fmt"
	"mygo/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 0.登录并判断
func Login(c *gin.Context) {

}

// 1.各种创建************************************************************************************
// 1.1接收players和teams信息然后登记创建并返回到客户端反馈
// 模型如下ChinPlayer函数：
// 		c.ShouldBind(NewPlayer)接收post的内容，其中NewPlayer是中转站，
// 		NewPlayer.Checkin()是存储数据库
//		c.JSON(..., ...)是返回反馈
func ChinPlayer(c *gin.Context) {
	NewPlayer := &model.Player{}                    // 空的东西
	if err := c.ShouldBind(NewPlayer); err != nil { //这里的c.ShouldBind是把客户端提交的(json)存到后端服务器内存的NewPlayer变量里 // 前端的信息应该先是在c这个*gin.Context里
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	NewPlayer.Checkin()           //这个是存到将内存变量(结构体)的值(内存暂时存的)存到数据库中
	c.JSON(200, "new player创建成功") //服务器给客户端返回信息
}

func ChinTeam(c *gin.Context) {
	NewTeam := &model.Team{}
	if err := c.ShouldBind(NewTeam); /*这句放if语句前面也行，后面才是*/ err == nil {
		NewTeam.Checkin() // 先识别客户端的信息，再存进数据库
		c.JSON(200, "new team创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// 1.2创建球赛信息
func CreateGame(c *gin.Context) {
	NewGame := &model.Game{}
	if err := c.ShouldBind(NewGame); err == nil {
		NewGame.CreateGame()
		c.JSON(200, "new game创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// 1.3创建用户
func CreateUser(c *gin.Context) {
	NewUser := &model.User{}
	if err := c.ShouldBind(NewUser); err == nil {
		NewUser.Register()
		c.JSON(200, "new user 创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// 1.4创建user的预约
func CreateAppoint(c *gin.Context) {
	UserAppointName := c.Query("username") + "appoints"
	NewAppoint := &model.UserAppoint{}

	if err := c.ShouldBind(NewAppoint); err == nil {
		fmt.Println(NewAppoint)
		NewAppoint.CreateAppoint(UserAppointName) //还传入指定的表名，预约比赛的不同用户放在对应不同的appoint表，就是插入的时候判断插到哪张表
		// 还应有个对games表对应game行的appoint列update加1, fangzai了model.CreateAppoint()中
		c.JSON(200, "new appoint创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// 2.各种查询*************************************************************************************
//2.1查询players和teams信息并返回到客户端
func GetPlayerById(c *gin.Context) {
	playerid, _ := strconv.Atoi(c.Query("playerid")) //1.接收前端get的query内容
	response := model.GetPlayerById(playerid)        //2.在数据库里查询得到结果
	c.JSON(200, response)                            //3.返回前端
}

func GetTeamById(c *gin.Context) {
	teamid, _ := strconv.Atoi(c.Query("teamid"))
	response := model.GetTeamById(teamid)
	fmt.Println(*response)
	c.JSON(200, *response)
}

// 2.2查询球赛
func GetAGames(c *gin.Context) {
	response := model.GetAllGames()
	c.JSON(200, response)
}

// 2.3查询用户预约
func GetAppoint(c *gin.Context) {
	UserName, OK := c.GetQuery("username")
	if OK {
		response := model.GetAppoint(UserName)
		c.JSON(200, response)
	}

}

// 3.更改**************************************************************************************
// 3.1 更新修改运动员的team

func UpdatePTeam(c *gin.Context) {
	PlayerId, _ := strconv.Atoi(c.Query("playerid"))
	NewTeamID, _ := strconv.Atoi(c.Query("newteamid"))

	model.UpdatePTeam(PlayerId, NewTeamID)

	c.JSON(200, "player的newteamid修改成功")
}

// 3.2 更新修改团队的信息
func UpdateTeamName(c *gin.Context) {
	TeamID, _ := strconv.Atoi(c.Query("teamid"))
	NewName, OK := c.GetQuery("newname") // 可以把所有的c.Query()换掉
	if OK {
		model.UpdateTeamName(TeamID, NewName)
		c.JSON(200, "修改teamname成功")
	} else {
		c.JSON(500, "获取newname错误")
	}
}
func UpdateTeamLogo(c *gin.Context) {
	TeamID, _ := strconv.Atoi(c.Query("teamid"))
	NewLogo, OK := c.GetQuery("newlogo")
	if OK {
		model.UpdateTeamLogo(TeamID, NewLogo)
		c.JSON(200, "修改teamlogo成功")
	} else {
		c.JSON(500, "获取newlogo错误")
	}
}
func UpdateTeamInfo(c *gin.Context) {
	TeamID, _ := strconv.Atoi(c.Query("teamid"))
	NewInfo, OK := c.GetQuery("newinfo")
	if OK {
		model.UpdateTeamInfo(TeamID, NewInfo)
		c.JSON(200, "修改teaminfo成功")
	} else {
		c.JSON(500, "获取newinfo错误")
	}
}

// 3.3 更改用户权限
func UpdateUserRight(c *gin.Context) {
	NewRole, _ := strconv.Atoi("newrole")
	UserName := c.Query("username")
	model.UpdateUserRight(UserName, NewRole)
}
