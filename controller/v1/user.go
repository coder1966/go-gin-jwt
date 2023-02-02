package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/model"
	"go-gin-jwt/response"
	"go-gin-jwt/util/errmsg"
	"go-gin-jwt/util/validator"
	"net/http"
	"strconv"
)

var code int

// UserExist 查询用户是否存在
func UserExist(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	// todo 添加用户
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	// 进一步数据验证。例如验证登录名、密码长度。在
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		//c.JSON(http.StatusOK, gin.H{"status": code, "message": msg})
		response.Response(c, http.StatusUnprocessableEntity, code, nil, msg) // todo 这个通用返回方法
		return
	}
	// 到这里，验证结束

	code = model.CheckUser(data.LoginName)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUser 查询单个用户
//func GetUser(c *gin.Context) {
//
//}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) // strconv.Atoi string转int
	currPageNum, _ := strconv.Atoi(c.Query("currpagenum"))
	if pageSize == 0 {
		pageSize = -1 // GORM 规定 -1 表示分页要求失效
	}
	if currPageNum == 0 {
		currPageNum = -1 // GORM 规定- 1 表示分页要求失效
	}
	data, total := model.GetUsers(pageSize, currPageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Query("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.LoginName) // 其实，这个检查也可以用勾子函数来写
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
