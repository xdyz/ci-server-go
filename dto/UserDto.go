package dto

// 定义UserDto结构体 json转换 bingding 格式校验
type GetUserByNameDto struct {
	UserName string `json:"username" bingding:"required"`
	Password string `json:"password" bingding:"required"`
}
