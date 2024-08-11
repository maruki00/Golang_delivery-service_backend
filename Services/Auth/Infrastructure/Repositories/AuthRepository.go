package user_repositories

import (
	"crypto/md5"
	auth_domain_dto "delivery/Services/Auth/Domain/DTOs"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/go-uuid"
)

type AuthRepository struct {
}

func (obj *AuthRepository) Login(login string, password string) (*auth_domain_dto.AuthDTO, error) {
	db := shareddb.NewDB()
	defer db.Close()
	dto := &auth_domain_dto.AuthDTO{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT id, email, user_type, user_level FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.User_id, &dto.Email, &dto.User_type, &dto.User_level)
	if dto.Email == "" || dto.User_level == "" || dto.User_type == "" {
		return nil, errors.New("invalid credentials")
	}
	dto.Token = base64.NewEncoding("base64").EncodeToString([]byte(uuid.New))

	return nil, nil
}

func (obj *AuthRepository) ForgetPassword(email string) error {

	db := shareddb.NewDB()
	defer db.Close()

	statement, err := db.Prepare(`
				INSERT INTO users (user_name, full_name, email, password, address, user_type) 
				VALUES (?, ?, ?, ?, ?, ?)`)

	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(userName, fullName, email, hashedPass, address, userType)
	if err != nil {
		return errors.New("could not register a new account")
	}
	return nil
}
