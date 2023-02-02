package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-gin-jwt/controller/v1"
	"go-gin-jwt/middleware"

	"go-gin-jwt/util"
)

// func InitRouter() *gin.Engine { // 也可以func InitRouter() *gin.Engine 返回一个*gin.Engine,在main里面调用

func InitRouter() {
	// 入参
	gin.SetMode(util.AppMode) // 运行模式。生产 release /开发 debug 环境。
	//r := gin.Default() // 路由初始化。都可以，Default() 默认加了两个中间件，日志和？？
	r := gin.New()             // 路由初始化。都可以，没有Default() 默认加的日志等两个中间件
	r.Use(middleware.Logger()) // 引进自定义日志中间件
	r.Use(gin.Recovery())      // gin.Default()本来默认有的中间件，因为自定义了gin.New(),所以补上。Recovery 中间件会恢复任何恐慌(panics) 如果存在恐慌，中间件将会写入500。

	// 需要鉴权的分一个路由组，加中间件
	authorRouterV1 := r.Group("api/v1")
	authorRouterV1.Use(middleware.JwtMiddleware())
	{
		// 用户模块的路由接口
		authorRouterV1.PUT("user/:id", v1.EditUser)
		authorRouterV1.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		authorRouterV1.POST("category/add", v1.AddCategory)
		authorRouterV1.PUT("category/:id", v1.EditCategory)
		authorRouterV1.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		authorRouterV1.POST("article/add", v1.AddArticle)
		authorRouterV1.PUT("article/:id", v1.EditArticle)
		authorRouterV1.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件的路由接口
		authorRouterV1.POST("upload", v1.Upload)

		//router.GET("hello", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": "OK，成功！",
		//	})
		//})
	}

	// 不需要鉴权的分一个路由组
	publicRouterV1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		publicRouterV1.GET("users", v1.GetUsers)

		// 分类模块的路由接口
		publicRouterV1.GET("categories", v1.GetCategories)

		// 文章模块的路由接口
		publicRouterV1.GET("articles", v1.GetArticles)
		publicRouterV1.GET("articles/bycategory/:id", v1.GetCategoryArticles)
		publicRouterV1.GET("article/:id", v1.GetArticle)

		// 登录
		publicRouterV1.POST("login", v1.Login)

		// 注册
		publicRouterV1.POST("logon", v1.Logon)
		// publicRouterV1.POST("user/add", v1.AddUser) // 用注册代替了

		//router.GET("hello", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": "OK，成功！",
		//	})
		//})
	}

	r.Run(util.HttpPort) // 跑起来。也可以func InitRouter() *gin.Engine 返回一个*gin.Engine,在main里面调用
	//return
}
