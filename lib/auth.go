package lib

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SECRETKEY = "fuwujisuanyyds"  // 密钥
)

type UserInfo struct {
	Username string
	ID uint64
}

func CreateToken(user *UserInfo) (tokeness string, err error) {
	claims := jwt.MapClaims{
		"id": user.ID,
		"username": user.Username,
		"nbf": time.Now().Unix(),
		"iat":time.Now().Unix(),
		"exp":time.Now().Add(time.Hour*24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokeness, err = token.SignedString([]byte(SECRETKEY))
	return
}

func ParseToken(tokeness string) (err error){
	//user = &UserInfo{}
	token, err := jwt.Parse(tokeness, func(token *jwt.Token) (i interface{}, err error) {
		value, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected method, token.Header['alg'] = ", value)
		}
		return []byte(SECRETKEY), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//user.ID = uint64(claims["id"].(float64))
		//user.Username = claims["usrname"].(string)
		return nil
	}else{
		return  err
	}
}
