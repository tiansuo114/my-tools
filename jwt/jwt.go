package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

var jwtSecret []byte

var jwtInstance *jwt.Token

type Claims struct {
	jwt.StandardClaims
	data interface{} `json:"data"`
}

func init() {
	token := jwt.New(jwt.SigningMethodHS256)
	jwtInstance = token
	//jwtSecret = config.JwtSecret
}

func GenerateToken(data interface{}) (tokenStr string, err error) {
	expireTime := time.Now().Add(time.Hour * 24)

	claim := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "",
		},
		data: data,
	}

	jwtInstance.Claims = claim

	tokenStr, err = jwtInstance.SignedString(jwtSecret)

	return
}

func ParseToken(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if data, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return data, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
