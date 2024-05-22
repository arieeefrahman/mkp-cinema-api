package middlewares

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4/middleware"
)

var ctx = context.Background()

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
	RedisClient     *redis.Client
}

func NewConfigJWT(secretJWT string, expiresDuration int, redisAddr string) *ConfigJWT {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &ConfigJWT{
		SecretJWT:       secretJWT,
		ExpiresDuration: expiresDuration,
		RedisClient:     redisClient,
	}
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(userID uuid.UUID) (string, error) {
	tokenClaims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))
	if err != nil {
		return "", err
	}

	// Store token in Redis
	err = jwtConf.RedisClient.Set(ctx, token, userID.String(), time.Duration(jwtConf.ExpiresDuration)*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (jwtConf *ConfigJWT) CheckToken(token string) bool {
	_, err := jwtConf.RedisClient.Get(ctx, token).Result()
	return err == nil
}

func (jwtConf *ConfigJWT) GetPayload(token *jwt.Token) *JwtCustomClaims {
	claims := token.Claims.(*JwtCustomClaims)
	return claims
}

func (jwtConf *ConfigJWT) Logout(token string) bool {
	err := jwtConf.RedisClient.Del(ctx, token).Err()
	return err == nil
}
