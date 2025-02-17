package shared_utils

import (
	shared_configs "delivery/internal/shared/Application/Configs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func JwtToken(email string, user_id int) (string, error) {

	conf, err := shared_configs.GetConfig()
	if err != nil {
		return "", errors.New("config error")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":   email,
			"user_id": user_id,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(conf.Jwt.Secret))
	if err != nil {

		return "", err
	}

	return tokenString, nil
}
