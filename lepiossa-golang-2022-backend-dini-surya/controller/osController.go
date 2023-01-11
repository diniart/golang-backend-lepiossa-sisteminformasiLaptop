package controller

import (
	"net/http"
	"time"

	"lepiosa/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OperasiSistemInput struct {
	NamaOs string `json:"namaOs"`
}

// GetAllOs godoc
// @Summary Get all OperasiSistem.
// @Description Get a list of OperasiSistem.
// @Tags OperasiSistem
// @Produce json
// @Success 200 {object} []models.OperasiSistem
// @Router /os [get]
func GetAllOs(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var os []models.OperasiSistem
	db.Find(&os)

	c.JSON(http.StatusOK, gin.H{"data": os})
}

// CreateOperasiSistem godoc
// @Summary Create New OperasiSistem.
// @Description Creating a new OperasiSistem.
// @Tags OperasiSistem
// @Param Body body OperasiSistemInput true "the body to create a new OperasiSistem"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.OperasiSistem
// @Router /os [post]
func CreateOperasiSistem(c *gin.Context) {
	// Validate input
	var input OperasiSistemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create os
	os := models.OperasiSistem{NamaOs: input.NamaOs}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&os)

	c.JSON(http.StatusOK, gin.H{"data": os})
}

// GetOperasiSistemById godoc
// @Summary Get OperasiSistem.
// @Description Get an OperasiSistem by id.
// @Tags OperasiSistem
// @Produce json
// @Param id path string true "OperasiSistem id"
// @Success 200 {object} models.OperasiSistem
// @Router /os/{id} [get]
func GetOperasiSistemById(c *gin.Context) { // Get model if exist
	var os models.OperasiSistem

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": os})
}

// GetLaptopsByOperasiSistemId godoc
// @Summary Get Laptops.
// @Description Get all Laptops by OperasiSistemId.
// @Tags OperasiSistem
// @Produce json
// @Param id path string true "OperasiSistem id"
// @Success 200 {object} []models.Laptop
// @Router /os/{id}/laptop [get]
func GetLaptopsByOperasiSistemId(c *gin.Context) { // Get model if exist
	var laptops []models.Laptop

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("OperasiSistemID = ?", c.Param("id")).Find(&laptops).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": laptops})
}

// UpdateOperasiSistem godoc
// @Summary Update OperasiSistem.
// @Description Update OperasiSistem by id.
// @Tags OperasiSistem
// @Produce json
// @Param id path string true "OperasiSistem id"
// @Param Body body OperasiSistemInput true "the body to update age os category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.OperasiSistem
// @Router /os/{id} [patch]
func UpdateOperasiSistem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var os models.OperasiSistem
	if err := db.Where("id = ?", c.Param("id")).First(&os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input OperasiSistemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.OperasiSistem
	updatedInput.NamaOs = input.NamaOs

	updatedInput.UpdatedAt = time.Now()

	db.Model(&os).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": os})
}

// DeleteOperasiSistem godoc
// @Summary Delete one OperasiSistem.
// @Description Delete a OperasiSistem by id.
// @Tags OperasiSistem
// @Produce json
// @Param id path string true "OperasiSistem id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /os/{id} [delete]
func DeleteOperasiSistem(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var os models.OperasiSistem
	if err := db.Where("id = ?", c.Param("id")).First(&os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&os)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
