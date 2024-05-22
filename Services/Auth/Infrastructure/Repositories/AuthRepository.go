package user_repositories

import (
	shareddb "delivery/Services/Sahared/Infrastructure/DB"
)

type AuthRepository struct {
}

func (obj *AuthRepository) Loginn(login string, password string) error {
	db := shareddb.NewDB()
	// data := []any{login, password}
	// db.Query("INSERT INTO users VALUES(1, 'sfda','asdf','asdf','asdf')")
	db.Exec("INSERT INTO users VALUES(10, 'sfda','asdf','asdf','asdf')", nil)
	return nil
}
