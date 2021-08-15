package controller

import (
	"log"
	"net/http"

	"github.com/Cheep2Workshop/proj-web/dashredis"
	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/go-redis/redis/v8"

	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	// step 1 check accont
	// step 2 modify info
	// step 3 save user
	// setp F ok
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
	user, err := repo.Client.BeginLogin(*req)
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
		repo.Client.CancelLogin()
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	// step3: save log
	err = repo.Client.SaveLoginLog(req.Email)
	if err != nil {
		log.Println(err.Error())
		repo.Client.CancelLogin()
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}

	repo.Client.EndLogin()
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetLoginLogs(ctx *gin.Context) {
	var email string
	err := ctx.Bind(&email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	// try get cache from redis
	redisClient := dashredis.NewRedisClient()
	logs, err := redisClient.GetLoginLogs(email)
	if err != nil && err != redis.Nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	if err == nil {
		ctx.JSON(http.StatusOK, logs)
		return
	}

	logs, err = repo.Client.GetLoginLogs(email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	// cache result into redis
	redisClient.SetLoginLogs(email, logs...)
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

	// delete cache
	redisClient := dashredis.NewRedisClient()
	redisClient.DeleteWithEmail(req.DeleteEmail)
	ctx.JSON(http.StatusOK, req.DeleteEmail)
}
