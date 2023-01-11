package controller

import (
	"net/http"
	"time"

	"lepiosa/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userInput struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetAllUser godoc
// @Summary Get all Users.
// @Description Get a list of Users.
// @Tags Users
// @Produce json
// @Success 200 {object} []models.Users
// @Router /user [get]
func GetAllUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var user []models.Users
	db.Find(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser godoc
// @Summary Create New Users.
// @Description Creating a new Users.
// @Tags Users
// @Param Body body userInput true "the body to create a new Users"
// @Produce json
// @Success 200 {object} models.Users
// @Router /user [post]
func CreateUser(c *gin.Context) {
	// Validate input
	var input userInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	user := models.Users{Username: input.UserName, Email: input.Email, Password: input.Password}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUserById godoc
// @Summary Get Users.
// @Description Get an Users by id.
// @Tags Users
// @Produce json
// @Param id path string true "Users id"
// @Success 200 {object} models.Users
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) { // Get model if exist
	var user models.Users

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser godoc
// @Summary Update Users.
// @Description Update Users by id.
// @Tags Users
// @Produce json
// @Param id path string true "Users id"
// @Param Body body userInput true "the body to update age user category"
// @Success 200 {object} models.Users
// @Router /user/{id} [patch]
func UpdateUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.Users
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input userInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Users
	updatedInput.Username = input.UserName
	updatedInput.Email = input.Email
	updatedInput.Password = input.Password
	updatedInput.UpdatedAt = time.Now()

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser godoc
// @Summary Delete one Users.
// @Description Delete a Users by id.
// @Tags Users
// @Produce json
// @Param id path string true "Users id"
// @Success 200 {object} map[string]boolean
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var user models.Users
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
