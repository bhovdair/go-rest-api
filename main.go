package main

import (
	"net/http"

	handlers "github.com/bhovdair/go-rest-api/handlers/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello world!")
	})

	userHandler := handlers.NewUserHandler()
	productCategoryHandler := handlers.NewProductCategoryHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetData)
		v1.GET("/users/:id", userHandler.GetDataByID)
		v1.POST("/users", userHandler.CreateData)
		v1.PUT("/users/:id", userHandler.UpdateData)
		v1.DELETE("/users/:id", userHandler.DeleteData)

		v1.GET("/productcategories", productCategoryHandler.GetData)
		v1.GET("/productcategories/:id", productCategoryHandler.GetDataByID)
		v1.POST("/productcategories", productCategoryHandler.CreateData)
		v1.PUT("/productcategories/:id", productCategoryHandler.UpdateData)
		v1.DELETE("/productcategories/:id", productCategoryHandler.DeleteData)
	}

	router.Run("localhost:8080")
}
