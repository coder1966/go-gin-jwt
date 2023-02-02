package middleware

// 跨域中间件
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	// return cors.Default() // 这种就全部不跨域
	return cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"}, // 允许的域名
		AllowOrigins: []string{"*"}, // 允许的域名。全部许可
		//AllowMethods:     []string{"PUT", "PATCH"}, // 允许的请求方法
		AllowMethods:  []string{"*"}, // 允许的请求方法
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length", "Content-Type", "Authorization"},
		//AllowCredentials: true, // 是否发送Cooke请求
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour, // 预请求时间，12小时内不再需要预请求
	})
}
