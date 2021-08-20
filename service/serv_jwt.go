package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/dgrijalva/jwt-go"
)

type DashboardClaims struct {
	jwt.StandardClaims
	Name  string
	Email string
}

// generate token with claims and key
func GenerateToken(claim jwt.Claims, key string) (string, error) {
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaim.SignedString([]byte(key))
	if err != nil {
		log.Printf("Key:%s, Error:%s", key, err.Error())
	}
	return token, err
}

// parse string token to claims with specific key
func ParseToken(token string, key string) (*DashboardClaims, error) {
	var err error
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&DashboardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		// check the detail of error
		debugTokenError(err)
		return nil, err
	}
	// try convert jwt.Claims to DashboardClaims
	// type assertion ref: (https://blog.kalan.dev/golang-type-assertion/)
	if claims, ok := tokenClaims.Claims.(*DashboardClaims); ok {
		return claims, nil
	}
	return nil, errors.New("JWT token payload is improper")
}

func debugTokenError(err error) {
	var message string
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			message = "token is malformed"
		} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
			message = "token could not be verified because of signing problems"
		} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			message = "signature validation failed"
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			message = "token is expired"
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			message = "token is not yet valid before sometime"
		} else {
			message = "can not handle this token"
		}
	}
	log.Println(message)
}

func GenerateClaims(user models.User, issuer string) jwt.Claims {
	now := time.Now()
	jwtId := user.Name + strconv.FormatInt(now.Unix(), 10)
	claims := DashboardClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Name,
			ExpiresAt: now.Add(24 * time.Hour).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    issuer,
			NotBefore: now.Unix(),
			Subject:   user.Name,
		},
		Name:  user.Name,
		Email: user.Email,
	}
	return claims
}
