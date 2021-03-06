package controller

import (
	"net/http"

	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/Cheep2Workshop/proj-web/service"
	"github.com/gin-gonic/gin"
)

func Buy(ctx *gin.Context) {
	var err error
	req := service.OrderReq{}

	err = ctx.Bind(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	order, err := service.AddOrder(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func ListProduct(ctx *gin.Context) {
	orders, err := repo.Client.GetAllProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
