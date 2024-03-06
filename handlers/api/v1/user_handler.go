package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all data"})
}

func (h *UserHandler) GetDataByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get data", "id": id})
}

func (h *UserHandler) CreateData(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create data"})
}

func (h *UserHandler) UpdateData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update data", "id": id})
}

func (h *UserHandler) DeleteData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete data", "id": id})
}
