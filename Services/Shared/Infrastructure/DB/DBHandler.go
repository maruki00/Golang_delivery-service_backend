package shared_infrastructure_db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "user:user@/delivery")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Second * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
