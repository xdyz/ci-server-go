package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// 后面是tag 标签 传进来的实际的属性 为id 和 username
type HamUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// 定义一个组合结构体，用于最终生成token，里面存储的是用户的id和username 和 token配置 过期时间等

type MyCustomClaims struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 定义一个字节数组 作为证书签名密钥 必须唯一
var jwtSecret = []byte("jwtSecret")

// 生成token 接收一个userid username 的结构体 返回token 和 err
func GenerateToken(u HamUser) (string, error) {

	// 实例化MyCustomClaims
	myCustomClaims := &MyCustomClaims{
		UserId:   u.Id,
		UserName: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000, // 过期时间

			Issuer: "gin-blog",
		},
	}

	// 生成  tonken 实例化
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myCustomClaims)

	// 生成token字符串 传入签名密钥
	tokenString, err := token.SignedString(jwtSecret)

	// 如果出现错误 则返回空字符串和错误
	if err != nil {
		return "", err
	}

	// 成功返回token
	return tokenString, nil
}

// 将token字符串解析成claims
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 实例化token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// 如果出现错误 则返回空和错误
	if err != nil {
		return nil, err
	}

	// 如果token不是空的 则返回claims 校验token 是不是合法  是不是符合
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	// 如果什么都没有返回 nil
	return nil, nil
}
