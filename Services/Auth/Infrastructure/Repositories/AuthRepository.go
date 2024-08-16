package auth_infrastructure_repository

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"fmt"
	"sync"

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

func (obj *AuthRepository) generateToken(dto *auth_domain_dtos.AuthDTO) string {

	return ""
}

func (obj *AuthRepository) Login(login, password string) (*auth_domain_dtos.LoggedInDTO, error) {

	hashedPassword := shared_utils.Md5Hash(password)
	uu := &shared_models.User{}
	_ = obj.db.Model(&shared_models.User{}) //.Where("user_name", login).Where("password", hashedPassword).Limit(1).Find(nil)
	// if u.RowsAffected == 0 {
	// 	return nil, errors.New("invalid credentials")
	// }
	// token, err := shared_utils.JwtToken(uu.Email, uu.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not generate token")
	// }

	fmt.Println(hashedPassword, uu)
	return &auth_domain_dtos.LoggedInDTO{
		Token: "",
	}, nil
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
