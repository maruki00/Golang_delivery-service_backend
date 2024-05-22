package Authusecases

import (
	"delivery/Services/Auth/Domain/DTOs"
	shareddb "delivery/Services/Sahared/Infrastructure/DB"
)

type LoginuserCase struct {
	db *shareddb.DBHandler
}

func (obj *LoginuserCase) Login(dto *DTOs.LoginDTO) string {
	data := []any{dto.getLogin(), dto.getPassword()}
	err := obj.db.Query("select * from users where username=? and password=?", data)
	if err != nil {
		return "Could not Login...."
	}
	return "Success"
}
