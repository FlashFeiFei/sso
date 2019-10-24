package user

import "github.com/jinzhu/gorm"

//用户模型
type User struct {
	gorm.Model
	//用户姓名
	Name string `gorm:"type:varchar(30);column:name"`
	//手机号
	Phone string `gorm:"type:char(12);column:phone"`
	//账号
	Account string `gorm:"type:varchar(30);column:account"`
	//密码
	Password string `gorm:"type:varchar(64);column:password"`
}

//数据库的表名
func (User) TableName() string {
	return "user"
}
