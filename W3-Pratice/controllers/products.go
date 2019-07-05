package controllers

import (
	"product-service/models"
	"product-service/tranfer"
	"github.com/gin-gonic/gin"
)

func HanldeGetProductsList(ctx *gin.Context) {
	products, err := models.GetProductsList()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Product List. Err:" + err.Error(),
		})

		return
	}

	categoriesData, _ := models.GetCategoryByIdList(&products)
	productsDto, _ := tranfer.MapProductForResponse(&products, categoriesData)

	ctx.JSON(200, &productsDto)
}
