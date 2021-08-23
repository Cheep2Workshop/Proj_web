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
	UserGroup(r, &mgr)
	ShopGroup(r, &mgr)
	ManageGroup(r, &mgr)
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

func UserGroup(r *gin.Engine, mgr *JWTManager) {
	g := r.Group("user")
	g.POST("/signup", Signup)
	g.POST("/setuser", mgr.AuthJwt, SetUser)
	g.POST("/login", Login)
	g.POST("/getloginlogs", mgr.AuthJwt, GetLoginLogs)
	g.POST("/getuser", GetUser)
	g.POST("/deleteuser", mgr.AuthJwt, DeleteUser)
	g.POST("/checktoken", mgr.AuthJwt)
}

func ShopGroup(r *gin.Engine, mgr *JWTManager) {
	g := r.Group("shop")
	g.POST("buy", mgr.AuthJwt, Buy)
	g.GET("list", ListProduct)
}

func ManageGroup(r *gin.Engine, mgr *JWTManager) {
	g := r.Group("manage")
	g.POST("addproduct", mgr.AuthJwt, AddProduct)
	g.POST("setproduct", mgr.AuthJwt, SetProduct)
	g.POST("deleteproduct", mgr.AuthJwt, DeleteProduct)
}
