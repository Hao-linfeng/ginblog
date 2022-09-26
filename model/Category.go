package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
}

// 新增用户
func CreatCategory(data *Category) int {
	// data.Password = ScryptPw(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	} else {
		return errmsg.SUCCSE
	}

}

// 查询用户列表
func GetCategorys(pageSize int, pageNum int) ([]Category, int) {
	var cates []Category
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, int(total)
}

// 编译用户
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询分类下文章
