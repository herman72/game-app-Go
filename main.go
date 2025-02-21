package main

import (
	"fmt"
	"game-app-go/config"
	"game-app-go/delivery/httpserver"
	"game-app-go/repository/mysql"
	"game-app-go/service/authservice"
	"game-app-go/service/userservice"
	"game-app-go/validator/uservalidator"
	"time"
)

const (
	JWTSignKey = "jwt_secret"
	AccessTokenSubject = "at"
	RefreshTokenSubject = "rt"
	AccessTokenExpirationDuration = time.Hour * 24
	RefreshTokenExpirationDuration = time.Hour * 24 * 7

)

func main(){
	

	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey: JWTSignKey,
			AccessExpirationTime: AccessTokenExpirationDuration,
			RefresExpirationTime: RefreshTokenExpirationDuration,
			AccessSubject: AccessTokenSubject,
			RefreshSubject: RefreshTokenSubject, 
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameapp",
			Port: 3308,
			Host: "localhost",
			DBName: "gameapp_db",
		},
	}
	// TODO: add command for up, down and status
	// mgr :=migrator.New(cfg.Mysql)
	// mgr.Up()
	
	authSvc, userSvc, userValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)
	fmt.Println("Start Echo server")
	server.Serve()

	// http.HandleFunc("/user/login", userLoginHandler)
	// http.HandleFunc("/user/profile",  userProfileHandler)
	// log.Print("serever is running on port 8080")
	// http.ListenAndServe(":8080", nil)
}


func setupServices(cfg config.Config)(authservice.Service, userservice.Service,
	uservalidator.Validator){
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userSvc := userservice.New(authSvc, MysqlRepo)
	uV := uservalidator.New(MysqlRepo)

	return authSvc, userSvc, uV

}