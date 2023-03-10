package response

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/util/errmsg"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, errmsg.SUCCESS, data, msg)
}

func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, errmsg.ERROR, data, msg)
}
