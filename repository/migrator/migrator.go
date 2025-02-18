package migrator

import (
	"database/sql"
	"fmt"
	"game-app-go/repository/mysql"

	migrate "github.com/rubenv/sql-migrate"
)

// TODO: set migration table name
// TODO: add limit for up and down

type Migrator struct {
	dialect string
	dbConfig mysql.Config
	migratations *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config,)Migrator{
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	} 

	return Migrator{dialect: "mysql" ,dbConfig: dbConfig ,migratations: migrations}
}

func (m Migrator)Up(){
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.dbConfig.Username, 
	m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))

	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migratations, migrate.Up)
	if err !=nil {
		panic(fmt.Errorf("cant apply migrations: %v", err))
	}

	fmt.Printf("Applied %d migrations\n", n)
}

func (m Migrator)Down(){

	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.dbConfig.Username, 
	m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))

	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migratations, migrate.Down)
	if err !=nil {
		panic(fmt.Errorf("cant rollback migrations: %v", err))
	}

	fmt.Printf("rollback %d migrations\n", n)

}

func (m Migrator)Status(){
	// TODO: add status
}

