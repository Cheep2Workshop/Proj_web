package controller

import (
	"github.com/gin-gonic/gin"
)

func RunGin(i chan int) {
	// run gin server (http)
	r := gin.Default()

	mgr := JWTManager{
		Issuer: Issuer,
		Secret: Secret,
	}
	r.Any("/health", Health)
	r.POST("/signup", Signup)
	r.POST("/setuser", mgr.AuthJwt, SetUser)
	r.POST("/login", Login)
	r.POST("/getloginlogs", mgr.AuthJwt, GetLoginLogs)
	r.POST("/getuser", GetUser)
	r.POST("/deleteuser", mgr.AuthJwt, DeleteUser)
	r.POST("/checktoken", mgr.AuthJwt)
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, "See u")
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
	i <- 0
}
