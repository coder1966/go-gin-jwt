package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/model"
	"go-gin-jwt/util/errmsg"
	"net/http"
	"strconv"
)

// 添加文章
// 查询单个文章
// 查询文章列表
// 编辑文章
// 删除文章

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	// todo 添加文章
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCategoryArticles 查询分类下所有文章
func GetCategoryArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) // strconv.Atoi string转int
	currPageNum, _ := strconv.Atoi(c.Query("currpagenum"))
	id, _ := strconv.Atoi(c.Query("id"))
	if pageSize == 0 {
		pageSize = -1 // GORM 规定 -1 表示分页要求失效
	}
	if currPageNum == 0 {
		currPageNum = -1 // GORM 规定- 1 表示分页要求失效
	}

	data, code, total := model.GetCategoryArticles(id, pageSize, currPageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo GetArticle 查询单个文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticles 查询文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) // strconv.Atoi string转int
	currPageNum, _ := strconv.Atoi(c.Query("currpagenum"))
	if pageSize == 0 {
		pageSize = -1 // GORM 规定 -1 表示分页要求失效
	}
	if currPageNum == 0 {
		currPageNum = -1 // GORM 规定- 1 表示分页要求失效
	}
	data, code, total := model.GetArticles(pageSize, currPageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Query("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
