package user_repositories

import (
	"crypto/md5"
	"delivery/Services/Auth/Domain/DTOs"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	"errors"
	"fmt"
)

type AuthRepository struct {
}

func (obj *AuthRepository) Login(login string, password string) error {
	db := shareddb.NewDB()
	defer db.Close()
	dto := &DTOs.LoginDTO{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
	if dto.Failed() {
		return errors.New("invalid credential")
	}
	return nil
}

func (obj *AuthRepository) Register(
	userName string,
	fullName string,
	email string,
	password string,
	address string,
	userType string) error {

	db := shareddb.NewDB()
	defer db.Close()

	hash := md5.Sum([]byte(password))
	hashedPass := fmt.Sprintf("%x", hash)

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
