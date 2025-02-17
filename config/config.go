package config

import (
	"game-app-go/repository/mysql"
	"game-app-go/service/authservice"
)

type HTTPServer struct {
	Port int
}

type Config struct{
	HTTPServer HTTPServer
	Auth authservice.Config 
	Mysql mysql.Config
}
