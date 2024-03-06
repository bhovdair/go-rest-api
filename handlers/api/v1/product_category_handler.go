package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
}

func NewProductCategoryHandler() *ProductCategoryHandler {
	return &ProductCategoryHandler{}
}

func (h *ProductCategoryHandler) GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all data"})
}

func (h *ProductCategoryHandler) GetDataByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get data", "id": id})
}

func (h *ProductCategoryHandler) CreateData(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create data"})
}

func (h *ProductCategoryHandler) UpdateData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update data", "id": id})
}

func (h *ProductCategoryHandler) DeleteData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete data", "id": id})
}
