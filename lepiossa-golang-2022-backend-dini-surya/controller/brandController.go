package controller

import (
	"net/http"
	"time"

	"lepiosa/models"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type brandInput struct {
	NamaBrand string `json:"namaBrand"`
}

// GetAllBrand godoc
// @Summary Get all Brand.
// @Description Get a list of Brand.
// @Tags Brand
// @Produce json
// @Success 200 {object} []models.Brand
// @Router /brand [get]
func GetAllBrand(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var brands []models.Brand
	db.Find(&brands)

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

// CreateBrand godoc
// @Summary Create New Brand.
// @Description Creating a new Brand.
// @Tags Brand
// @Param Body body brandInput true "the body to create a new Brand"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Brand
// @Router /brand [post]
func CreateBrand(c *gin.Context) {
	// Validate input
	var input brandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create brand
	brand := models.Brand{NamaBrand: input.NamaBrand}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&brand)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// GetBrandById godoc
// @Summary Get Brand.
// @Description Get an Brand by id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} models.Brand
// @Router /brand/{id} [get]
func GetBrandById(c *gin.Context) { // Get model if exist
	var brand models.Brand

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// GetLaptopsByBrandId godoc
// @Summary Get Laptops.
// @Description Get all Laptop by BrandID.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} []models.Laptop
// @Router /brand/{id}/laptop [get]
func GetLaptopByBrandId(c *gin.Context) { // Get model if exist
	var Laptops []models.Laptop

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("brand_id = ?", c.Param("id")).Find(&Laptops).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Laptops})
}

// UpdateBrand godoc
// @Summary Update Brand.
// @Description Update Brand by id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Param Body body brandInput true "the body to update brand "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Brand
// @Router /brand/{id} [patch]
func UpdateBrand(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input brandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Brand
	updatedInput.NamaBrand = input.NamaBrand

	updatedInput.UpdatedAt = time.Now()

	db.Model(&brand).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// DeleteBrand godoc
// @Summary Delete one Brand.
// @Description Delete a Brand by id.
// @Tags Brand
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} map[string]boolean
// @Router /brand/{id} [delete]
func DeleteBrand(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var brand models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&brand)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
