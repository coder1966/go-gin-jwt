package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-jwt/model"
	"go-gin-jwt/util/errmsg"
	"net/http"
)

// Logon 注册模块Api入口
func Logon(c *gin.Context) {
	var data model.User
	//fmt.Println("1 logon接口进来！！！")
	c.ShouldBindJSON(&data)
	fmt.Println("2 执行 c.ShouldBindJSON(&data)！结果如下：", data)
	//fmt.Println(data)
	var code int

	// todo 检查来数据是否格式合格。validate:"required,min=4,max=12" code = model.CheckLogin(data.LoginName, data.Password)

	//  todo 在数据库创建
	code = model.Logon(data.LoginName, data.Password)
	//  todo 直接变成登录状态

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
