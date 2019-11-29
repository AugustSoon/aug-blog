package token

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")

type Context struct {
	ID      uint
	Account string
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		id, _ := strconv.Atoi(claims.Id)
		ctx.ID = uint(id)
		ctx.Account = claims.Audience
		return ctx, nil
	}

	return ctx, err
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string

	_, err := fmt.Sscanf(header, "Bearer %s", &t)

	if err != nil {
		return &Context{}, ErrMissingHeader
	}

	return Parse(t, viper.GetString("jwt_secret"))
}

func Sign(c Context) (tokenString string, err error) {
	secret := viper.GetString("jwt_secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  c.Account,
		ExpiresAt: time.Now().Unix() + viper.GetInt64("jwt_ttl"),
		Id:        string(c.ID),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Aug",
		NotBefore: time.Now().Unix(),
		Subject:   "login",
	})

	tokenString, err = token.SignedString([]byte(secret))

	return
}
