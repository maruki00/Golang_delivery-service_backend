package shared_middlewares

import (
	shared_configs "delivery/internal/shared/Application/Configs"
	shareddb "delivery/internal/shared/Infrastructure/DB"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func verifyToken(tokenString string) error {
	conf, err := shared_configs.GetConfig()
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Jwt.Secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func AuthRequired() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" || len(tokenString) < 8 {
			ctx.AbortWithStatusJSON(401, map[string]string{
				"error: ": "Missing authorization header",
			})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 7)
		err := verifyToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(401, map[string]string{
				"error: ": "unauthorized",
			})
			return
		}

		db := shareddb.NewDB()

		st, err := db.Prepare("select user_id from auths where token = ?")
		if err != nil {
			ctx.AbortWithStatusJSON(401, map[string]string{
				"error: ": "auth table error",
			})
			return
		}

		var user_id int
		err = st.QueryRow(tokenString).Scan(&user_id)
		if err != nil {
			ctx.AbortWithStatusJSON(401, map[string]string{
				"error: ": "token not found " + err.Error(),
			})
			return
		}

		ctx.Next()
	}
}
