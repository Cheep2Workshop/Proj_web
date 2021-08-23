package controller

import (
	"net/http"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/gin-gonic/gin"
)

func AddProduct(ctx *gin.Context) {
	var err error
	p := []models.Product{}
	err = ctx.Bind(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}

	err = repo.Client.AddProducts(p...)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func SetProduct(ctx *gin.Context) {
	var err error
	p := models.Product{}
	err = ctx.Bind(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	err = repo.Client.SetProduct(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, p)
}

func DeleteProduct(ctx *gin.Context) {
	var err error
	ids := []int{}
	err = ctx.Bind(&ids)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}
	err = repo.Client.DeleteProductById(ids...)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
