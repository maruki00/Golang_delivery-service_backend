package Authusecases

import (
	"delivery/Services/Auth/Domain/DTOs"
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
)

type LoginuserCase struct {
	repo *user_repositories.AuthRepository
}

func (obj *LoginuserCase) Login(dto DTOs.LoginDTO) string {
	err := obj.repo.Loginn(dto.GetLogin(), dto.GetPassword())
	if err != nil {
		return err.Error()
	}
	return "success"
}
