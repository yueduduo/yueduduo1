package handler

// 这个包是用gin写接送与处理客户端信息、调取发送服务端信息
import (
	"fmt"
	"mygo/model"
	"strconv"

	"github.com/gin-gonic/gin"

)

// 注释不显示原因：注释其后的代码有空行😂


// @Summary 登记运动员
// @Description 接收players信息然后登记创建并返回到客户端反馈
// @Tags player
// @Accept json
// @Produce json
// @Failure 500 {string} json {"error": err.Error()}
// @Success 200 {string} string "new player创建成功"
// @Router /players [post]
func ChinPlayer(c *gin.Context) {
	NewPlayer := &model.Player{}                    // 空的东西
	if err := c.ShouldBind(NewPlayer); err != nil { //这里的c.ShouldBind是把客户端提交的(json)存到后端服务器内存的NewPlayer变量里 // 前端的信息应该先是在c这个*gin.Context里
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	NewPlayer.Checkin()           //这个是存到将内存变量(结构体)的值(内存暂时存的)存到数据库中
	c.JSON(200, "new player创建成功") //服务器给客户端返回信息
}


// @Summary 登记队伍
// @Description 接收teams信息然后登记创建并返回到客户端反馈
// @Tags team
// @Accept json
// @Produce json
// @Failure 500 {string} json {"error": err.Error()}
// @Success 200 {string} string "new team创建成功"
// @Router /teames [post]
func ChinTeam(c *gin.Context) {
	NewTeam := &model.Team{}
	if err := c.ShouldBind(NewTeam); /*这句放if语句前面也行，后面才是*/ err == nil {
		NewTeam.Checkin() // 先识别客户端的信息，再存进数据库
		c.JSON(200, "new team创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// @Summary 创建球赛
// @Description 接收game信息然后创建并返回到客户端反馈
// @Tags game
// @Accept json
// @Produce json
// @Failure 500 {string} json {"error": err.Error()}
// @Success 200 {string} string "new game创建成功"
// @Router /games [post]
func CreateGame(c *gin.Context) {
	NewGame := &model.Game{}
	if err := c.ShouldBind(NewGame); err == nil {
		NewGame.CreateGame()
		c.JSON(200, "new game创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// @Summary  创建用户
// @Description 接收user信息然后创建并返回到客户端反馈
// @Tags user
// @Accept json
// @Produce json
// @Failure 500 {string} json {"error": err.Error()}
// @Success 200 {string} string "new user创建成功"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	NewUser := &model.User{}
	if err := c.ShouldBind(NewUser); err == nil {
		NewUser.Register()
		c.JSON(200, "new user 创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

// @Summary  创建用户预约
// @Description 接收user预约信息然后创建并返回到客户端反馈
// @Tags appoint
// @Accept json
// @Produce json
// @Failure 500 {string} json {"error": err.Error()}
// @Success 200 {string} string "new appoint创建成功"
// @Router /users/appoint/{username} [post]
func CreateAppoint(c *gin.Context) {
	UserAppointName := c.Query("username") + "appoints"
	NewAppoint := &model.UserAppoint{}

	if err := c.ShouldBind(NewAppoint); err == nil {
		fmt.Println(NewAppoint)
		NewAppoint.CreateAppoint(UserAppointName) //还传入指定的表名，预约比赛的不同用户放在对应不同的appoint表，
		// 还应有个对games表对应game行的appoint列update加1, fangzai了model.CreateAppoint()中
		c.JSON(200, "new appoint创建成功")
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}


// @Summary  查询player
// @Description 查询指定id的player信息并返回到客户端
// @Tags player
// @Accept json
// @Produce json
// @Success 200 {object} model.Player
// @Router /players/{playerid} [get]
func GetPlayerById(c *gin.Context) {
	playerid, _ := strconv.Atoi(c.Query("playerid")) //1.接收前端get的query内容
	response := model.GetPlayerById(playerid)        //2.在数据库里查询得到结果
	c.JSON(200, response)                            //3.返回前端
}
// @Summary  查询team
// @Description 查询指定id的team信息并返回到客户端
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {object} model.Team
// @Router /teams/{teamid} [get]
func GetTeamById(c *gin.Context) {
	teamid, _ := strconv.Atoi(c.Query("teamid"))
	response := model.GetTeamById(teamid)
	fmt.Println(*response)
	c.JSON(200, *response)
}


// @Summary  查询game
// @Description 查询all games信息并返回到客户端
// @Tags game
// @Accept json
// @Produce json
// @Success 200 {object} []model.Game
// @Router /games [get]
func GetAGames(c *gin.Context) {
	response := model.GetAllGames()
	c.JSON(200, response)
}


// @Summary  查询用户预约
// @Description 查询用户预约的比赛详情
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} []model.UserAppoint
// @Router /users/appoint/{username} [get]
func GetAppoint(c *gin.Context) {
	UserName, OK := c.GetQuery("username")
	if OK {
		response := model.GetAppoint(UserName)
		c.JSON(200, response)
	}
}

// @Summary  更新运动员的team信息
// @Description 更新修改运动员的team
// @Tags player
// @Accept json
// @Produce json
// @Success 200 {string} string "player的newteamid修改成功"
// @Router /teams/{playerid} [put]
func UpdatePTeam(c *gin.Context) {
	PlayerId, _ := strconv.Atoi(c.Query("playerid"))
	NewTeamID, _ := strconv.Atoi(c.Query("newteamid"))

	model.UpdatePTeam(PlayerId, NewTeamID)

	c.JSON(200, "player的newteamid修改成功")
}


// @Summary  更新team的name
// @Description 更新team的name
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {string} string "player的newteamid修改成功"
// @Router /teams/{teamid}/{newname} [put]
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
// @Summary  更新team的logo
// @Description 更新team的logo
// @Tags team
// @Accept json
// @Produce json
// @Failure 500 {string} string "获取newlogo错误" 
// @Success 200 {string} string "修改teamlogo成功"
// @Router /teams/{teamid}/{newlogo} [put]
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
// @Summary  更新team的info
// @Description 更新team的info
// @Tags team
// @Accept json
// @Produce json
// @Failure 500 {string} string "获取newinfo错误" 
// @Success 200 {string} string "修改teaminfo成功"
// @Router /teams/{teamid}/{newinfo} [put]
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

// @Summary  更改用户权限
// @Description 更改用户权限role
// @Tags team
// @Accept json
// @Router /teams/{usename/{newrole} [put]
func UpdateUserRight(c *gin.Context) {
	NewRole, _ := strconv.Atoi("newrole")
	UserName := c.Query("username")
	model.UpdateUserRight(UserName, NewRole)
}
