package auth_infrastructure_repository

import (
	auth_domain_entities "delivery/Services/Auth/Domain/Entities"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shared_entities "delivery/Services/Shared/Domain/Entities"
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

func (obj *AuthRepository) Register(user shared_entities.UserEntity) (shared_entities.UserEntity, error) {

	u := obj.db.Model(&shared_models.User{}).Create(user)
	if u.RowsAffected == 0 {
		return nil, errors.New("could not create the user")
	}
	return user, nil
}

func (obj *AuthRepository) TwoFactoryCreate(twofactory *auth_domain_entities.TwoFactoryPinEntity) (bool, error) {

	u := obj.db.Model(&auth_infrastructure_models.TwoFactoryPin{}).Create(twofactory)
	if u.RowsAffected == 0 {
		return false, errors.New("could not create the user")
	}
	return true, nil
}

func (obj *AuthRepository) TwoFactoryConfirm(email string, pin int) (bool, error) {

	var twofactory *auth_infrastructure_models.TwoFactoryPin
	obj.db.Model(&twofactory).Where("pin = ? and email = ? ", pin, email).Find(twofactory)
	if twofactory == nil {
		return false, errors.New("could not confirm your account")
	}
	res := obj.db.Model(&twofactory).Where("pin = ? and email = ? ", pin, email).Delete(twofactory)
	if res.RowsAffected == 0 {
		return false, fmt.Errorf("could not confirm your account")
	}
	return true, nil
}

// func (obj *AuthRepository) ForgetPassword(entity *auth_domain_entities.ForgetPasswordEntity) error {
// 	res := obj.db.Model(&auth_infrastructure_models.ResetPassword{}).Create(entity)
// 	if res.RowsAffected == 0 {
// 		return errors.New("record doesnt saved")
// 	}
// 	return nil
// }

// func (obj *AuthRepository) ResetPasswordByPin(entity *auth_domain_entities.ForgetPasswordEntity, password string) error {
// 	var resetPass *auth_infrastructure_models.ResetPassword
// 	obj.db.Model(&auth_infrastructure_models.ResetPassword{}).Where("pin = ? or token = ?", entity.GetPin(), entity.GetToken()).Find(resetPass)
// 	if resetPass == nil {
// 		return errors.New("invalid token or pin")
// 	}
// 	var user shared_models.User

// 	obj.db.Model(&user).Where("email = ? ", entity.GetEmail()).Find(&user)
// 	if user == nil {
// 		return errors.New("invalid token or pin")
// 	}

// 	res := obj.db.Model(&user).Update("password", shared_utils.Md5Hash(password))

// 	if res.RowsAffected == 0 {
// 		return errors.New("record doesnt updated")
// 	}
// 	obj.clearToken(user.Id)
// 	return nil
// }
