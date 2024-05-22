package shareddb

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBHandler struct {
	db *sql.DB
}

func (obj *DBHandler) NewDB() *DBHandler {
	if obj.db == nil {
		db, er := sql.Open("mysql", "user:P@ssW0rd@tcp://127.0.0.1:3306/delivery")
		if er != nil {
			panic("Error With Database!")
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		obj.db = db
	}
	return obj
}

func (obj *DBHandler) Query(sql string, params []any) error {
	obj.NewDB()
	_, err := obj.db.Exec(sql, params...)
	if err != nil {
		return err
	}
	return nil
}
