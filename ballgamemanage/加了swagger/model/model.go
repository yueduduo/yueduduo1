package model

// 这个包是用gorm写服务端的数据库各种模型与对应的crud操作
import (
	"fmt"
	"time"

	// "time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// 1.定义数据库model**********************************************************
type Player struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Info   string `json:"info"`
	TeamID int    `json:"team_id"`
}

type Team struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	TeamInfo string `json:"info"`
}

type Game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Data        string `json:"data"`
	Place       string `json:"place"`
	GameInfo    string `json:"info"`
	Appointment int    `gorm:"default:0"`
}

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Role            uint32 `json:"role"`
	AppointGameName string `json:"appoint_game_name"`
}

type UserAppoint struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	AppointGameName string `json:"appoint_game_name"`
}

// 2.连接数据库，得到db连接，并进行初始化*******************************************************************
func Init() *gorm.DB {
	// 连库
	dsn := "root:@tcp(localhost:3306)/bmanagement?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil { //检查连接错误
		panic(err)
	}
	db = d
	// 建表
	db.AutoMigrate(&Team{})
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Game{})
	db.AutoMigrate(&User{})
	return db
}

// 3.定义对象的操作*******************************************************************
// 3.0 登录询前端输入的数据，判断用户名是否存在来登录，并判断类型
func Login(UserName string) (string, int) {
	var user User
	err := db.Where("Name = ?", UserName).First(&user).Error
	if err == nil {
		return "登录成功", int(user.Role)
	} else {
		return fmt.Sprint(err) + "用户不存在，请注册用户！", -1
	}
}

// 3.1.1创建注册登记创建
func (u *User) Register() {
	db.Create(u)
	db.AutoMigrate(&UserAppoint{})
}
func (p *Player) Checkin() {
	db.Create(p)
}
func (t *Team) Checkin() {
	db.Create(t)
}

// 3.1.2创建比赛
func (g *Game) CreateGame() {
	db.Create(g)
}

// 3.1.3创建预约
func (ua *UserAppoint) CreateAppoint(uan string) {
	db.Table(uan).Create(ua)
	db.Exec("UPDATE games SET appoint = appoint + 1 where name = ?", ua.AppointGameName)
}

// ——————————————————————————————————————————————————————————————————————————————————————————————————
// 3.2.1以id查询
func GetPlayerById(Id int) Player {
	var getPlayer Player
	db.Where("ID = ?", Id).First(&getPlayer)
	fmt.Println(getPlayer)
	return getPlayer
}

func GetTeamById(Id int) *Team {
	GetTeam := &Team{}
	err := db.Where("ID = ?", Id).First(GetTeam)
	fmt.Println(err)
	return GetTeam
}

func GetAllGames() []Game {
	var Games []Game
	db.Exec("SELECT * FROM games ORDERED BY date DESC LIMIT 0,10").Scan(&Games)
	return Games
}

// // 3.2.3查询用户预约
func GetAppoint(UserName string) []UserAppoint {
	var UAppoints []UserAppoint
	UserName = UserName + "Appoints"
	db.Table(UserName).Find(&UAppoints)
	return UAppoints
}

// 3.3更新———————————————————————————————————————————————————————————————————

// 3.3.1更新修改运动员的team
func UpdatePTeam(ID int, NewTeamID int) {
	db.Model(&Player{}).Where("id = ?", ID).Update("TeamID", NewTeamID)
}

// 3.3.2更新修改团队的信息
func UpdateTeamName(TeamID int, NewName string) {
	db.Model(&Team{}).Where("id = ?", TeamID).Update("Name", NewName)
}
func UpdateTeamLogo(TeamID int, NewLogo string) {
	db.Model(&Team{}).Where("id = ?", TeamID).Update("Logo", NewLogo)
}
func UpdateTeamInfo(TeamID int, NewInfo string) {
	db.Model(&Team{}).Where("id = ?", TeamID).Update("Info", NewInfo)
}

// 3.3.3更改用户权限
func UpdateUserRight(UserName string, NewRole int) {
	db.Model(&User{}).Where("name = ?", UserName).Update("role", NewRole)
}
