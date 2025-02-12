package main

import (
	"encoding/json"
	"fmt"
	"game-app-go/entity"
	"game-app-go/repository/mysql"
	"game-app-go/service/userservice"
	"io"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/user/register", userRegisterHandler)
	http.HandleFunc("/user/login", userLoginHandler)
	log.Print("serever is running on port 8080")
	http.ListenAndServe(":8080", nil)
}


func userRegisterHandler(writer http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"invalid method"}`)
	}

	data, err := io.ReadAll(req.Body)

	if err != nil {
		writer.Write(
			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
		)
	}

	var uReq userservice.RegisterRequest

	err = json.Unmarshal(data, &uReq)

	if err != nil {
		writer.Write(
			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
		)
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)

	_, err = userSvc.Register(uReq)

	if err != nil {
		writer.Write(
			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
		)
		return
	}

	writer.Write([]byte(`{message:"user created"}`))
}

func userLoginHandler(writer http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost{
		fmt.Fprintf(writer, `{"error":"invalid method"}`)
	}
	data, err := io.ReadAll(req.Body)

	if err != nil {
		writer.Write(
			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
		)
	}

	var lReq userservice.LoginRequest

	err = json.Unmarshal(data, &lReq)

	if err != nil {
		writer.Write(
			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
		)
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)

	_, err = userSvc.Login(lReq)

	if err != nil {
		writer.Write(
			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
		)
		return
	}

	writer.Write([]byte(`{"messge": ""user credential is ok}`))

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