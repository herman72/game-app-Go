package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port int `koanf:"port"`
	Host string `koanf:"host"`
	DBName string `koanf:"db_name"`
}

type MySQLDB struct {
	config Config
	db *sql.DB
}

func New(config Config)*MySQLDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", config.Username, 
	config.Password, config.Host, config.Port, config.DBName))

	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v",err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return &MySQLDB{config: config, db: db}
}