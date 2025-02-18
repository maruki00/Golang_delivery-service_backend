package contracts

import (
	auth_domain_entities "delivery/internal/auth/domain/entities"
	auth_infra_models "delivery/internal/auth/infra/Models"
	shared_entities "delivery/internal/shared/Domain/Entities"
	shared_models "delivery/internal/shared/infra/Models"
)

type IAuthRepository interface {
	// Login(login, password string) (*auth_domain_dtos.AuthDTO, error)
	CheckToken(token string) bool
	Login(login, password string) (*shared_models.User, error)
	CreateAuth(token string, user *shared_models.User) *auth_infra_models.Auth
	ClearToken(id int)
	Register(user shared_entities.UserEntity) (shared_entities.UserEntity, error)
	CleanPins(email string)
	TwoFactoryCreate(twofactory auth_domain_entities.TwoFactoryPinEntity) (bool, error)
	TwoFactoryConfirm(email string, pin int) (bool, error)
	LockUser(email string, lock string)
	Logout(token string)
}
