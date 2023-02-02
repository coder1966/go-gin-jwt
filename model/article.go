package model

import (
	"github.com/jinzhu/gorm"
	"go-gin-jwt/util/errmsg"
)

// Article 文章本身
type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Category Category `gorm:"foreignKey:Cid"`                // 对应文章结构提.foreignKey外键
	CID      int      `gorm:"type:int" json:"cid"`           // Article 的id
	Desc     string   `gorm:"type:varchar(200)" json:"desc"` // description 描述
	Content  string   `gorm:"type:longtext" json:"content"`  // (文章的)内容
	Img      string   `gorm:"type:varchar(200)" json:"img"`  // (文章的)图片
}

// CreateArticle 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategoryArticles 查询分类下所有文章
func GetCategoryArticles(id int, pageSize int, currPageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int // 查询到的记录的总条数。为了方便分页
	err := db.Preload("Category").Limit(pageSize).Offset((currPageNum-1)*pageSize).Where("cid = ?", id).Find(&articleList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGRORY_NOT_EXIST, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// GetArticle GetArticle 查询单个文章
func GetArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// GetArticles 查询文章列表。牵扯到分页。
func GetArticles(pageSize int, currPageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int                                                                                                            // 查询到的记录的总条数。为了方便分页
	err = db.Preload("Category").Limit(pageSize).Offset((currPageNum - 1) * pageSize).Find(&articleList).Count(&total).Error // gorm使用preload解决一对多关系
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// EditArticle 编辑文章。不含修改密码。修改/重置密码是一个单独的接口。map传参比struct传参，0值也会更新，所以用map
func EditArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.CID
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&Article{}).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询文章下的所有文章

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
