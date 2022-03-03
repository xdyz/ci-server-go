package service

import (
	"fmt"
	"go-basic-web/dto"
	"go-basic-web/global"
	"go-basic-web/model"
	"go-basic-web/utils"
	"go-basic-web/vo"

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
	var uDto dto.GetUsersDto

	// 解析请求的query参数 并且校验参数 失败返回错误
	if err := c.ShouldBindQuery(&uDto); err != nil {
		utils.Faild(c, err.Error())
		return
	}

	uList := make([]vo.UserVo, 10)
	var uCount int64

	// 查询用户列表 分页查询
	tx := global.DB.Model(&model.UserEntity{}).Offset((uDto.Page - 1) * uDto.Size).Limit(uDto.Size).Find(&uList).Count(&uCount)

	if tx.Error != nil {
		utils.Faild(c, tx.Error.Error())
		return
	}
	utils.Success(c, "", gin.H{
		"data":  uList,
		"total": uCount,
	})
}

// 登录做校验，通过账号密码查询用户是否存在
func GetUserByName(c *gin.Context) {

	// 实例化一个结构体
	var uDto dto.GetUserByNameDto

	var uVo vo.UserVo

	// 解析请求参数 并且校验参数 失败返回错误
	if err := c.ShouldBindJSON(&uDto); err != nil {
		utils.Faild(c, err.Error())

		return
	}

	// 查询用户是否存在 只取其中某些字段 这里需要用到 vo 的概念

	// 为 vo的结构体 这样只返回需要的字段就可以了
	tx := global.DB.Model(&model.UserEntity{}).Where("username = ?", uDto.UserName).First(&uVo)

	if tx.Error != nil {
		utils.Faild(c, "查询数据失败")
		return
	}
	utils.Success(c, "", uVo)

}

// GetUserById godoc
// @Tags     用户
// @Summary  通过id 查询用户
// @Param    id  path  int     true   "用户id"
// @Router   /user/{id} [get]
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.Faild(c, "id不能为空")
		return
	}

	// 实例化一个结构体
	var user vo.UserVo

	tx := global.DB.Model(&model.UserEntity{}).Where("id = ?", id).First(&user) // 查询用户
	if tx.Error != nil {
		utils.Faild(c, tx.Error.Error())
		return
	}

	utils.Success(c, "", user)
}

// CreateUser godoc
// Tags 	 用户
// Summary 	 创建用户
// @Param 	 body 	 body 	 dto.CreateUserDto 	 true 	 "用户信息"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var uDto dto.CreateUserDto

	// 解析请求参数 并且校验参数 失败返回错误
	if err := c.ShouldBindJSON(&uDto); err != nil {
		utils.Faild(c, err.Error())
		return
	}

	fmt.Printf("uDto: %v\n", uDto)

	user := model.UserEntity{

	}

	tx := global.DB.Create(&user).
	if tx.Error != nil {
		utils.Faild(c, tx.Error.Error())
		return
	}

	utils.Success(c, "", nil)
}
