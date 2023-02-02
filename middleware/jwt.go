package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-gin-jwt/util"
	"go-gin-jwt/util/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(util.JwtKey) // 盐

type MyClaims struct {
	LoginName string `gorm:"json:"login_name"` // 要和User模块的字段一样
	//Password           string `gorm:"json:"password"`
	jwt.StandardClaims // 嵌套 dgrijalva/jwt-go 自己带的结构体，有签发时间、签发人...等字段
}

// SetToken 生成 token
func SetToken(loginName string) (string, int) {
	expireTime := time.Now().Add(util.ReadTimeOut) // 过期时间
	SetClaims := MyClaims{
		LoginName: loginName, // todo 这里 userId比较合理
		// todo 应该加上权限码
		//Password:  Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "super",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims) // 生成token。SigningMethodHS256 是之中的一种模式，哈希
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// CheckToken 验证 token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { // ParseWithClaims 比对token
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// JwtMiddleware jwt中间件，控制验证的东西。
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 认证的东西
		tokenHeader := c.Request.Header.Get("Authorization") // 一种固定的规范
		code := errmsg.SUCCESS

		// token 在不在
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// token格式检查。这是固定写法
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, checkCode := CheckToken(checkToken[1])
		if checkCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 查过期
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_TIMEOUT
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 查完了。用户信息写入上下文
		c.Set("loginName", key.LoginName) // todo 反user全套信息
		c.Next()
	}
}
