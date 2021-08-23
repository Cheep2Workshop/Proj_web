package controller

import (
	"net/http"

	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/gin-gonic/gin"
)

func Buy(ctx *gin.Context) {
	req := repo.OrderReq{}

	ctx.Bind(req)
	if err := repo.Client.AddOrder(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func ListProduct(ctx *gin.Context) {
	orders, err := repo.Client.GetAllProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
