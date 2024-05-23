package Authusecases

import (
	"delivery/Services/Auth/Domain/DTOs"
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
)

type RegisteruserCase struct {
	repo *user_repositories.AuthRepository
}

// userName string, fullName string, email string, password string, country string, city string, street string, house int, flat int
func (obj *LoginuserCase) Register(dto DTOs.UserDTO) string {

	err := obj.repo.Register(
		dto.GetUserName(),
		dto.GetFullName(),
		dto.GetEmail(),
		dto.GetPassword(),
		dto.GetCountry(),
		dto.GetCity(),
		dto.GetStreet(),
		dto.GetHouse(),
		dto.GetFlat(),
	)

	if err != nil {
		return "invalid credentials"
	}
	return "success"
}
