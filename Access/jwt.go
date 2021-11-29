package Access

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	JwtSignKey = "my_sign_key"
	JwtExpires = 60 * 60
)

/*
jwt 通过 json 传输
传输数据通过数据签名相对比较安全。客户端与服务端通过 jwt 交互，服务端通过解密 token 信息，来实现用户认证。
不需要服务端集中维护 token 信息，便于扩展。当然 jwt 也有其缺点
1.客户单将username和Password 进行base64编码，然后进行传递
2.后端解析，加入一些过期信息等，传递给前端
3.前端每次将token传递给后端，后端解析出username进行用户信息管理
*/
func Tokens(c *gin.Context) (string ,error) {
	//解析 Authorization
	splits := strings.Split(c.GetHeader("Authorization"), " ")
	if len(splits) != 2 {
		return "", errors.New("用户名或密码格式错误")
	}

	//base64 解码
	appSecret, err := base64.StdEncoding.DecodeString(splits[1])
	if err != nil {
		return "", err
	}

	fmt.Println(string(appSecret))
	//获得username和password
	parts := strings.Split(string(appSecret), ":")
	if len(parts) != 2 {
		return "", errors.New("用户名或密码格式错误")
	}
	//TODO
	//通常是username 和password校验
	//校验通过生成token
	claims := jwt.StandardClaims{
		Issuer:    parts[0],
		ExpiresAt: time.Now().Add(JwtExpires * time.Second).Unix(),
	}

	token,err := JwtEncode(claims)
	return token, err
}

//JwtAuth jwt验证接口，返回username 和 err
func JwtAuth(c *gin.Context) (string, error) {
	claims, err := JwtDecode(strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", ""))
	if err != nil {
		return "", err
	}
	fmt.Println(claims)

	//TODO 一般会根据claims.Issuer获取用户信息
	userName := claims.Issuer
	return userName, nil
}


func JwtDecode(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSignKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("token is not jwt.StandardClaims")
	}
}

func JwtEncode(claims jwt.StandardClaims) (string, error) {
	mySigningKey := []byte(JwtSignKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}