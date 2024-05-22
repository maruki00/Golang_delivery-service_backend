package Authusecases

import (
	"delivery/Services/Auth/Domain/DTOs"
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
)

type LoginuserCase struct {
	repo *user_repositories.AuthRepository
}

func (obj *LoginuserCase) Login(dto DTOs.LoginDTO) string {

	err := obj.repo.Login(dto.GetLogin(), dto.GetPassword())

	if err != nil {
		return "invalid credentials"
	}
	return "success"
}
