package main

import (
	"fmt"
	"game-app-go/entity"
	"game-app-go/repository/mysql"
)

func main(){

	testUserMysqlRepo()
}


func testUserMysqlRepo() {
	mysqlRepo := mysql.New()
	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "09273423",
		Name:        "Mohammad",
	})

	if err != nil {
		fmt.Println("register user", err)
	} else {
		fmt.Println("created user", createdUser)
	}
	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique err", err)
	}

	fmt.Println("isUnique", isUnique)
}