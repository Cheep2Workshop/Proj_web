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

type JWTManager struct {
	Issuer string
	Secret string
}

func (mgr *JWTManager) VerifyJwt(token string) (*utils.DashboardClaims, error) {
	// parse token
	claims, err := utils.ParseToken(token, mgr.Secret)
	if err != nil {
		return nil, err
	}
	// check name and email matched
	_, err = orm.Client.CheckAuth(claims.Name, claims.Email)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (mgr *JWTManager) GenerateJwt(user models.User) (string, error) {
	claims := utils.GenerateClaims(user, mgr.Issuer)
	token, err := utils.GenerateToken(claims, mgr.Secret)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return token, nil
}

// (middleware) authorizate the jwt
func (mgr *JWTManager) AuthJwt(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	splited := strings.Split(auth, "Bearer ")
	if len(splited) != 2 {
		// abort
		ctx.AbortWithStatusJSON(http.StatusForbidden, "No jwt in header.")
		return
	}
	// parse token
	token := splited[1]
	_, err := mgr.VerifyJwt(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.Next()
}
