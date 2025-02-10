package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New()*MySQLDB {
	db, err := sql.Open("mysql", "gameapp:gameapp@(localhost:3308)/gameapp_db")

	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v",err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return &MySQLDB{db}
}