package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	ID uint `json:"userId"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2
var CustomSecret = []byte("dapan")

func GenToken(userId uint) (string, error) {
	fmt.Println(CustomSecret)
	c := CustomClaims {
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), 
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(CustomSecret)
}

// 解析token
func parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(tokenString *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})

	if err != nil {
		return nil, err
	}
	
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	
	return nil, errors.New("invalid token")
}

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	claims, err := parseToken(token)

	if err != nil {
		c.JSON(422, gin.H{
			"msg":  err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("user_id", claims.ID)
	c.Next() 
}