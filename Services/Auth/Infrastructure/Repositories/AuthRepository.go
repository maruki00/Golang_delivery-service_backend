package auth_infrastructure_repository

import (
	"crypto/md5"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"errors"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type AuthRepository struct {
	sync.Mutex
	db *gorm.DB
}

func NewAuthRepository(config shared_configs.Config) *AuthRepository {
	return &AuthRepository{
		db: shareddb.NewMysqlDB_GORM(&config),
	}
}

func (obj *AuthRepository) CheckToken(token string) bool {

	return false
}

func (obj *AuthRepository) generateToken(dto *auth_domain_dtos.AuthDTO) string {

	return ""
}

func (obj *AuthRepository) Login(login, password string) (*auth_domain_dtos.AuthDTO, error) {

	dto := &auth_domain_dtos.AuthDTO{}
	//dto.Token = "helloworld" //base64.NewEncoding("base64").EncodeToString([]byte("asfdf"))
	hash := md5.Sum([]byte(password))
	hashedPassword := fmt.Sprintf("%x", hash)
	uu := &shared_models.User{}
	u := obj.db.Model(&shared_models.User{}).Where("user_name", login).Where("password", hashedPassword).Limit(1).Find(&uu)
	if u.RowsAffected == 0 {
		return nil, errors.New("invalid credentials")
	}
	err := statement.QueryRow(login, h).Scan(&dto.User_id, &dto.Email, &dto.User_type, &dto.User_level)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("invalid creadentails " + err.Error())
	}

	if dto.Email == "" || dto.User_level == "" || dto.User_type == "" {
		return nil, errors.New("invalid credentials")
	}
	return dto, nil
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
