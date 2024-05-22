package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// router := gin.Default()
	// routes.AuthRouter(router)

	db, err := sql.Open("mysql", "user:P@ssW0rd@tcp(localhost:3306)/delivery")

	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.Query("insert into users values(1, 'asdfasd', 'asdf', 'zxdf')", nil)
	fmt.Println("Server run on port 3000")
	//router.Run(":3000")
}
