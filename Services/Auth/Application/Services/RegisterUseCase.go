package authu_services

import (
	user_repositories "delivery/Services/Auth/Infrastructure/Repositories"
)

type RegisteruserCase struct {
	repo *user_repositories.AuthRepository
}

// userName string, fullName string, email string, password string, country string, city string, street string, house int, flat int
// func (obj *LoginuserCase) Register(dto auth_domain_dto.DTO) (string, error) {

// 	fmt.Println(dto)
// 	_ = obj.repo.Register(
// 		dto.UserName,
// 		dto.FullName,
// 		dto.Email,
// 		dto.Password,
// 		dto.Address,
// 		dto.UserType,
// 	)

// 	// fmt.Print("Hello world")
// 	// if err != nil {
// 	// 	return "invalid Data"
// 	// }
// 	return "success"
// }
