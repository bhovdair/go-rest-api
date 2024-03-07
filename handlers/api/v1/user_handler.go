package handlers

import (
	"net/http"

	"github.com/bhovdair/go-rest-api/config"
	"github.com/bhovdair/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetData(c *gin.Context) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get all data", "data": users})
}

func (h *UserHandler) GetDataByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get data", "data": user})
}

func (h *UserHandler) CreateData(c *gin.Context) {

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	result := config.DB.Create(&newUser)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Create data", "data": newUser})
}

func (h *UserHandler) UpdateData(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	var existingUser models.User

	if err := config.DB.First(&existingUser, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	existingUser.Name = updatedUser.Name
	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email

	result := config.DB.Model(&existingUser).Updates(&updatedUser)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update data", "data": updatedUser})
}

func (h *UserHandler) DeleteData(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete data", "id": id})
}
