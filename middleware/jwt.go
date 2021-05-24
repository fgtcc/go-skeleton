package middleware

import (
	"go-skeleton/serializer"
	"go-skeleton/utils"
	"go-skeleton/utils/errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	LoginName string `json:"login_name"`
	jwt.StandardClaims
}

func SetToken(loginName string) (string, error) {
	// expireTime := time.Now().Add(7 * 24 * time.Hour)
	expireTime := time.Now().Add(60 * time.Minute)
	SetClaims := MyClaims{
		loginName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tallsafe",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, err
}

func CheckToken(token string) (*MyClaims, error) {
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(errors.ERROR_TOKEN_WRONG, errors.GetErrMsg(errors.ERROR_TOKEN_WRONG))
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New(errors.ERROR_TOKEN_EXPIRED, errors.GetErrMsg(errors.ERROR_TOKEN_EXPIRED))
			} else {
				return nil, errors.New(errors.ERROR_TOKEN_TYPE_WRONG, errors.GetErrMsg(errors.ERROR_TOKEN_TYPE_WRONG))
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, nil
		} else {
			return nil, errors.New(errors.ERROR_TOKEN_WRONG, errors.GetErrMsg(errors.ERROR_TOKEN_WRONG))
		}
	}

	return nil, errors.New(errors.ERROR_TOKEN_WRONG, errors.GetErrMsg(errors.ERROR_TOKEN_WRONG))
}

func getToken(c *gin.Context) (string, error) {
	tokenHeader := c.Request.Header.Get("Authorization")
	if tokenHeader == "" {
		return "", errors.New(errors.ERROR_TOKEN_NOT_EXIST, errors.GetErrMsg(errors.ERROR_TOKEN_NOT_EXIST))
	}

	checkToken := strings.Split(tokenHeader, " ")
	if len(checkToken) == 0 {
		return "", errors.New(errors.ERROR_TOKEN_TYPE_WRONG, errors.GetErrMsg(errors.ERROR_TOKEN_TYPE_WRONG))
	}

	if len(checkToken) != 2 && checkToken[0] != "Bearer" {
		return "", errors.New(errors.ERROR_TOKEN_TYPE_WRONG, errors.GetErrMsg(errors.ERROR_TOKEN_TYPE_WRONG))
	}

	return checkToken[1], nil
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		claims, err := CheckToken(token)
		if err != nil {
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		c.Set("login_name", claims.LoginName)
		c.Next()
	}
}

func AdminJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		key, err := CheckToken(token)
		if err != nil {
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		if key.LoginName != "admin" {
			err := errors.New(errors.ERROR_USER_NO_RIGHT, errors.GetErrMsg(errors.ERROR_USER_NO_RIGHT))
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
