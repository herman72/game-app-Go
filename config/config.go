package config

import (
	"game-app-go/repository/mysql"
	"game-app-go/service/authservice"
)

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct{
	HTTPServer HTTPServer `koanf:"http_server"`
	Auth authservice.Config `koanf:"auth"`
	Mysql mysql.Config `koanf:"mysql"`
}
