package main

import (
	"crypto/md5"
	auth_domain_dto "delivery/Services/Auth/Domain/DTOs"
	shared_infrastructure_db "delivery/Services/Shared/Infrastructure/DB"
	"encoding/base64"
	"errors"
	"fmt"
)

//import (
// 	authRoutes "delivery/Services/Auth/Application/Routes"
// 	orderRoutes "delivery/Services/Order/Application/Routes"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

func stdOTP(login, password string) (*auth_domain_dto.AuthDTO, error) {
	db := shared_infrastructure_db.NewDB()
	defer db.Close()
	dto := &auth_domain_dto.AuthDTO{}
	dto.Token = base64.NewEncoding("base64").EncodeToString([]byte("helloworld"))
	hash := md5.Sum([]byte(password))
	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT id, email, user_type, user_level FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	err := statement.QueryRow(login, h).Scan(&dto.User_id, &dto.Email, &dto.User_type, &dto.User_level)
	if err != nil {
		return nil, errors.New("invalid creadentails ")
	}
	if dto.Email == "" || dto.User_level == "" || dto.User_type == "" {
		return nil, errors.New("invalid credentials")
	}

	return dto, nil
}

func main() {

	// router := gin.Default()
	// authRoutes.AuthRouter(router)
	// orderRoutes.OrderRouter(router)

	// router.Run(":3000")

	y, z := stdOTP("abdellah@mail.com", "1234")
	fmt.Println("asdf")
}
