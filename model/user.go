package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"go-gin-jwt/util/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
)

// User 用户
type User struct {
	gorm.Model
	LoginName string `gorm:"type:varchar(20);not null" json:"loginName" validate:"required,min=4,max=12" label:"登录名"`
	Password  string `gorm:"type:varchar(64);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email     string `gorm:"type:varchar(20);not null" json:"email"`
	NiceName  string `gorm:"type:varchar(20);not null" json:"nice_name"`
	Sex       int32  `gorm:"type:int" json:"sex"`
	Role      int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2,max=20" label:"角色码"` // 角色，默认2，validate数据校验tag
	Avatar    string `gorm:"type:varchar(200);not null" json:"avatar"`                                    // 头像
}

// 在模型里，写数据库操作的方法

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 添加用户
func CreateUser(data *User) int {
	data.Password = ScryptPassword(data.Password) // 密码转为加密的。（另一个方法是用GORM的勾子BeforeSave()前加密）
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUser 查询单个用户
//func GetUser(c *gin.Context) {
//
//}

// GetUsers 查询用户列表。牵扯到分页。
func GetUsers(pageSize int, currPageNum int) ([]User, int) {
	var users []User
	var total int // 查询到的记录的总条数。为了方便分页
	err = db.Limit(pageSize).Offset((currPageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// EditUser 编辑用户。不含修改密码。修改/重置密码是一个单独的接口。map传参比struct传参，0值也会更新，所以用map
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["loginName"] = data.LoginName
	maps["email"] = data.Email
	maps["niceName"] = data.NiceName
	maps["sex"] = data.Sex
	maps["role"] = data.Role
	maps["avatar"] = data.Avatar
	err = db.Model(&User{}).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPassword 密码加密。专业方法。https://pkg.go.dev/golang.org/x/crypto/scrypt
func ScryptPassword(password string) string {
	// func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error) // salt =程序员定义盐；N=cpu开销，2~2^29；keyLen=最后哈希长度
	const keyLen = 32                                                            // 最后的哈希的长度
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}               // 盐
	HashPassword, err := scrypt.Key([]byte(password), salt, 1<<11, 8, 1, keyLen) // 得到专业级加密的密码
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(HashPassword) // 返回base64字符串
}

//// BeforeSave 钩子函数方法，写入前自动执行。另外的方法是data.password从api一进来就加密。
//func (u *User) BeforeSave() {
//	u.Password = ScryptPassword(u.Password)
//}

// CheckLogin 登录验证
func CheckLogin(loginName string, password string) int {
	var user User

	db.Where("login_name =?", loginName).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ERROR_PASSWORD
	}
	//if user.Role != 1 {
	//	return errmsg.ERROR_USER_NO_RIGHT
	//}
	return errmsg.SUCCESS
}

// Logon 检查登录名在数据库是否重名
func Logon(loginName string, password string) int {
	var user User

	db.Where("login_name =?", loginName).First(&user)
	// 重名
	if user.ID != 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	// 生成密码

	// 创建用户
	user.LoginName = loginName
	user.Password = password
	var code int
	code = CreateUser(&user)

	return code
}
