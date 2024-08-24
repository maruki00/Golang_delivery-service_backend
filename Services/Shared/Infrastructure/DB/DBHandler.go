package shared_infrastructure_db

import (
	"database/sql"
	shared_configs "delivery/Services/Shared/Application/Configs"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *sql.DB {
	config, err := shared_configs.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@/%s",
		config.Database.Mysql.Username,
		config.Database.Mysql.Password,
		config.Database.Mysql.DBName)
	db, err := sql.Open(config.Database.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Second * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func NewMysqlDB_GORM(config *shared_configs.Config) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Mysql.Username,
		config.Database.Mysql.Password,
		config.Database.Mysql.Host,
		config.Database.Mysql.Port,
		config.Database.Mysql.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
