package authu_services

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
)

var repo *user_repositories.AuthRepository

func Login(dto auth_domain_dtos.LoginDTO) (string, error) {

	_, err := repo.Login(dto.GetLogin(), dto.GetPassword())
	if err != nil {
		return "invalid credentials", nil
	}
	return "success", nil
}
