package model

import (
	"github.com/jinzhu/gorm"
	"go-gin-jwt/util/errmsg"
)

// Category 分类
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var categories Category
	db.Where("name = ?", name).First(&categories)
	if categories.ID > 0 {
		return errmsg.ERROR_CATEGRORYNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCategory 添加分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategory 查询单个分类
//func GetCategory(c *gin.Context) {
//
//}

// GetCategories 查询分类列表。牵扯到分页。
func GetCategories(pageSize int, currPageNum int) ([]Category, int) {
	var categories []Category
	var total int // 查询到的记录的总条数。为了方便分页
	err = db.Limit(pageSize).Offset((currPageNum - 1) * pageSize).Find(&categories).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return categories, total
}

// EditCategory 编辑分类。不含修改密码。修改/重置密码是一个单独的接口。map传参比struct传参，0值也会更新，所以用map
func EditCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&Category{}).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类下的所有文章

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
