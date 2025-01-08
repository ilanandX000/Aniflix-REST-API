package controllers

import (
	"aniflix-restapi-web/initializers"
	"aniflix-restapi-web/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAnimes(c *gin.Context) {
    var animes []models.Anime
    if err := initializers.DB.Find(&animes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, animes)
}

func GetAnimeByID(c *gin.Context) {
    id := c.Param("id")
    var anime models.Anime
    if err := initializers.DB.First(&anime, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
        return
    }
    c.JSON(http.StatusOK, anime)
}

func CreateAnime(c *gin.Context) {
    var anime models.Anime
    if err := c.ShouldBindJSON(&anime); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := initializers.DB.Create(&anime).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, anime)
}

func UpdateAnime(c *gin.Context) {
    id := c.Param("id")
    var anime models.Anime
    if err := initializers.DB.First(&anime, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
        return
    }
    if err := c.ShouldBindJSON(&anime); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := initializers.DB.Save(&anime).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, anime)
}

func DeleteAnime(c *gin.Context) {
    id := c.Param("id")
    if err := initializers.DB.Delete(&models.Anime{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Anime deleted"})
}
