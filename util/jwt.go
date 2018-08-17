package util

import (
	"time"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"errors"
	"strings"
	"encoding/json"
)

var timeout = 300 * 24 * time.Hour            // 超时
var sign = "db3b91483842f012e24966f9640192a8" // 签名

// 用户会话
type UserSession struct {
	Id        int    `json:"user_id"`    // 用户id
	Nickname  string `json:"nickname"`   // 昵称
	AvatarUrl string `json:"avatar_url"` // 头像
}

func (us *UserSession) TokenGenerator() string {
	user, err := json.Marshal(us)
	if err != nil {
		return ""
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = string(user)
	claims["exp"] = time.Now().Add(timeout).Unix()
	claims["iat"] = time.Now().Unix()

	userToken, err := token.SignedString([]byte(sign))
	if err != nil {
		return ""
	}

	return userToken
}

func (us *UserSession) TokenAuthenticator(userToken string) error {
	if userToken == "" {
		return errors.New("header not Authorization")
	}

	parts := strings.SplitN(userToken, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return errors.New("bad header")
	}

	token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != t.Method {
			return nil, errors.New("jwt sign method not HS256")
		}

		return []byte(sign), nil
	})
	if err != nil {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)

	err = json.Unmarshal([]byte(claims["sub"].(string)), us)
	if err != nil {
		return err
	}

	return nil
}
