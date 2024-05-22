package user_repositories

import (
	shareddb "delivery/Services/Sahared/Infrastructure/DB"
	"errors"
)

type AuthRepository struct {
	shareddb.DBHandler
}

func (obj AuthRepository) login(login string, password string) error {
	data := []any{login, password}
	err := obj.Query("select * from users where username=? and password=?", data)
	if err != nil {
		return errors.New("invalid operation")
	}
	return nil
}
