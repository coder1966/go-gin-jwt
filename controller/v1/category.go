package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/model"
	"go-gin-jwt/util/errmsg"
	"net/http"
	"strconv"
)

// 查询分类是否准在

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	// todo 添加分类
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGRORYNAME_USED {
		code = errmsg.ERROR_CATEGRORYNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCategory 查询单个分类
//func GetCategory(c *gin.Context) {
//
//}

// todo 查询分类下的所有文章

// GetCategories 查询分类列表
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) // strconv.Atoi string转int
	currPageNum, _ := strconv.Atoi(c.Query("currpagenum"))
	if pageSize == 0 {
		pageSize = -1 // GORM 规定 -1 表示分页要求失效
	}
	if currPageNum == 0 {
		currPageNum = -1 // GORM 规定- 1 表示分页要求失效
	}
	data, total := model.GetCategories(pageSize, currPageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Query("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name) // 其实，这个检查也可以用勾子函数来写
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATEGRORYNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
