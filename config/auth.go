package config

import "time"

var (
	day                   = time.Hour * 24
	secretKey             = "my-secret-key"
	accessExpirationTime  = time.Now().Add(day * 7).Unix()
	refreshExpirationTime = time.Now().Add(day * 30).Unix()
)

type AuthConfig struct {
	signKey               []byte
	accessExpirationTime  int64
	refreshExpirationTime int64
}

func (authConfig AuthConfig) New() AuthConfig {
	authConfig.signKey = []byte(secretKey)
	authConfig.accessExpirationTime = accessExpirationTime
	authConfig.refreshExpirationTime = refreshExpirationTime
	return authConfig
}
