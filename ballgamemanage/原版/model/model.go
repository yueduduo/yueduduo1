package model

// 这个包是用gorm写服务端的数据库各种模型与对应的crud操作
import (
	"fmt"
	// "time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// 1.定义数据库model******************************************************************
type Player struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Info   string `json:"info"`
	TeamID int    `json:"team_id"`
}

type Team struct {
	Id   	int    `json:"id"`
	Name 	string `json:"name"`
	Logo 	string `json:"logo"`
	TeamInfo 	string `json:"info"`
} 

type Game struct {
	Id   	int    `json:"id"`
	Name 		string `json:"name"`
	// CreateAt time.Time	//`json:"createat"`//没用好
	Data        string `json:"data"`
	Place 		string `json:"place"`
	// Teams   []string `json:"teams"`// 报错未解决，应该就直接存不了切片
	// Team, 晕了不知道Game的表怎么写了
	GameInfo    	string `json:"info"`
	Appointment int    `gorm:"default:0"`
}

type User struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Avatar              string `json:"avatar"`
	Role                uint32 `json:"role"`
	AppointGameName     string `json:"appoint_game_name"`
	/* Game   Game  `gorm:"foreignKey:GameAppoint; association_foreignkey:Appoint"`
	出错原因是先是没写foreignKey:GameAppoint
	既要写定义外键，又要写定义所属表的关联（外）键, 或者是我写错字母了*/
}

type UserAppoint struct {
	gorm.Model
	/* UserUser   `gorm:""`报错invalid field found for struct mygo/model.UserAppoint's field User: 
	define a valid foreign key for relations or implement the Valuer/Scanner interface*/
	AppointGameName string `json:"appoint_game_name"`
}

// 2.连接数据库，得到db连接，并进行初始化***********************************************************************************
func Init() *gorm.DB {
	// 一直报错一方面是因为在main包里注释了Init函数就没连上数据库
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
func Login(UserName string)(string, int){
	var user User
	err := db.Where("Name = ?", UserName).First(&user).Error
	if err == nil{
		return "登录成功", int(user.Role)
	}else{
		return fmt.Sprint(err) + "用户不存在，请注册用户！", -1
		// 不知道如何去掉F:/Golang/Agoproject/src/22spring_study/ballgamemange/model/model.go:79 record not found服务器的报错，应该就不能去掉
	}
}

// 3.1.1创建注册登记创建
func (u *User) Register() {
	db.Create(u)
	db.AutoMigrate(&UserAppoint{})
	// 将appoint表重命名为u.Name+"Appoint"
	db.Migrator().RenameTable(&UserAppoint{}, u.Name+"Appoints") // 测试知，不会自己加上s
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
	/*db.Exec("INSERT INTO `?` (appoint_game_name) VALUES(?)", uan, ua.AppointGameName) 
	报错Table 'bmanagement.?' doesn't exist
	这里的指针没问题,但是引号也传进进去了，导致表名有引号，错误*/
	db.Exec("UPDATE games SET appoint = appoint + 1 where name = ?", ua.AppointGameName) //
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
// 不明白为什么上面不行
// func GetTeamById(Id int) Team {
// 	var GetTeam Team
// 	err := db.Where("ID = ?", Id).First(&GetTeam).Error
// 	fmt.Println(err)
// 	return GetTeam
// }

// 3.2.2查询比赛
// type ListResponse struct {
// 	Name        string
// 	Date        string
// 	Place       string
// 	Info        string
// 	Appointment int 
// 	TeamA       *Team
// 	TeamB       *Team
// }
// type TeamNames struct{
// 	team1 		string
// 	team2 		string
// }
// func GetAllGames() []ListResponse {
// 	var Games []ListResponse
// 	var Team1And2 TeamNames
// 	db.Table("game").Select("team1name", "team2name").First(&Team1And2)

// 	db.Table("game").Select("name", "place", "info", "appointment").Find(&Games)
// 	db.Table("team").Where("name = ?", Team1And2.team1).First(&Games)
// 	return Games
// }
//  晕了不会了
func GetAllGames() []Game{
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

// 3.3更新——————————————————————————————————————————————————————————————————————————————————————————————————

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
