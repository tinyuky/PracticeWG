package controllers

import (
	"product-service/constants"
	"product-service/models"
	"product-service/tranfer"

	"github.com/gin-gonic/gin"
)

func HandleGetProductsList(ctx *gin.Context) {
	page := constants.PRODUCTLIST_PAGE
	limit := constants.PRODUCTLIST_LIMIT
	sort := constants.PRODUCTLIST_SORT
	by := constants.PRODUCTLIST_BY

	if ctx.Query("page") != "" {
		page = ctx.Query("page")
	}
	if ctx.Query("limit") != "" {
		limit = ctx.Query("limit")
	}
	if ctx.Query("sort") != "" {
		sort = ctx.Query("sort")
	}
	if ctx.Query("by") != "" {
		by = ctx.Query("by")
	}
	filter := ctx.QueryArray("filter")
	products, err := models.GetProductsList(page, limit, sort, by, &filter)

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
