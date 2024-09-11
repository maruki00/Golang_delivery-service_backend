package auth_domain_contracts

import (
	auth_domain_entities "delivery/Services/Auth/Domain/Entities"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	shared_entities "delivery/Services/Shared/Domain/Entities"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
)

type IAuthRepository interface {
	// Login(login, password string) (*auth_domain_dtos.AuthDTO, error)
	CheckToken(token string) bool
	Login(login, password string) (*shared_models.User, error)
	CreateAuth(token string, user *shared_models.User) *auth_infrastructure_models.Auth
	ClearToken(id int)
	Register(user shared_entities.UserEntity) (shared_entities.UserEntity, error)
	CleanPins(email string)
	TwoFactoryCreate(twofactory auth_domain_entities.TwoFactoryPinEntity) (bool, error)
	TwoFactoryConfirm(email string, pin int) (bool, error)
	LockUser(email string, lock string)
	Logout(token string)
}
