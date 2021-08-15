package controller

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func RunGin(c context.Context, wg *sync.WaitGroup) {
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

	// wait for cacnel
	<-c.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Shutdown http server ...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}
	wg.Done()
}
