package user_repositories

import (
	"crypto/md5"
	auth_domain_dto "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	"errors"
	"fmt"
)

type AuthRepository struct {
}

func (obj *AuthRepository) Login(login string, password string) (*auth_domain_dto.AuthDTO, error) {
	db := shareddb.NewDB()
	defer db.Close()
	entity := &auth_infrastructure_models.AuthModel{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT id, email, password, user_type FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.id)
	if entity.Failed() {
		return nil, errors.New("invalid credential")
	}
	return nil
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
