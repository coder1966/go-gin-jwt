package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-jwt/middleware"
	"go-gin-jwt/model"
	"go-gin-jwt/util/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	fmt.Println("2 login 执行 c.ShouldBindJSON(&data)！结果如下：", data)
	var token string
	var code int
	code = model.CheckLogin(data.LoginName, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.LoginName)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
