package controller

import (
	"log"
	"net/http"

	"github.com/Cheep2Workshop/proj-web/models"

	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var err error
	var user *models.User
	err = ctx.Bind(&user)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	result, err := repo.Client.Signup(*user)
	if result {
		ctx.JSON(http.StatusOK, result)
		return
	}
	log.Println(err.Error())
	ctx.JSON(http.StatusBadRequest, err.Error())
}

func SetUser(ctx *gin.Context) {
	var err error
	var req *repo.SetUserReq
	err = ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
	}
	err = repo.Client.SetUser(*req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
	}
}

func Login(ctx *gin.Context) {
	var err error
	var req *repo.LoginReq
	err = ctx.Bind(&req)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	// step1: get login user
	user, err := repo.Client.Login(*req)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	// step2: generate token
	mgr := JWTManager{
		Issuer: Issuer,
		Secret: Secret,
	}
	token, err := mgr.GenerateJwt(*user)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	// step3: save log
	err = repo.Client.SaveLoginLog(req.Email)
	if err != nil {
		log.Println(err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetLoginLogs(ctx *gin.Context) {
	var email *string
	err := ctx.Bind(&email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	logs, err := repo.Client.GetLoginLogs(*email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, logs)
}

// Search and get user info by email
func GetUser(ctx *gin.Context) {
	var err error
	var email *string
	err = ctx.Bind(&email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
	}
	user, err := repo.Client.GetUserByEmail(*email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
	}
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	var err error
	var req *repo.DeleteUserReq
	err = ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	err = repo.Client.DeleteUser(*req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, req.DeleteEmail)
}
