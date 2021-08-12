package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func RunGin() {
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

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen failed: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}
}
