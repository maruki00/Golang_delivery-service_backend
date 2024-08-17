package auth_infrastructure_repository

import (
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"errors"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

type AuthRepository struct {
	sync.Mutex
	db *gorm.DB
}

func NewAuthRepository(config *shared_configs.Config) *AuthRepository {
	return &AuthRepository{
		db: shareddb.NewMysqlDB_GORM(config),
	}
}

func (obj *AuthRepository) CheckToken(token string) bool {

	return false
}

// func (obj *AuthRepository) generateToken(dto *auth_domain_dtos.AuthDTO) string {

// 	return ""
// }

func (obj *AuthRepository) Login(login, password string) (string, error) {

	hashedPassword := shared_utils.Md5Hash(password)
	uu := &shared_models.User{}

	u := obj.db.Model(&shared_models.User{}).Where("user_name", login).Where("password", hashedPassword).Limit(1).Find(uu)
	if u.RowsAffected == 0 {
		return "", errors.New("invalid credentials")
	}

	token, err := shared_utils.JwtToken(uu.Email, uu.Id)
	auth := &auth_infrastructure_models.Auth{
		Email:     uu.Email,
		Token:     token,
		UserId:    uu.Id,
		UserType:  uu.UserType,
		UserLevel: uu.UserLevel,
		ExpiresAt: string(time.Now().Format("y-m-d H:i:s")),
	}
	obj.clearToken(auth.UserId)
	obj.Create(auth)
	if err != nil {
		return "", fmt.Errorf("could not generate token " + err.Error())
	}

	return token, nil
}

func (obj *AuthRepository) clearToken(id int) {
	obj.db.Exec("delete from auths  where user_id = ? ", id)
}

func (obj *AuthRepository) Create(auth *auth_infrastructure_models.Auth) *auth_infrastructure_models.Auth {
	obj.db.Create(auth)
	return auth
}

func (obj *AuthRepository) ForgetPassword(email string) error {

	// db := shareddb.NewDB()
	// defer db.Close()

	// statement, err := db.Prepare(`
	// 			INSERT INTO users (user_name, full_name, email, password, address, user_type)
	// 			VALUES (?, ?, ?, ?, ?, ?)`)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer statement.Close()

	// _, err = statement.Exec(userName, fullName, email, hashedPass, address, userType)
	// if err != nil {
	// 	return errors.New("could not register a new account")
	// }
	return nil
}

func (obj *AuthRepository) Register(user shared_models.User) error {

	u := obj.db.Model(&shared_models.User{}).Create(user)
	if u.RowsAffected == 0 {
		return errors.New("could not create the user")
	}
	return nil
}
