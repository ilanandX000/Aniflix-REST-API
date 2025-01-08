package controllers

import (
	"aniflix-restapi-web/initializers"
	"aniflix-restapi-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEpisodes(c *gin.Context) {
	var episodes []models.AnimeEpisode
	if err := initializers.DB.Find(&episodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, episodes)
}

func GetEpisodeByID(c *gin.Context) {
	id := c.Param("id")
	var episode models.AnimeEpisode
	if err := initializers.DB.First(&episode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
		return
	}
	c.JSON(http.StatusOK, episode)
}

func CreateEpisode(c *gin.Context) {
	var episode models.AnimeEpisode
	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, episode)
}

func UpdateEpisode(c *gin.Context) {
	id := c.Param("id")
	var episode models.AnimeEpisode
	if err := initializers.DB.First(&episode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
		return
	}
	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Save(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, episode)
}

func DeleteEpisode(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.AnimeEpisode{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Episode deleted"})
}
