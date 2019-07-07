
package controllers

import (
	"product-service/models"
	"strings"
	"github.com/gin-gonic/gin"
)

func HandleGetCategoriesList(ctx *gin.Context) {
	idList := ctx.Query("ids")
	idListArray := strings.Split(idList, ",")
	categories, err := models.GetListCategory(idListArray)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Categories List. Err:" + err.Error(),
		})

		return
	}

	ctx.JSON(200, categories)
}
