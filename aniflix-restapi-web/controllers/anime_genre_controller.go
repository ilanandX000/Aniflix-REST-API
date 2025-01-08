package controllers

import (
	"aniflix-restapi-web/initializers"
	"aniflix-restapi-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAnimeGenres(c *gin.Context) {
	var animeGenres []models.AnimeGenre
	if err := initializers.DB.Find(&animeGenres).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, animeGenres)
}

func GetAnimeGenreByID(c *gin.Context) {
	id := c.Param("id")
	var animeGenre models.AnimeGenre
	if err := initializers.DB.First(&animeGenre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AnimeGenre not found"})
		return
	}
	c.JSON(http.StatusOK, animeGenre)
}

func CreateAnimeGenre(c *gin.Context) {
	var animeGenre models.AnimeGenre
	if err := c.ShouldBindJSON(&animeGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&animeGenre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, animeGenre)
}

func UpdateAnimeGenre(c *gin.Context) {
	id := c.Param("id")
	var animeGenre models.AnimeGenre
	if err := initializers.DB.First(&animeGenre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AnimeGenre not found"})
		return
	}
	if err := c.ShouldBindJSON(&animeGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Save(&animeGenre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, animeGenre)
}

func DeleteAnimeGenre(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.AnimeGenre{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "AnimeGenre deleted"})
}
