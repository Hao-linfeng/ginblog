package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:int DEFAULT:2 json:"role" validate:"required,gte=2"`
}

// 新增用户
func CreatUser(data *User) int {
	// data.Password = ScryptPw(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	} else {
		return errmsg.SUCCSE
	}

}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编译用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 加密
func (u *User) BeforeSave(tx *gorm.DB) (err error) { //狗子函数 gorm
	u.Password = ScryptPw(u.Password)
	return
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 43, 56, 234, 23, 12, 3, 45}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证
func CheckLogin(username string, passssword string) int {
	var user User

	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(passssword) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 2 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE

}
