package model

import "gorm.io/gorm"

// 定义UserEntity结构体  设置字段各种属性，如主键，默认值，索引等，配置json格式名称 如果不需要json序列化，可以不设置

type UserEntity struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50);column:username;not null;unique;comment:'用户名'" json:"username"`
	Password string `gorm:"type:varchar(50);not null;comment:'密码'" json:"password`
	NickName string `gorm:"type:varchar(50);column:nickname;not null;comment:'昵称'" json:"nickname"`
}

// 表名
func (UserEntity) TableName() string {
	return "user"
}
