package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/orm"
	"github.com/Cheep2Workshop/proj-web/utils"
	"github.com/gin-gonic/gin"
)

const (
	Issuer = "Coinmouse"
	Secret = "SsdDifdoDz"
)

func CheckJwt(token string) (bool, error) {
	// parse token
	claims, err := utils.ParseToken(token, Secret)
	if err != nil {
		return false, err
	}
	// check name and email matched
	_, err = orm.Client.CheckAuth(claims.Name, claims.Email)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateJWT(user models.User) (string, error) {
	claims := utils.GenerateClaims(user, Issuer)
	token, err := utils.GenerateToken(claims, Secret)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return token, nil
}

// (middleware) authorizate the jwt
func AuthJwt(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	splited := strings.Split(auth, "Bearer ")
	if len(splited) != 2 {
		// abort
		ctx.AbortWithStatusJSON(http.StatusForbidden, "No jwt in header.")
		return
	}
	// parse token
	token := splited[1]
	result, err := CheckJwt(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	if !result {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Access not incorrect.")
	}
	ctx.Next()
}
