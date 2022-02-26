package common

import (
	"fmt"
	"go-basic-web/global"
	"go-basic-web/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDataBase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		"root",            // 用户名
		"123456",          // 密码
		"127.0.0.1",       // 主机
		"3306",            // 端口
		"go_basic",        // 数据库
		"utf8",            // 字符集
		"Asia%2FShanghai", // 时区
	)

	// 连接数据库 如果没有错误，则返回数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 如果有错误，打印错误信息 并终止程序
	if err != nil {
		panic(err)
	}

	// 返回数据库连接
	return db
}

// 初始化数据库 项目启动时调用
func init() {
	fmt.Println("初始化数据库")
	global.DB = InitDataBase()

	// 初始化数据库表
	global.DB.AutoMigrate(&model.UserEntity{})
}
