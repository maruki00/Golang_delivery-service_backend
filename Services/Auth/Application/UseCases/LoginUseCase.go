package Authusecases

import (
	shareddb "delivery/Services/Sahared/Infrastructure/DB"
)

type LoginuserCase struct {
	db *shareddb.DBHandler
}

func (obj *LoginuserCase) Login(login string, password string) string {
	data := []any{login, password}
	err := obj.db.Query("select * from users where username=? and password=?", data)
	if err != nil {
		return "Could not Login...."
	}
	return "Success"
}
