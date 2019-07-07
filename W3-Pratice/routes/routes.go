package routes

import (
	"os"
	"product-service/config"
	"product-service/controllers"

	"github.com/gin-gonic/gin"
)

func Start() {
	port := os.Getenv("PRODUCT_SERVICE_PORT")

	if port == "" {
		port = config.PORT
	}

	route := gin.Default()
	route.GET("/products", controllers.HandleGetProductsList)
	route.GET("/categories", controllers.HandleGetCategoriesList)
	route.Run(":" + port)
}
