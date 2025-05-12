package config

import "time"

const (
	JWTSignKey = "jwt_secret"
	AccessTokenSubject = "at"
	RefreshTokenSubject = "rt"
	AccessTokenExpirationDuration = time.Hour * 24
	RefreshTokenExpirationDuration = time.Hour * 24 * 7
	AuthMiddlewareContextKey = "claim"

)