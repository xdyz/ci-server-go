package service

import (
	"go-basic-web/dto"
	"go-basic-web/global"
	"go-basic-web/model"
	"go-basic-web/utils"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Tags     用户
// @Summary  请求用户列表
// @Param    page  query  int     true   "页码"    default(1)
// @Param    size  query  int     true   "每页数量"  default(10)
// @Param    name  query  string  false  "用户名"
// @Router   /user [get]
func GetUsers(c *gin.Context) {

}

// 登录做校验，通过账号密码查询用户是否存在
func GetUserByName(c *gin.Context) {

	// 实例化一个结构体
	var u dto.GetUserByNameDto

	// 解析请求参数 并且校验参数 失败返回错误
	if err := c.ShouldBindJSON(&u); err != nil {
		utils.Faild(c, err.Error())

		return
	}

	// 查询用户是否存在 只取其中某些字段 这里需要用到 vo 的概念
	// global.DB.Where("username = ?", u.UserName).First(&u)

	// 这里需要将&u 改为 vo的结构体 这样只返回需要的字段就可以了
	global.DB.Model(&model.UserEntity{}).Where("username = ?", u.UserName).First(&u)

}

// GetUserById godoc
// @Tags     用户
// @Summary  通过id 查询用户
// @Param    id  path  int     true   "用户id"
// @Router   /user/{id} [get]

func GetUserById(c *gin.Context) {

}
