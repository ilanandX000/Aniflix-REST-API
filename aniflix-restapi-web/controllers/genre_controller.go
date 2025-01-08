package controllers

import (
	"aniflix-restapi-web/initializers"
	"aniflix-restapi-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGenres(c *gin.Context) {
	var genres []models.Genre
	if err := initializers.DB.Find(&genres).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genres)
}

func GetGenreByID(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre
	if err := initializers.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}
	c.JSON(http.StatusOK, genre)
}

func CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&genre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, genre)
}

func UpdateGenre(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre
	if err := initializers.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Save(&genre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genre)
}

func DeleteGenre(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.Genre{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted"})
}
