package main

import (
	"fmt"
	"game-app-go/config"
	"game-app-go/delivery/httpserver"
	"game-app-go/repository/mysql"
	"game-app-go/service/authservice"
	"game-app-go/service/userservice"
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
	fmt.Println("Start Echo server")

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
	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()

	// http.HandleFunc("/user/login", userLoginHandler)
	// http.HandleFunc("/user/profile",  userProfileHandler)
	// log.Print("serever is running on port 8080")
	// http.ListenAndServe(":8080", nil)
}


// func userLoginHandler(writer http.ResponseWriter, req *http.Request){
// 	if req.Method != http.MethodPost{
// 		fmt.Fprintf(writer, `{"error":"invalid method"}`)
// 	}
// 	data, err := io.ReadAll(req.Body)

// 	if err != nil {
// 		writer.Write(
// 			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
// 		)
// 	}

// 	var lReq userservice.LoginRequest

// 	err = json.Unmarshal(data, &lReq)

// 	if err != nil {
// 		writer.Write(
// 			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
// 		)
// 		return
// 	}
// 	authSvc := authservice.New(JWTSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpirationDuration, RefreshTokenExpirationDuration)
// 	mysqlRepo := mysql.New()
// 	userSvc := userservice.New(authSvc ,mysqlRepo)

// 	resp, err := userSvc.Login(lReq)

// 	if err != nil {
// 		writer.Write(
// 			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
// 		)
// 		return
// 	}

// 	data, err = json.Marshal(resp)
// 	if err != nil {
// 		writer.Write(
// 			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
// 		)
// 		return
// 	}

// 	writer.Write(data)


// 	// writer.Write([]byte(`{"messge": ""user credential is ok}`))

// }

// func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodGet{
// 		fmt.Fprintf(writer, `{"error":"invalid method"}`)
// 	}

// 	// sessionID := req.Header.Get("SessionID")
// 	// TODO: Validate sessionid by database and get user id

// 	// validate jwt token and retrive userID from pyload

// 	authSvc := authservice.New(JWTSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpirationDuration, RefreshTokenExpirationDuration)

// 	authToken := req.Header.Get("Authorization")
// 	claims, err := authSvc.ParseToken(authToken)
// 	if err != nil {
// 		fmt.Fprintf(writer, `{"error":"token is not valid"}`)
// 		return
// 	}
// 	mysqlRepo := mysql.New()
// 	userSvc := userservice.New(authSvc ,mysqlRepo)

// 	resp, err := userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})

// 	if err != nil {
// 		writer.Write(
// 			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
// 		)
// 		return
// 	}

// 	data, err := json.Marshal(resp)
// 	if err != nil {
// 		writer.Write(
// 			[]byte(fmt.Sprintf(`{error: %s}`, err.Error())),
// 		)
// 		return
// 	}

// 	writer.Write(data)


// }

// func testUserMysqlRepo() {
// 	mysqlRepo := mysql.New()
// 	createdUser, err := mysqlRepo.Register(entity.User{
// 		ID:          0,
// 		PhoneNumber: "09273423",
// 		Name:        "Mohammad",
// 	})

// 	if err != nil {
// 		fmt.Println("register user", err)
// 	} else {
// 		fmt.Println("created user", createdUser)
// 	}
// 	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
// 	if err != nil {
// 		fmt.Println("unique err", err)
// 	}

// 	fmt.Println("isUnique", isUnique)
// }


func setupServices(cfg config.Config)(authservice.Service, userservice.Service){
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userSvc := userservice.New(authSvc, MysqlRepo)

	return authSvc, userSvc

}