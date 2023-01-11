// package brandModel
// package osModel
// package ratingModel

package controller

import (
	"net/http"
	"time"

	"lepiosa/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type laptopInput struct {
	NamaLaptop string `json:"namaLaptop"`
	OsID       uint   `json:"osID"`
	BrandID    uint   `json:"brandID"`
	Layar      string `json:"layar"`
	Hardisk    string `json:"hardisk"`
	Ram        string `json:"ram"`
	Grafis     string `json:"grafis"`
	RatingID   uint   `json:"ratingID"`
	Harga      string `json:"harga"`
	Image      string `json:"image"`
	Prosesor   string `json:"prosesor"`
	Detail     string `json:"detail"`
}

// GetAllLaptop godoc
// @Summary Get all laptops.
// @Description Get a list of Laptops.
// @Tags Laptop
// @Produce json
// @Success 200 {object} []models.Laptop
// @Router /laptop [get]
func GetAllLaptop(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var laptops []models.Laptop
	db.Find(&laptops)

	c.JSON(http.StatusOK, gin.H{"data": laptops})
}

// CreateLaptop godoc
// @Summary Create New Laptop.
// @Description Creating a new Laptop.
// @Tags Laptop
// @Param Body body laptopInput true "the body to create a new laptop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Laptop
// @Router /laptop [post]
func CreateLaptop(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input brand
	var input laptopInput
	var brand models.Brand
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.BrandID).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brandID not found!"})
		return
	}

	// Validate  os

	var os models.OperasiSistem

	if err := db.Where("id = ?", input.OsID).First(&os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "osID not found!"})
		return
	}
	// Validate  rating
	var rating models.Rating

	if err := db.Where("id = ?", input.RatingID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ratingID not found!"})
		return
	}

	// Create Laptop
	laptop := models.Laptop{NamaLaptop: input.NamaLaptop, OperasiSistemID: input.OsID, BrandID: input.BrandID, Layar: input.Layar, Hardisk: input.Hardisk, Ram: input.Ram, Grafis: input.Grafis, RatingID: input.RatingID, Harga: input.Harga, Image: input.Image, Prosesor: input.Prosesor, Detail: input.Detail}
	db.Create(&laptop)

	c.JSON(http.StatusOK, gin.H{"data": laptop})
}

// GetLaptopById godoc
// @Summary Get Laptop.
// @Description Get a Laptop by id.
// @Tags Laptop
// @Produce json
// @Param id path string true "laptop id"
// @Success 200 {object} models.Laptop
// @Router /laptop/{id} [get]
func GetLaptopById(c *gin.Context) { // Get model if exist
	var laptop models.Laptop

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&laptop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": laptop})
}

// UpdateLaptop godoc
// @Summary Update Laptop.
// @Description Update laptop by id.
// @Tags Laptop
// @Produce json
// @Param id path string true "laptop id"
// @Param Body body laptopInput true "the body to update an laptop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Laptop
// @Router /laptop/{id} [patch]
func UpdateLaptop(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var laptop models.Laptop
	var brand models.Brand
	var os models.OperasiSistem
	var rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&laptop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input laptopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.BrandID).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brandID not found!"})
		return
	}
	if err := db.Where("id = ?", input.OsID).First(&os).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "osID not found!"})
		return
	}
	if err := db.Where("id = ?", input.RatingID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "osID not found!"})
		return
	}

	var updatedInput models.Laptop
	updatedInput.NamaLaptop = input.NamaLaptop
	updatedInput.OperasiSistemID = input.OsID
	updatedInput.BrandID = input.BrandID
	updatedInput.Layar = input.Layar
	updatedInput.Hardisk = input.Hardisk
	updatedInput.Ram = input.Ram
	updatedInput.Grafis = input.Grafis
	updatedInput.RatingID = input.RatingID
	updatedInput.Harga = input.Harga
	updatedInput.Image = input.Image
	updatedInput.Prosesor = input.Prosesor
	updatedInput.Detail = input.Detail
	updatedInput.UpdatedAt = time.Now()

	db.Model(&laptop).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": laptop})
}

// DeleteLaptop godoc
// @Summary Delete one laptop.
// @Description Delete a laptop by id.
// @Tags Laptop
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "laptop id"
// @Success 200 {object} map[string]boolean
// @Router /laptop/{id} [delete]
func DeleteLaptop(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var laptop models.Laptop
	if err := db.Where("id = ?", c.Param("id")).First(&laptop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&laptop)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
