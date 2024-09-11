package auth_infrastructure_repository

import (
	auth_domain_entities "delivery/Services/Auth/Domain/Entities"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	shared_entities "delivery/Services/Shared/Domain/Entities"
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

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (obj *AuthRepository) CheckToken(token string) bool {

	return false
}

func (obj *AuthRepository) Login(login, password string) (*shared_models.User, error) {

	var uu shared_models.User

	obj.db.Model(&shared_models.User{}).Where("user_name", login).Where("password", password).Find(&uu)

	fmt.Println(" hello : ", uu, password)

	return &uu, nil
}

func (obj *AuthRepository) CreateAuth(token string, user *shared_models.User) *auth_infrastructure_models.Auth {
	auth := &auth_infrastructure_models.Auth{
		Email:     user.Email,
		Token:     token,
		UserId:    user.Id,
		UserType:  user.UserType,
		UserLevel: user.UserLevel,
		ExpiresAt: string(time.Now().Format("y-m-d H:i:s")),
	}

	obj.db.Create(auth)
	return auth

}
func (obj *AuthRepository) ClearToken(id int) {
	obj.db.Exec("delete from auths  where user_id = ? ", id)
}

func (obj *AuthRepository) Register(user shared_entities.UserEntity) (shared_entities.UserEntity, error) {

	u := obj.db.Model(&shared_models.User{}).Create(user)
	if u.RowsAffected == 0 {
		return nil, errors.New("could not create the user")
	}
	return user, nil
}

func (obj *AuthRepository) CleanPins(email string) {
	obj.db.Exec("delete from two_factory_pins where email = ? ", email)
}

func (obj *AuthRepository) TwoFactoryCreate(twofactory auth_domain_entities.TwoFactoryPinEntity) (bool, error) {

	u := obj.db.Model(&auth_infrastructure_models.TwoFactoryPin{}).Create(twofactory)
	if u.RowsAffected == 0 {
		return false, errors.New("could not create the user")
	}
	return true, nil
}

func (obj *AuthRepository) TwoFactoryConfirm(email string, pin int) (bool, error) {

	var twofactory *auth_infrastructure_models.TwoFactoryPin
	obj.db.Model(&auth_infrastructure_models.TwoFactoryPin{}).Where("pin", pin).Where("email", email).Find(&twofactory)
	if twofactory == nil {
		return false, errors.New("invalid data")
	}
	obj.db.Exec("delete from two_factory_pins where pin = ? and email = ? ", pin, email)
	return true, nil
}

func (obj *AuthRepository) LockUser(email string, lock string) {
	obj.db.Exec("update users set is_locked = ? where (email = ? or user_name = ?)", lock, email, email)
}

func (obj *AuthRepository) Logout(token string) {
	obj.db.Exec("delete from auths where token = ? ", token)
}
