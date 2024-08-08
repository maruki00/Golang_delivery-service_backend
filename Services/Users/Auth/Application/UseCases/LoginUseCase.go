package Authusecases

import (
	"delivery/Services/Auth/Domain/DTOs"
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
	"fmt"
)

type LoginuserCase struct {
	repo *user_repositories.AuthRepository
}

func (obj *LoginuserCase) Login(dto DTOs.LoginDTO) string {

	fmt.Println("hello : ", dto)
	err := obj.repo.Login(dto.GetLogin(), dto.GetPassword())
	fmt.Println("Error result: ", err)
	if err != nil {
		return "invalid credentials"
	}
	return "success"
}
